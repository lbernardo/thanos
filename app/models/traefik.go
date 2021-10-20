package models

import (
	"encoding/json"
	"io/ioutil"
)

func NewTraefikService(ip string) *TraefikService {
	return &TraefikService{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: TraefikServiceMetadata{
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "traefik",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "traefik",
				"helm.sh/chart":                "traefik-10.6.0",
			},
			Name:      "traefik",
			Namespace: "default",
		},
		Spec: TraefikServiceSpec{
			ExternalIPS: []string{
				ip,
			},
			Ports: []TraefikServicePort{
				{
					Name:       "web",
					NodePort:   31576,
					Port:       80,
					Protocol:   "TCP",
					TargetPort: "web",
				},
				{
					Name:       "websecure",
					NodePort:   31373,
					Port:       443,
					Protocol:   "TCP",
					TargetPort: "websecure",
				},
			},
			Selector: map[string]string{
				"app.kubernetes.io/instance": "traefik",
				"app.kubernetes.io/name":     "traefik",
			},
			Type: "LoadBalancer",
		},
	}
}

func (t *TraefikService) ToFile(filename string) {
	b, _ := json.Marshal(t)
	ioutil.WriteFile(filename, b, 0755)
}

type TraefikService struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   TraefikServiceMetadata `json:"metadata"`
	Spec       TraefikServiceSpec     `json:"spec"`
}

type TraefikServiceMetadata struct {
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
}

type TraefikServiceSpec struct {
	ExternalIPS []string             `json:"externalIPs"`
	Ports       []TraefikServicePort `json:"ports"`
	Selector    map[string]string    `json:"selector"`
	Type        string               `json:"type"`
}

type TraefikServicePort struct {
	Name       string `json:"name"`
	NodePort   int64  `json:"nodePort"`
	Port       int64  `json:"port"`
	Protocol   string `json:"protocol"`
	TargetPort string `json:"targetPort"`
}
