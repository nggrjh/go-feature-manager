package manager

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/spf13/viper"
)

func NewUnleashManager(options ...Option) (FeatureManager, error) {
	viper.AutomaticEnv()
	cfg := &unleashConfiguration{
		appName:         viper.GetString("UNLEASH_APP_NAME"),
		url:             viper.GetString("UNLEASH_URL"),
		apiToken:        viper.GetString("UNLEASH_API_TOKEN"),
		refreshInterval: viper.GetDuration("UNLEASH_REFRESH_INTERVAL"),
	}

	var opt option
	for _, o := range options {
		o(&opt)
	}

	if opt.unleash != nil {
		cfg = opt.unleash
	}

	client, err := unleash.NewClient(
		unleash.WithAppName(cfg.appName),
		unleash.WithUrl(cfg.url),
		unleash.WithCustomHeaders(http.Header{"Authorization": {cfg.apiToken}}),
		unleash.WithRefreshInterval(cfg.refreshInterval),
		unleash.WithListener(&struct{}{}),
	)
	if err != nil {
		return nil, err
	}
	client.WaitForReady()
	return &unleashManager{client: client}, nil
}

type unleashManager struct {
	client *unleash.Client
}

func (u *unleashManager) IsEnabled(name string, fallback bool) bool {
	opts := []unleash.FeatureOption{unleash.WithFallback(fallback)}
	return u.client.IsEnabled(name, opts...)
}

func (u *unleashManager) String() string {
	b, _ := json.Marshal(map[string]any{
		"manager": "unleash",
	})
	return string(b)
}

type unleashConfiguration struct {
	appName         string
	url             string
	apiToken        string
	refreshInterval time.Duration
}
