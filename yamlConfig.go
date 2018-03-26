package main

type yamlConfig struct {
	Descriptions []description `yaml:"descriptions"`
}

type description struct {
	ExecutionEnabled bool      `yaml:"executionEnabled"`
	Commands         []command `yaml:"commands"`
}

type command struct {
	Configuration configuration `yaml:"configuration"`
	RemoteUrl     string        `yaml:"remoteUrl"`
}

type configuration struct {
	Command           string `yaml:"command"`
	RemoteUrl         string `yaml:"remoteUrl"`
	CheckResponseCode string `yaml:"checkResponseCode"`
	HttpMethod        string `yaml:"method"`
}
