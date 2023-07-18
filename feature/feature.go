package feature

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nggrjh/go-feature-manager/manager"
)

//go:generate mockgen -source=feature.go -destination=mock/feature.go -package=mock

type IFeature interface {
	IsEnabled() bool
	String() string
}

func NewFeature(config *featureConfiguration, options ...Option) *feature {
	f := &feature{}

	var opt option
	for _, o := range options {
		o(&opt)
	}

	if opt.manager != nil {
		f.manager = opt.manager
	}

	f.name = config.buildName()
	f.fallback = opt.fallback
	return f
}

type feature struct {
	manager  manager.FeatureManager
	name     string
	fallback bool
}

func (f *feature) IsEnabled() bool {
	return f.manager.IsEnabled(f.name, f.fallback)
}

func (f *feature) String() string {
	b, _ := json.Marshal(map[string]any{
		"manager":  f.manager.String(),
		"name":     f.name,
		"fallback": f.fallback,
	})
	return string(b)
}

func NewConfiguration(app, behaviour string, category categoryType) *featureConfiguration {
	return &featureConfiguration{
		app:       app,
		behaviour: behaviour,
		category:  category,
	}
}

type featureConfiguration struct {
	app       string
	behaviour string
	category  categoryType
}

func (fc *featureConfiguration) buildName() string {
	toSnakeCase := func(s string) string {
		return strings.ToLower(strings.NewReplacer(" ", "-", "'", "").Replace(strings.TrimSpace(s)))
	}
	return fmt.Sprintf("%s_%s_%s", toSnakeCase(fc.app), fc.category.String(), toSnakeCase(fc.behaviour))
}
