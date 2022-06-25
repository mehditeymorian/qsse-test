package subscriber

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	SubscriberCounter *prometheus.GaugeVec
}

func NewMetrics(namespace string) Metrics {
	var metric Metrics

	metric.SubscriberCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{ //nolint:exhaustruct
		Namespace: namespace,
		Subsystem: "qsse",
		Name:      "subscriber_count",
		Help:      "count of topic's subscribers",
	}, []string{"topic"})

	return metric
}

func (m Metrics) IncSubscriber(topic string) {
	m.SubscriberCounter.With(map[string]string{
		"topic": topic,
	}).Inc()
}

func (m Metrics) DecSubscriber(topic string) {
	m.SubscriberCounter.With(map[string]string{
		"topic": topic,
	}).Dec()
}
