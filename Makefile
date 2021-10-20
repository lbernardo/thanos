build-proto:
	protoc --go_out=. --go-grpc_out=. proto/communicate.proto
build:
	go build -o bin/thanos .
install:	
	./resources/helm/linux-amd64/helm repo add traefik https://helm.traefik.io/traefik
	./resources/helm/linux-amd64/helm repo update
	./resources/helm/linux-amd64/helm install traefik traefik/traefik
build-with-docker:
	docker run --rm -v $PWD:/var/app -w /var/app golang:1.16 make build