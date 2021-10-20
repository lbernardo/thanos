package run

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"thanos/app/models"
	"thanos/app/proto"
	"time"

	"google.golang.org/grpc"
)

func ApplyRun(address, secret, env, tag string, content []byte) {
	var applyContent models.ApplyFile
	json.Unmarshal(content, &applyContent)
	environmentContent, ok := applyContent.Environments[env]
	if !ok {
		fmt.Printf("Environment %v not found\n", env)
		os.Exit(1)
	}
	environmentContent.Image = fmt.Sprintf("%v:%v", environmentContent.Image, tag)
	deployment := models.NewDeployment(applyContent.Name, environmentContent.Image, environmentContent.Replicas, environmentContent.Service.Port)
	service := models.NewService(applyContent.Name, environmentContent.Service.Port)
	ingress := models.NewRoute(applyContent.Name, environmentContent.Service.Host, environmentContent.Service.Port)

	deploymentBytes, _ := json.Marshal(deployment)
	serviceBytes, _ := json.Marshal(service)
	ingressBytes, _ := json.Marshal(ingress)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*10))
	if err != nil {
		fmt.Printf("Error to connect thanos: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("-> Connected with server")

	c := proto.NewCommunicateClient(conn)
	res, err := c.Apply(context.Background(), &proto.ApplyRequest{
		Deployment: deploymentBytes,
		Service:    serviceBytes,
		Ingress:    ingressBytes,
		Secret:     secret,
	})
	fmt.Println("-> Sended content to server")
	if err != nil {
		fmt.Printf("could not apply: %v\n", err)
		os.Exit(1)
	}
	if !res.Success {
		fmt.Println("Error to apply proto")
		os.Exit(1)
	}
	fmt.Printf("Success apply %v\n", applyContent.Name)
}
