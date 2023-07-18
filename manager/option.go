package manager

type Option func(opt *option)

type option struct {
	fallback bool
	unleash  *unleashConfiguration
}

func WithFallback(fallback bool) Option {
	return func(opt *option) {
		opt.fallback = fallback
	}
}

func WithUnleashConfiguration(config *unleashConfiguration) Option {
	return func(opt *option) {
		opt.unleash = config
	}
}
