package feature

import (
	"context"

	"github.com/nggrjh/go-feature-manager/manager"
)

type Option func(opt *option)

type option struct {
	context  context.Context
	manager  manager.FeatureManager
	fallback bool
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

func WithFallback(fallback bool) Option {
	return func(opt *option) {
		opt.fallback = fallback
	}
}
