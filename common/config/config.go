/*
Copyright © 2023 xujiahao <1787619881@qq.com>
*/
package config

import (
	"fmt"
	"reflect"
	"strings"
)

// GetFlagMap is a function in Golang that retrieves a map. In this map,
// each key corresponds to a configuration key,
// while the associated value represents a reflect.Value object
// for the respective field in the Configs struct.
func GetFlagMap() map[string]reflect.Value {
	m := make(map[string]reflect.Value)
	err := fillMap(&m, "", "", Configs)
	if err != nil {
		panic(err)
	}
	return m
}

func fillMap(m *map[string]reflect.Value, parentKey string, name string, c interface{}) error {
	var key string
	if parentKey != "" && name != "" {
		key = parentKey + "." + name
	} else {
		key = parentKey + name
	}

	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
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
			err := fillMap(m, key, n, fv.Interface())
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("error occurred while parsing flag config: unsupport type in struct")
	}
}
