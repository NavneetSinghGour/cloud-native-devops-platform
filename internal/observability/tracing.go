package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer(serviceName string) (*sdktrace.TracerProvider, error) {

	ctx := context.Background()

	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(

		sdktrace.WithBatcher(exporter),

		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName(serviceName),
			),
		),
	)

	otel.SetTracerProvider(tp)

	return tp, nil
}
