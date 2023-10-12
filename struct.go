package main

type Parent struct {
	Id    string `json:"appId" yaml:"appId"`
	Slice []Child
}

type Child struct {
	Src              string   `json:"src" yaml:"src"`
	Dst              string   `json:"dst" yaml:"dst"`
	DstVersion       string   `json:"dstVersion" yaml:"dstVersion"`
	Routes           []string `json:"routes" yaml:"routes"`
	DeployDependency bool     `json:"deployDependency" yaml:"deployDependency"`
}
