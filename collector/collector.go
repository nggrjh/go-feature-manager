package collector

import "time"

//go:generate mockgen -source=collector.go -destination=mock/collector.go -package=mock

type FeatureCollector interface {
	Observe(duration, lifetime time.Duration, labelValues ...string)
}
