package logger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// XdLogger的接口
type XdLogger interface {
	SetOutput(out, shadowOut io.Writer)
	SetFormatter(formatter, shadowFormatter Formatter)
	SetReportCaller(include, shadowInclude bool)
	SetLevel(level, shadowLevel logrus.Level)
	AddHook(hook, shadowHook logrus.Hook)
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

type CtxLogger struct {
	n *logrus.Logger // normal logger
	s *logrus.Logger // shadow logger
}

func NewCtxLogger() XdLogger {
	formatter := new(JSONFormatter)
	formatter.TimestampFormat = "2006-01-02T15:04:05.000Z07:00"

	n := logrus.Logger{
		Out:          os.Stderr,
		Formatter:    formatter,
		Hooks:        make(LevelHooks),
		Level:        InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	s := logrus.Logger{
		Out:          os.Stderr,
		Formatter:    formatter,
		Hooks:        make(LevelHooks),
		Level:        InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	return &CtxLogger{&n, &s}
}

func (cl *CtxLogger) newXdLogShadowEntry() XdLoggerEntry {
	return &XdLogShadowEntry{logrus.NewEntry(cl.n), cl.n, cl.s}
}

func (cl *CtxLogger) hasUniqShadowLogger() bool {
	return cl.s != nil && cl.s != cl.n
}

func (cl *CtxLogger) SetOutput(out, shadowOut io.Writer) {
	cl.n.SetOutput(out)
	if cl.hasUniqShadowLogger() {
		cl.s.SetOutput(shadowOut)
	}
}

func (cl *CtxLogger) SetFormatter(formatter, shadowFormatter Formatter) {
	cl.n.SetFormatter(formatter)
	if cl.hasUniqShadowLogger() {
		cl.s.SetFormatter(shadowFormatter)
	}
}

func (cl *CtxLogger) SetReportCaller(include, shadowInclude bool) {
	cl.n.SetReportCaller(include)
	if cl.hasUniqShadowLogger() {
		cl.s.SetReportCaller(shadowInclude)
	}
}

func (cl *CtxLogger) SetLevel(level, shadowLevel logrus.Level) {
	cl.n.SetLevel(level)
	if cl.hasUniqShadowLogger() {
		cl.s.SetLevel(shadowLevel)
	}
}

func (cl *CtxLogger) AddHook(hook, shadowHook logrus.Hook) {
	cl.n.AddHook(hook)
	if cl.hasUniqShadowLogger() && shadowHook != nil {
		cl.s.AddHook(shadowHook)
	}
}

func (cl *CtxLogger) WithField(key string, value interface{}) XdLoggerEntry {
	// 借用logrus.Logger本身Entry的管理机制来创建Entry,下同
	return &XdLogShadowEntry{cl.n.WithField(key, value), cl.n, cl.s}
}

func (cl *CtxLogger) WithFields(fields Fields) XdLoggerEntry {
	return &XdLogShadowEntry{cl.n.WithFields(fields), cl.n, cl.s}
}

func (cl *CtxLogger) WithError(err error) XdLoggerEntry {
	return &XdLogShadowEntry{cl.n.WithError(err), cl.n, cl.s}
}

func (cl *CtxLogger) WithTime(t time.Time) XdLoggerEntry {
	return &XdLogShadowEntry{cl.n.WithTime(t), cl.n, cl.s}
}

func (cl *CtxLogger) WithObject(obj interface{}) XdLoggerEntry {
	fields := parseFieldsFromObj(obj)
	return &XdLogShadowEntry{cl.n.WithFields(fields), cl.n, cl.s}
}

func (cl *CtxLogger) Tracef(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Tracef(ctx, format, args...)
}

func (cl *CtxLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Debugf(ctx, format, args...)
}

func (cl *CtxLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Infof(ctx, format, args...)
}

func (cl *CtxLogger) Printf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Printf(ctx, format, args...)
}

func (cl *CtxLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Warnf(ctx, format, args...)
}

func (cl *CtxLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Warningf(ctx, format, args...)
}

func (cl *CtxLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Errorf(ctx, format, args...)
}

func (cl *CtxLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Fatalf(ctx, format, args...)
}

func (cl *CtxLogger) Panicf(ctx context.Context, format string, args ...interface{}) {
	cl.newXdLogShadowEntry().Panicf(ctx, format, args...)
}

func (cl *CtxLogger) Log(ctx context.Context, level Level, args ...interface{}) {
	cl.newXdLogShadowEntry().Log(ctx, level, args...)
}

func (cl *CtxLogger) Trace(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Trace(ctx, args...)
}

func (cl *CtxLogger) Debug(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Debug(ctx, args...)
}

func (cl *CtxLogger) Info(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Info(ctx, args...)
}

func (cl *CtxLogger) Print(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Print(ctx, args...)
}

func (cl *CtxLogger) Warn(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Warn(ctx, args...)
}

func (cl *CtxLogger) Warning(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Warning(ctx, args...)
}

func (cl *CtxLogger) Error(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Error(ctx, args...)
}

func (cl *CtxLogger) Fatal(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Fatal(ctx, args...)
}

func (cl *CtxLogger) Panic(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Panic(ctx, args...)
}

func (cl *CtxLogger) Logln(ctx context.Context, level Level, args ...interface{}) {
	cl.newXdLogShadowEntry().Logln(ctx, level, args...)
}

func (cl *CtxLogger) Traceln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Traceln(ctx, args...)
}

func (cl *CtxLogger) Debugln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Debugln(ctx, args...)
}

func (cl *CtxLogger) Infoln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Infoln(ctx, args...)
}

func (cl *CtxLogger) Println(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Println(ctx, args...)
}

func (cl *CtxLogger) Warnln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Warnln(ctx, args...)
}

func (cl *CtxLogger) Warningln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Warningln(ctx, args...)
}

func (cl *CtxLogger) Errorln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Errorln(ctx, args...)
}

func (cl *CtxLogger) Fatalln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Fatalln(ctx, args...)
}

func (cl *CtxLogger) Panicln(ctx context.Context, args ...interface{}) {
	cl.newXdLogShadowEntry().Panicln(ctx, args...)
}
