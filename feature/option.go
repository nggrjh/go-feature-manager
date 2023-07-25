package feature

import (
	"context"

	"github.com/nggrjh/go-feature-manager/collector"
	"github.com/nggrjh/go-feature-manager/manager"
)

type Option func(opt *option)

type option struct {
	context   context.Context
	manager   manager.FeatureManager
	collector collector.FeatureCollector
	fallback  bool
}

func WithContext(ctx context.Context) Option {
	return func(opt *option) {
		opt.context = ctx
	}
}

func WithManager(manager manager.FeatureManager) Option {
	return func(opt *option) {
		opt.manager = manager
	}
}

func WithCollector(collector collector.FeatureCollector) Option {
	return func(opt *option) {
		opt.collector = collector
	}
}

func WithFallback(fallback bool) Option {
	return func(opt *option) {
		opt.fallback = fallback
	}
}
