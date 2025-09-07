package obs

import "time"

type MockMetrics struct {
	incrCalls   []string
	timingCalls []time.Duration
}

func (m *MockMetrics) Close() error {
	return nil
}

func (m *MockMetrics) Incr(name string, tags []string, rate float64) error {
	m.incrCalls = append(m.incrCalls, name)
	return nil
}

func (m *MockMetrics) Timing(name string, value time.Duration, tags []string, rate float64) error {
	m.timingCalls = append(m.timingCalls, value)
	return nil
}

func (m *MockMetrics) Histogram(name string, value float64, tags []string, rate float64) error {
	return nil
}
