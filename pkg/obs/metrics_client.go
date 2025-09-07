package obs

import (
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
)

type MetricsClient interface {
	Close() error
	Incr(name string, tags []string, rate float64) error
	Timing(name string, value time.Duration, tags []string, rate float64) error
	Histogram(name string, value float64, tags []string, rate float64) error
}

type ddMetrics struct {
	*statsd.Client
}

func (d *ddMetrics) Close() error {
	return d.Client.Close()
}

func (d *ddMetrics) Incr(name string, tags []string, rate float64) error {
	return d.Client.Incr(name, tags, rate)
}

func (d *ddMetrics) Timing(name string, value time.Duration, tags []string, rate float64) error {
	return d.Client.Timing(name, value, tags, rate)
}

func (d *ddMetrics) Histogram(name string, value float64, tags []string, rate float64) error {
	return d.Client.Histogram(name, value, tags, rate)
}
