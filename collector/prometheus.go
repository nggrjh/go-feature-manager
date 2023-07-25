package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func NewPrometheusCollector() (FeatureCollector, error) {
	return &prometheusCollector{}, nil
}

type prometheusCollector struct {
	gauge     *prometheus.GaugeVec
	histogram *prometheus.HistogramVec
}

func (p *prometheusCollector) Observe(duration, lifetime time.Duration, labelValues ...string) {
	p.histogram.WithLabelValues(labelValues...).Observe(duration.Seconds())
	p.gauge.WithLabelValues(labelValues...).Set(lifetime.Seconds())
}
