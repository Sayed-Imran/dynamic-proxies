package config

type BaseConfig struct {
	Port           string `env:"PORT,8080"`
	Env            string `env:"ENV,dev"`
	Namespace      string `env:"NAMESPACE,"`
	KubeConfigPath string `env:"KUBECONFIG_PATH,config.yml"`
}
