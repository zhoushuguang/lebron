package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	ztrace "github.com/zeromicro/go-zero/core/trace"
)

func StartSpan(ctx context.Context, method string) (context.Context, trace.Span) {
	tracer := otel.Tracer(ztrace.TraceName)
	return tracer.Start(ctx, method, trace.WithSpanKind(trace.SpanKindInternal))
}
