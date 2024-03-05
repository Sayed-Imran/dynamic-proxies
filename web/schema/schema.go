package apischema

type DeployConfig struct {
	AppName  string `json:"name"`
	Image    string `json:"image"`
	Replicas int32    `json:"replicas"`
	Port     int32    `json:"port"`
}

type DeleteConfig struct {
	AppName string `json:"name"`
}

