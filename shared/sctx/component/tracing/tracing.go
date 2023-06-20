package tracing

import (
	"context"
	"flag"

	"jetshop/shared/sctx"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

type tracingClient struct {
	id            string
	serviceName   string
	version       string
	collectorHost string
	logger        sctx.Logger
	tp            *tracesdk.TracerProvider
}

func NewTracingClient(id string, serviceName string, version string) *tracingClient {
	tracer = otel.Tracer(serviceName)

	return &tracingClient{id: id, serviceName: serviceName, version: version}
}

func (t *tracingClient) ID() string {
	return t.id
}

func (t *tracingClient) InitFlags() {
	flag.StringVar(&t.collectorHost, "collector_host", "localhost:5555", "collector host")
}

func (t *tracingClient) Activate(sc sctx.ServiceContext) error {
	t.logger = sc.Logger(t.id)

	ctx := context.Background()

	exporter, err := otlptrace.New(ctx, otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(t.collectorHost),
		otlptracegrpc.WithInsecure(),
	))

	if err != nil {
		return err
	}

	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(t.serviceName),
		semconv.ServiceVersionKey.String(t.version),
	)

	t.tp = tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource),
	)

	otel.SetTracerProvider(t.tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}

func (t *tracingClient) Stop() error {
	return t.tp.Shutdown(context.Background())
}
