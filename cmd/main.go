package main

import (
	"context"
	"io"

	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"
	"github.com/joshuabezaleel/library-o11y/persistence"
	"github.com/joshuabezaleel/library-o11y/server"

	"github.com/fluent/fluent-logger-golang/fluent"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
)

func main() {
	// time.Sleep(5 * time.Second)
	logger := log.NewLogger()
	logger.Log.SetFormatter(&logrus.JSONFormatter{})

	fluentLogger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "fluentd"})
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer fluentLogger.Close()

	tracer, closer := initJaeger("book-service")
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	tracer = opentracing.GlobalTracer()
	span := tracer.StartSpan("test-tracer")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	bookRepository := persistence.NewBookRepository(ctx, logger, fluentLogger)

	bookService := book.NewBookService(ctx, bookRepository, logger, fluentLogger)

	srv := server.NewServer(ctx, bookService, logger, fluentLogger)
	srv.Run("8082")
}

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := jaegermetrics.NullFactory

	tracer, closer, err := cfg.NewTracer(config.Logger(jLogger), config.Metrics(jMetricsFactory))
	if err != nil {
		panic(err)
	}

	return tracer, closer
}
