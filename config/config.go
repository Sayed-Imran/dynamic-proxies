package config

import (
	"fmt"
	"reflect"
	"strings"
)

type BaseConfig struct {
	Port           string `env:"PORT,8080"`
	Env            string `env:"ENV,dev"`
	Namespace      string `env:"NAMESPACE,"`
	KubeConfigPath string `env:"KUBECONFIG_PATH,config.yml"`
}

func ListDefaultValues(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("env")

		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		defaultValue := parts[1]

		if defaultValue == "" {
			defaultValue = "default"
		}

		fmt.Printf("%s: %s\n", f.Name, defaultValue)
	}
}
