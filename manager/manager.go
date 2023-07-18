package manager

//go:generate mockgen -source=manager.go -destination=mock/manager.go -package=mock

type FeatureManager interface {
	IsEnabled(name string, fallback bool) bool
	String() string
}
