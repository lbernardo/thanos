package models

import "fmt"

func NewIngress(appname, host string) *Ingress {
	return &Ingress{
		APIVersion: "extensions/v1beta1",
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
								Path: "/",
								Backend: Backend{
									ServiceName: appname,
									ServicePort: "http",
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
	Path    string  `json:"path"`
	Backend Backend `json:"backend"`
}

type Backend struct {
	ServiceName string `json:"serviceName"`
	ServicePort string `json:"servicePort"`
}
