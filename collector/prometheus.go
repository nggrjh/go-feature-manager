package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func NewPrometheusCollector(app string) (FeatureCollector, error) {
	const subsystem = "Feature"
	var labels = []string{"behaviour", "category", "evaluation"}

	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   app,
			Subsystem:   subsystem,
			Name:        "ReleaseLifetime",
			Help:        "Release lifetime metrics",
			ConstLabels: map[string]string{},
		},
		labels,
	)
	if err := prometheus.Register(gauge); err != nil {
		return nil, err
	}

	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   app,
			Subsystem:   subsystem,
			Name:        "EvaluationRate",
			Help:        "Evaluation rate metrics",
			ConstLabels: map[string]string{},
			Buckets:     []float64{.02, .05, 0.2, 0.3, 0.5, 1, 2.5, 5, 10},
		},
		labels,
	)
	if err := prometheus.Register(histogram); err != nil {
		return nil, err
	}

	return &prometheusCollector{gauge: gauge, histogram: histogram}, nil
}

type prometheusCollector struct {
	gauge     *prometheus.GaugeVec
	histogram *prometheus.HistogramVec
}

func (p *prometheusCollector) Observe(duration, lifetime time.Duration, labelValues ...string) {
	p.histogram.WithLabelValues(labelValues...).Observe(duration.Seconds())
	p.gauge.WithLabelValues(labelValues...).Set(lifetime.Seconds())
}
