package logger

import (
	"context"
	"time"

	"gitlab.xiaoduoai.com/golib/xd_sdk/metadata"

	"github.com/sirupsen/logrus"
)

// 实现接口 XdLoggerEntry
type XdLogShadowEntry struct {
	*logrus.Entry
	nl *logrus.Logger
	sl *logrus.Logger
}

// logger 以及 entry的公用接口，方法一致，业务使用体验一致
type XdLoggerEntry interface {
	WithField(key string, value interface{}) XdLoggerEntry
	WithFields(fields Fields) XdLoggerEntry
	WithError(err error) XdLoggerEntry
	WithTime(t time.Time) XdLoggerEntry
	WithObject(obj interface{}) XdLoggerEntry
	Tracef(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Printf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Log(ctx context.Context, level Level, args ...interface{})
	Trace(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Print(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Panic(ctx context.Context, args ...interface{})
	Logln(ctx context.Context, level Level, args ...interface{})
	Traceln(ctx context.Context, args ...interface{})
	Debugln(ctx context.Context, args ...interface{})
	Infoln(ctx context.Context, args ...interface{})
	Println(ctx context.Context, args ...interface{})
	Warnln(ctx context.Context, args ...interface{})
	Warningln(ctx context.Context, args ...interface{})
	Errorln(ctx context.Context, args ...interface{})
	Fatalln(ctx context.Context, args ...interface{})
	Panicln(ctx context.Context, args ...interface{})
}

func (en XdLogShadowEntry) WithField(key string, value interface{}) XdLoggerEntry {
	return &XdLogShadowEntry{en.Entry.WithField(key, value), en.nl, en.sl}
}

func (en XdLogShadowEntry) WithFields(fields Fields) XdLoggerEntry {
	return &XdLogShadowEntry{en.Entry.WithFields(fields), en.nl, en.sl}
}

func (en XdLogShadowEntry) WithError(err error) XdLoggerEntry {
	return &XdLogShadowEntry{en.Entry.WithError(err), en.nl, en.sl}
}

func (en XdLogShadowEntry) WithTime(t time.Time) XdLoggerEntry {
	return &XdLogShadowEntry{en.Entry.WithTime(t), en.nl, en.sl}
}

func (en XdLogShadowEntry) WithObject(obj interface{}) XdLoggerEntry {
	fields := parseFieldsFromObj(obj)
	return &XdLogShadowEntry{en.Entry.WithFields(fields), en.nl, en.sl}
}

// setlogger，根据ctx压测标记，切换entry绑定的logger,将日志输出到不同的地方
func (en XdLogShadowEntry) setLogger(ctx context.Context) {
	if metadata.IsTestFlow(ctx) {
		en.Entry.Logger = en.sl
		return
	} else {
		en.Entry.Logger = en.nl
	}
}

func (en XdLogShadowEntry) Tracef(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logf(TraceLevel, format, args...)
}

func (en XdLogShadowEntry) Debugf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logf(DebugLevel, format, args...)
}

func (en XdLogShadowEntry) Infof(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logf(InfoLevel, format, args...)
}

func (en XdLogShadowEntry) Printf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Printf(format, args...)
}

func (en XdLogShadowEntry) Warnf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logf(WarnLevel, format, args...)
}

func (en XdLogShadowEntry) Warningf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Warnf(format, args...)
}

func (en XdLogShadowEntry) Errorf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logf(ErrorLevel, format, args...)
}

func (en XdLogShadowEntry) Fatalf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Fatalf(format, args...)
}

func (en XdLogShadowEntry) Panicf(ctx context.Context, format string, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logf(PanicLevel, format, args...)
}

func (en XdLogShadowEntry) Log(ctx context.Context, level Level, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Log(level, args...)
}

func (en XdLogShadowEntry) Trace(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Log(TraceLevel, args...)
}

func (en XdLogShadowEntry) Debug(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Log(DebugLevel, args...)
}

func (en XdLogShadowEntry) Info(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Log(InfoLevel, args...)
}

func (en XdLogShadowEntry) Print(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Print(args...)
}

func (en XdLogShadowEntry) Warn(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Log(WarnLevel, args...)
}

func (en XdLogShadowEntry) Warning(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Warn(args...)
}

func (en XdLogShadowEntry) Error(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Log(ErrorLevel, args...)
}

func (en XdLogShadowEntry) Fatal(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Fatal(args...)
}

func (en XdLogShadowEntry) Panic(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Panic(args...)
}

func (en XdLogShadowEntry) Logln(ctx context.Context, level Level, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(level, args...)
}

func (en XdLogShadowEntry) Traceln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(TraceLevel, args...)
}

func (en XdLogShadowEntry) Debugln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(DebugLevel, args...)
}

func (en XdLogShadowEntry) Infoln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(InfoLevel, args...)
}

func (en XdLogShadowEntry) Println(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Println(args...)
}

func (en XdLogShadowEntry) Warnln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(WarnLevel, args...)
}

func (en XdLogShadowEntry) Warningln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(WarnLevel, args...)
}

func (en XdLogShadowEntry) Errorln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(ErrorLevel, args...)
}

func (en XdLogShadowEntry) Fatalln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Fatalln(args...)
}

func (en XdLogShadowEntry) Panicln(ctx context.Context, args ...interface{}) {
	en.setLogger(ctx)
	en.Entry.WithContext(ctx).Logln(PanicLevel, args...)
}
