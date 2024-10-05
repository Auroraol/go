package logger

import (
	"context"
	"os"
	"path/filepath"
	"reflect"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const shadowDirName = "shadow"

var AppName string

func NewLogger(opts ...Option) (l Logger, err error) {
	return NewLoggerWithOptions(newOptions(opts...))
}

func NewLoggerWithOptions(options Options) (l Logger, err error) {
	l = xdStdLoggerNew()
	if err = initLoggerWithOptions(l, options); err != nil {
		return nil, errors.Wrap(err, "failed to initialize logger")
	}
	return l, nil
}

func initLoggerWithOptions(l Logger, options Options) (err error) {
	AppName = options.AppName
	if options.Level != "" { // 如果配置里指定了日志等级，则解析并设置，否则默认等级是info。
		level, err := ParseLevel(options.Level)
		if err != nil {
			return errors.Wrapf(err, "failed to parse level(%s)", options.Level)
		}
		l.SetLevel(level, level)
	}

	if options.File != "" { // 如果配置里指定了日志文件，则解析并设置，否则默认写到stderr。
		err = handleFileOutput(l, options.File) // 设置output、压测标志
		if err != nil {
			errors.Wrapf(err, "failed to set logger.Output and set flow_control")
		}
	}
	l.AddHook(NewTraceHook(), NewTraceHook())       // 拦截日志里面跟trace相关的字段。
	l.AddHook(NewFileLineHook(), NewFileLineHook()) // 在日志中输出文件名和行号。
	if options.Format == "json" || options.Format == "" {
		l.AddHook(NewMergeHook(), NewMergeHook()) //拦截日志里面其余的信息到custom中，只保留固定的字段
		l.SetFormatter(newJSONFormatter(), newJSONFormatter())
	} else {
		l.SetFormatter(newTextFormatter(), newTextFormatter())
	}

	if options.ErrFile != "" { // 如果配置里指定了错误日志文件，则额外将等级为error(及以上)的日志复制一份写到该文件中。
		errWriter, err := os.OpenFile(options.ErrFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return errors.Wrapf(err, "failed to open err file(%s)", options.ErrFile)
		}
		l.AddHook(NewErrWriterHook(errWriter), nil) // shadow区域日志默认不复制
	}
	return
}

func handleFileOutput(l Logger, fileName string) error {
	writer, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return errors.Wrapf(err, "failed to open file(%s)", fileName)
	}
	sWriter := writer
	sFilename, err := genShadowFileName(fileName)
	if err == nil {
		sf, err := os.OpenFile(sFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			Warningf(context.Background(), "failed to open file(%s), err(%v)", sFilename, err)
		} else {
			sWriter = sf
		}
	}
	l.SetOutput(writer, sWriter) // 设置正常日志、影子区域日志输出
	return nil
}

func genShadowFileName(fn string) (string, error) {
	if fn == "" {
		return "", errors.New("FileName empty")
	}

	path, name := filepath.Split(fn)
	if name == "" {
		return "", errors.New("FileName empty after split")
	}

	shadowPath := filepath.Join(path, shadowDirName) // 构造影子目录，并尝试创建目录
	_, err := os.Stat(shadowPath)
	if !os.IsExist(err) {
		err = os.MkdirAll(shadowPath, 0666)
	}

	return filepath.Join(path, shadowDirName, name), err
}

// see from https://gitlab.xiaoduoai.com/marketing/base/blob/master/log/log.go
func parseFieldsFromObj(o interface{}) logrus.Fields {
	logFields := logrus.Fields{}

	val := reflect.ValueOf(o)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return logFields
		}
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		fValue := val.Field(i)
		fType := val.Type().Field(i)
		if !isZero(fValue) && fValue.IsValid() && fType.PkgPath == "" { // exported fields
			if fValue.Kind() == reflect.Struct ||
				(fValue.Kind() == reflect.Ptr &&
					fValue.Elem().Kind() == reflect.Struct) {
				fields := parseFieldsFromObj(fValue.Interface())
				if fType.Anonymous {
					for k, v := range fields {
						logFields[k] = v
					}
				} else {
					logFields[fType.Name] = fields
				}
			} else {
				logFields[fType.Name] = fValue.Interface()
			}
		}
	}
	return logFields
}

// see https://gitlab.xiaoduoai.com/marketing/base/blob/master/log/log.go
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct: // 不去确认
		return false
	}
	return false
}
