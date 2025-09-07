package obs

import (
	"fmt"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"

	"github.com/Anacardo89/fizzbuzz-api/config"
)

func Start(cfg config.DD) (MetricsClient, error) {
	tracer.Start(
		tracer.WithService(cfg.Service),
		tracer.WithEnv(cfg.Env),
		tracer.WithServiceVersion(cfg.Version),
		tracer.WithAgentAddr(fmt.Sprintf("%s:%s", cfg.Agent, cfg.TracerPort)),
		tracer.WithSamplerRate(cfg.TracerSampleRate),
	)

	statsClient, err := statsd.New(
		fmt.Sprintf("%s:%s", cfg.Agent, cfg.MetricsPort),
		statsd.WithNamespace(cfg.MetricsNamespace),
		statsd.WithTags([]string{fmt.Sprintf("env:%s", cfg.Env), fmt.Sprintf("service:%s", cfg.Service)}),
	)
	if err != nil {
		tracer.Stop()
		return nil, err
	}

	return &ddMetrics{
		statsClient,
	}, nil
}
