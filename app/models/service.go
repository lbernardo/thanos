package models

import "fmt"

func NewService(appname string, port int64) *Service {
	return &Service{
		Kind:       "Service",
		APIVersion: "v1",
		Metadata: Metadata{
			Name:      fmt.Sprintf("service-%v", appname),
			Namespace: "default",
		},
		Spec: Spec{
			Selector: ServiceSelector{
				App: appname,
			},
			Ports: []ServicePort{
				ServicePort{
					Name:       "web",
					Port:       port,
					TargetPort: port,
				},
			},
		},
	}
}

type Service struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type Spec struct {
	Ports    []ServicePort   `json:"ports"`
	Selector ServiceSelector `json:"selector"`
}

type ServicePort struct {
	Name       string `json:"name"`
	Port       int64  `json:"port"`
	TargetPort int64  `json:"targetPort"`
}

type ServiceSelector struct {
	App string `json:"app"`
}
