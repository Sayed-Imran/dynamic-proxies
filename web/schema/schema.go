package apischema

type DeployConfig struct {
	AppName  string `json:"name"`
	Image    string `json:"image"`
	Replicas int    `json:"replicas"`
	Port     int    `json:"port"`
}

type DeleteConfig struct {
	AppName string `json:"name"`
}

