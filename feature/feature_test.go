package feature

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/nggrjh/go-feature-manager/manager/mock"
)

func Test_feature_IsEnabled(t *testing.T) {
	t.Parallel()
	type fields struct {
		app          string
		behaviour    string
		category     categoryType
		creationTime time.Time
		fallback     *bool
	}
	type expectIsEnabled struct {
		name       string
		fallback   bool
		rIsEnabled bool
	}

	tests := map[string]struct {
		fields fields
		want   bool

		expectIsEnabled *expectIsEnabled
	}{
		"it_should_return_true": {
			fields: fields{
				app:       "demo-app",
				behaviour: "Enable API Readiness",
				category:  CategoryTemporary,
			},
			want: true,

			expectIsEnabled: &expectIsEnabled{
				name:       "demo-app_t_enable-api-readiness",
				fallback:   false,
				rIsEnabled: true,
			},
		},
		"it_should_return_true_with_fallback": {
			fields: fields{
				app:       "demo-app",
				behaviour: "Enable API Readiness",
				category:  CategoryTemporary,
				fallback:  func(b bool) *bool { return &b }(true),
			},
			want: true,

			expectIsEnabled: &expectIsEnabled{
				name:       "demo-app_t_enable-api-readiness",
				fallback:   true,
				rIsEnabled: true,
			},
		},
		"it_should_return_false": {
			fields: fields{
				app:       "demo-app",
				behaviour: "Enable API Readiness",
				category:  CategoryTemporary,
			},
			want: false,

			expectIsEnabled: &expectIsEnabled{
				name:       "demo-app_t_enable-api-readiness",
				fallback:   false,
				rIsEnabled: false,
			},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			control := gomock.NewController(t)
			t.Cleanup(control.Finish)

			mockManager := mock.NewMockFeatureManager(control)

			if expect := tt.expectIsEnabled; expect != nil {
				mockManager.EXPECT().IsEnabled(expect.name, expect.fallback).Return(expect.rIsEnabled)
			}

			opts := []Option{
				WithContext(context.Background()),
				WithManager(mockManager),
			}
			if fallback := tt.fields.fallback; fallback != nil {
				opts = append(opts, WithFallback(*fallback))
			}

			f := NewFeature(
				NewConfiguration(tt.fields.app, tt.fields.behaviour, tt.fields.category, tt.fields.creationTime),
				opts...,
			)
			if got := f.IsEnabled(); got != tt.want {
				t.Errorf("IsEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_featureConfiguration_buildName(t *testing.T) {
	t.Parallel()
	type fields struct {
		app          string
		behaviour    string
		category     categoryType
		creationTime time.Time
	}
	tests := map[string]struct {
		fields fields
		want   string
	}{
		"it_should_return_temporary": {
			fields: fields{
				app:       "demo-app",
				behaviour: "Enable API Readiness",
				category:  CategoryTemporary,
			},
			want: "demo-app_t_enable-api-readiness",
		},
		"it_should_return_permanent": {
			fields: fields{
				app:       "demo-app",
				behaviour: "Enable API Readiness",
				category:  CategoryPermanent,
			},
			want: "demo-app_p_enable-api-readiness",
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fc := NewConfiguration(
				tt.fields.app,
				tt.fields.behaviour,
				tt.fields.category,
				tt.fields.creationTime,
			)
			if got := fc.buildName(); got != tt.want {
				t.Errorf("buildName() = %v, want %v", got, tt.want)
			}
		})
	}
}
