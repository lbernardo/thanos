package models

import "fmt"

func NewService(appname string, port int64) *Service {
	return &Service{
		Kind:       "Service",
		APIVersion: "v1",
		Metadata: Metadata{
			Name: fmt.Sprintf("service-%v", appname),
		},
		Spec: Spec{
			Selector: ServiceSelector{
				App: appname,
			},
			Ports: []ServicePort{
				ServicePort{
					Name: "http",
					Port: port,
				},
			},
		},
	}
}

type Service struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Spec struct {
	Ports    []ServicePort   `json:"ports"`
	Selector ServiceSelector `json:"selector"`
}

type ServicePort struct {
	Name string `json:"name"`
	Port int64  `json:"port"`
}

type ServiceSelector struct {
	App string `json:"app"`
}
