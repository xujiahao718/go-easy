/*
Copyright © 2023 xujiahao <1787619881@qq.com>
*/
package common

import (
	"fmt"
	"reflect"
	"strings"
)

var Configs Config

type Config struct {
	Application Application `mapstructure:"application"`
	IsOk        bool        `mapstructure:"isok"`
	Slice       []string    `mapstructure:"slice"`
}

type Application struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

func GetFlagMap() map[string]reflect.Value {
	m := make(map[string]reflect.Value)
	err := parseFlagConfig("", "", Configs, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func parseFlagConfig(parentKey string, name string, c interface{}, m *map[string]reflect.Value) error {
	var key string
	if parentKey != "" && name != "" {
		key = parentKey + "." + name
	} else {
		key = parentKey + name
	}

	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	fmt.Println("===", key)
	switch t.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String,
		reflect.Array, reflect.Slice:
		(*m)[key] = v
		return nil
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fv := v.Field(i)
			n := strings.ToLower(f.Name)
			err := parseFlagConfig(key, n, fv.Interface(), m)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("error occurred while parsing flag config: unsupport type in struct")
	}
}
