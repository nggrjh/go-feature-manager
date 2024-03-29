package feature

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nggrjh/go-feature-manager/collector"
	"github.com/nggrjh/go-feature-manager/manager"
)

//go:generate mockgen -source=feature.go -destination=mock/feature.go -package=mock

type Feature interface {
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

	if opt.collector != nil {
		f.collector = opt.collector
	}

	f.name = config.buildName()
	f.category = config.category
	f.releaseTime = config.creationTime
	f.fallback = opt.fallback
	return f
}

type feature struct {
	manager     manager.FeatureManager
	collector   collector.FeatureCollector
	name        string
	category    categoryType
	releaseTime time.Time
	fallback    bool
}

func (f *feature) IsEnabled() bool {
	start := time.Now()
	isEnabled := f.manager.IsEnabled(f.name, f.fallback)
	duration := time.Since(start)

	if f.collector != nil {
		f.collect(f.name, isEnabled, duration)
	}

	return isEnabled
}

func (f *feature) String() string {
	b, _ := json.Marshal(map[string]any{
		"manager":     f.manager.String(),
		"name":        f.name,
		"category":    f.category.String(),
		"releaseTime": f.releaseTime.String(),
		"fallback":    f.fallback,
	})
	return string(b)
}

func (f *feature) collect(name string, evaluation bool, duration time.Duration) {
	f.collector.Observe(duration, time.Since(f.releaseTime), name, f.category.String(), strconv.FormatBool(evaluation))
}

func NewConfiguration(app, behaviour string, category categoryType, creationTime time.Time) *featureConfiguration {
	return &featureConfiguration{
		app:          app,
		behaviour:    behaviour,
		category:     category,
		creationTime: creationTime,
	}
}

type featureConfiguration struct {
	app          string
	behaviour    string
	category     categoryType
	creationTime time.Time
}

func (fc *featureConfiguration) buildName() string {
	toSnakeCase := func(s string) string {
		return strings.ToLower(strings.NewReplacer(" ", "-", "'", "").Replace(strings.TrimSpace(s)))
	}
	return fmt.Sprintf("%s_%s_%s", toSnakeCase(fc.app), fc.category.String(), toSnakeCase(fc.behaviour))
}
