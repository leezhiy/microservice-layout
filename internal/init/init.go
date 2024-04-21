package init

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"git.corp.doulaoban.com/backend/kratos-layout/internal/conf"
	"git.corp.doulaoban.com/backend/kratos-layout/pkg/log/aliyun"
)

// Config 初始化配置文件
func Config() *Configures {
	return &Configures{
		Config: &conf.Bootstrap{},
	}
}

// Logger 初始化日志
func Logger(conf *conf.Logger, server *conf.Server, id string, version string) log.Logger {
	logstore := server.Environment + "-" + server.Name
	if conf.Logstore != "" {
		logstore = conf.Logstore
	}
	return log.With(aliyun.NewAliyunLog(
		aliyun.WithProject(conf.Project),
		aliyun.WithEndpoint(conf.Endpoint),
		aliyun.WithLogstore(logstore),
		aliyun.WithAccessKey(conf.AccessKey),
		aliyun.WithAccessSecret(conf.AccessSecret),
		aliyun.WithLogLevel(conf.Level),
	),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", server.Name,
		"service.version", version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

}

// Sentry 初始化监控告警
func Sentry(c *conf.Sentry, logger log.Logger) error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              c.Dsn,
		Debug:            c.Debug,
		DebugWriter:      log.NewWriter(logger),
		Environment:      c.Env,
		AttachStacktrace: true,
	})
}

// Tracer 初始化链路追踪
func Tracer(tracer *conf.Tracer, server *conf.Server) error {
	ctx := context.Background()

	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(tracer.Endpoint))
	if err != nil {
		return err
	}

	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100% 采样率
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(tracer.Sampler))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(server.Name),
			attribute.String("env", server.Environment),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
