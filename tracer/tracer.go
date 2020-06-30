package tracer

import (
	"fmt"
	"io"
	stdlog "log"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var closeTracer io.Closer

func InitTracer(serviceName string) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	tracer, c, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return err
	}
	closeTracer = c

	opentracing.SetGlobalTracer(tracer)
	return nil
}

func Close() {
	if err := closeTracer.Close(); err != nil {
		UntracedLogf("error flushing traces: %e", err)
	}
}

func SetError(span opentracing.Span, err error) {
	ext.Error.Set(span, true)
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)
}

func Log(span opentracing.Span, str string) {
	UntracedLog(str)

	span.LogFields(log.String("message", str))
}

func Logf(span opentracing.Span, format string, args ...interface{}) {
	Log(span, fmt.Sprintf(format, args...))
}

func UntracedLog(str string) {
	stdlog.Print(str)
}

func UntracedLogf(format string, args ...interface{}) {
	stdlog.Printf(format, args...)
}
