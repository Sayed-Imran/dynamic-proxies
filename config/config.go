package config

import (
	"reflect"

	errorHandler "github.com/sayed-imran/dynamic-proxies/handlers/error_handler"
	"github.com/spf13/viper"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type BaseConfig struct {
	Port           string `mapstructure:"PORT" default:"8080"`
	Env            string `mapstructure:"ENV" default:"incluster"`
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

type kubenetesConfig struct {
	Clientset     *kubernetes.Clientset
	DynamicClient dynamic.Interface
}

func KubeClients() *kubenetesConfig {
	if Configuration.Env == "dev" {
		kubeConfig, err := clientcmd.BuildConfigFromFlags("", Configuration.KubeconfigPath)
		errorHandler.ErrorHandler(err, "Error building kubeconfig")
		ClientSet, err := kubernetes.NewForConfig(kubeConfig)
		errorHandler.ErrorHandler(err, "Error building clientset")
		DynamicClient, err := dynamic.NewForConfig(kubeConfig)
		errorHandler.ErrorHandler(err, "Error building dynamic client")
		return &kubenetesConfig{
			Clientset:     ClientSet,
			DynamicClient: DynamicClient,
		}
	} else {
		restConfig, err := rest.InClusterConfig()
		errorHandler.ErrorHandler(err, "Error building inClusterConfig")
		ClientSet, err := kubernetes.NewForConfig(restConfig)
		errorHandler.ErrorHandler(err, "Error building clientset")
		DynamicClient, err := dynamic.NewForConfig(restConfig)
		errorHandler.ErrorHandler(err, "Error building dynamic client")
		return &kubenetesConfig{
			Clientset:     ClientSet,
			DynamicClient: DynamicClient,
		}
	}

}

var Configuration, _ = LoadConfig()
var KubeClient = KubeClients()
