package config

import (
	"reflect"

	"github.com/sayed-imran/dynamic-proxies/handlers"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

func KubeClient(env string) *kubernetes.Clientset {
	if env == "dev" {
		baseConfig, err := LoadConfig()
		handlers.ErrorHandler(err, "Error loading config")
		config, err := clientcmd.BuildConfigFromFlags("", baseConfig.KubeconfigPath)
		handlers.ErrorHandler(err, "Error building kubeconfig")
		clientset, err := kubernetes.NewForConfig(config)
		handlers.ErrorHandler(err, "Error building clientset")
		return clientset
	} else {
		var config *rest.Config
		config, err := rest.InClusterConfig()
		handlers.ErrorHandler(err, "Error building kubeconfig")
		clientset, err := kubernetes.NewForConfig(config)
		handlers.ErrorHandler(err, "Error building clientset")
		return clientset
	}
}
