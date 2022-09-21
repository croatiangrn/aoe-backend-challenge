package app

var (
	Config *ApplicationConfig
)

type ApplicationConfig struct {
	superHeroesJSONPath string
}

func NewApplicationConfig(superHeroesJSONPath string) *ApplicationConfig {
	return &ApplicationConfig{superHeroesJSONPath: superHeroesJSONPath}
}

func (a *ApplicationConfig) SuperHeroesJSONPath() string {
	return a.superHeroesJSONPath
}
