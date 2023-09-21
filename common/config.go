/*
Copyright © 2023 xujiahao <1787619881@qq.com>
*/
package common

var Configs Config

type Config struct {
	Application Application `mapstructure:"application"`
}

type Application struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}
