package run

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"thanos/app/proto"

	"google.golang.org/grpc"
)

type server struct {
	secret string
	proto.UnimplementedCommunicateServer
}

type file struct {
	name    string
	content []byte
}

func (s *server) Apply(ctx context.Context, in *proto.ApplyRequest) (*proto.StatusResponse, error) {
	if in.Secret != s.secret {
		return nil, errors.New("error with secret")
	}
	dir, _ := ioutil.TempDir(".thanos_files", "*")
	defer os.RemoveAll(dir)

	files := []file{
		{
			name:    "deployment.json",
			content: in.Deployment,
		},
		{
			name:    "service.json",
			content: in.Service,
		},
		{
			name:    "ingress.json",
			content: in.Ingress,
		},
	}

	for _, f := range files {
		filename := fmt.Sprintf("%v/%v", dir, f.name)
		ioutil.WriteFile(filename, f.content, 0755)
		log.Printf("-> Created %v\n", filename)
		log.Printf("-> Apply %v\n", filename)
		exec.Command("kubectl", "apply", "-f", filename).Run()
	}

	return &proto.StatusResponse{Success: true}, nil
}

func RunServer(address, secret string) {
	os.Mkdir(".thanos_files", 0775)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCommunicateServer(s, &server{secret: secret})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
