package main

import (
	"context"
	"fmt"

	"github.com/nggrjh/go-feature-manager/feature"
	"github.com/nggrjh/go-feature-manager/manager"
)

func main() {
	mngr, err := manager.NewUnleashManager()
	if err != nil {
		panic(err)
	}

	f := feature.NewFeature(
		feature.NewConfiguration("demo-app", "", feature.CategoryTemporary),
		feature.WithContext(context.Background()),
		feature.WithManager(mngr),
		feature.WithFallback(false),
	)
	fmt.Println(f, f.IsEnabled())
}
