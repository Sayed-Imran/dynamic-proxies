package config

import (
	"reflect"

	"github.com/spf13/viper"
)

type BaseConfig struct {
	Port           string `mapstructure:"PORT" default:"8080"`
	Env            string `mapstructure:"ENV"`
	Namespace      string `mapstructure:"NAMESPACE" default:"default"`
	KubeconfigPath string `mapstructure:"KUBECONFIG_PATH"`
	NodeAffinity   string `mapstructure:"NODE_AFFINITY"`
}

func LoadConfig() (config BaseConfig, err error) {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return
	}
	v := reflect.ValueOf(&config).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			defaultTag := t.Field(i).Tag.Get("default")
			field.SetString(defaultTag)
		}
	}
	return
}
