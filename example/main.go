package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nggrjh/go-feature-manager/collector"
	"github.com/nggrjh/go-feature-manager/feature"
	"github.com/nggrjh/go-feature-manager/manager"
)

func main() {
	const app = "demo-app"

	mngr, err := manager.NewUnleashManager()
	if err != nil {
		panic(err)
	}

	coll, err := collector.NewPrometheusCollector(app)
	if err != nil {
		panic(err)
	}

	f := feature.NewFeature(
		feature.NewConfiguration(app, "test", feature.CategoryTemporary, time.Now()),
		feature.WithContext(context.Background()),
		feature.WithManager(mngr),
		feature.WithCollector(coll),
		feature.WithFallback(false),
	)
	fmt.Println(f, f.IsEnabled())
}
