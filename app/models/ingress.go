package models

import "fmt"

func NewIngress(appname, host string) *Ingress {
	return &Ingress{
		APIVersion: "networking.k8s.io/v1",
		Kind:       "Ingress",
		Metadata: IngressMetadata{
			Name: fmt.Sprintf("%v-ingress-controller", appname),
			Annotations: Annotations{
				KubernetesIoIngressClass: "traefik",
			},
		},
		Spec: IngressSpec{
			Rules: []Rule{
				Rule{
					Host: host,
					HTTP: HTTP{
						Paths: []Path{
							Path{
								Path:     "/",
								PathType: "Exact",
								Backend: Backend{
									Service: ServiceBackend{
										Name: appname,
										Port: PortServiceBackend{
											Name: "http",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

type Ingress struct {
	Kind       string          `json:"kind"`
	APIVersion string          `json:"apiVersion"`
	Metadata   IngressMetadata `json:"metadata"`
	Spec       IngressSpec     `json:"spec"`
}

type IngressMetadata struct {
	Name              string      `json:"name"`
	CreationTimestamp interface{} `json:"creationTimestamp"`
	Annotations       Annotations `json:"annotations"`
}

type Annotations struct {
	KubernetesIoIngressClass string `json:"kubernetes.io/ingress.class"`
}

type IngressSpec struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Host string `json:"host"`
	HTTP HTTP   `json:"http"`
}

type HTTP struct {
	Paths []Path `json:"paths"`
}

type Path struct {
	Path     string  `json:"path"`
	PathType string  `json:"pathType"`
	Backend  Backend `json:"backend"`
}

type Backend struct {
	Service ServiceBackend `json:"service"`
}

type ServiceBackend struct {
	Name string             `json:"name"`
	Port PortServiceBackend `json:"port"`
}

type PortServiceBackend struct {
	Name string `json:"name"`
}
