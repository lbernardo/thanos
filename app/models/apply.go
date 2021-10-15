package models

type ApplyFile struct {
	Name         string                 `json:"name"`
	Environments map[string]Environment `json:"environments"`
}

type Environment struct {
	Image    string       `json:"image"`
	Replicas int64        `json:"replicas"`
	Service  ServiceApply `json:"service"`
}

type ServiceApply struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}
