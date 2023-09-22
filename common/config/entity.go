package config

var Configs Config

type Config struct {
	Application Application `mapstructure:"application"`
}

type Application struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}
