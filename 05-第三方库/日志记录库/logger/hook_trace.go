package logger

import (
	"gitlab.xiaoduoai.com/golib/xd_sdk/octrace"
)

// 从context中提取出与trace相关的一些字段写入日志中。
const TraceKey = "trace"
const SpanKey = "span"
const SampledKey = "sampled"
const ParentSpanKey = "pspan"
const TimeStamp = "timestamp"
const customKey = "custom"

type TraceHook struct {
}

func NewTraceHook() *TraceHook {
	return &TraceHook{}
}

func (h *TraceHook) Fire(entry *Entry) error {
	if ctx := entry.Context; ctx != nil {
		span := octrace.ExtractSpan(ctx)
		if span != nil {
			entry.Data[TraceKey] = span.SpanContext().TraceID.String()
		}
	}

	return nil
}

func (h *TraceHook) Levels() []Level {
	return AllLevels
}
