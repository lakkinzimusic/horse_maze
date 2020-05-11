PROJECT?=horse_maze
APP?=horsemaze
PORT?=8050

RELEASE?=0.0.2
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=docker.io/lakkinzimusic/${APP}


GOOS?=linux
GOARCH?=amd64

protoc: 
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto

commit: 
	git add .
	git commit -m "Docker, kubernetes, clean implemented"
	git push "https://$(shell git config user.name):$(shell git config user.password)@github.com/lakkinzimusic/horse_maze.git"

clean: 
	rm -f ${APP}

build: clean
		CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --network="host" --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

test:
	go test -v -race ./...

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

minikube:  push
	kubectl delete secret mysql-secret
	kubectl create -f ./kubernetes/mysql-secret.yaml
	kubectl delete daemonsets,replicasets,services,deployments,pods,rc --all
	kubectl apply -f ./kubernetes/app-mysql-deployment.yaml
	kubectl apply -f ./kubernetes/app-mysql-service.yaml
	kubectl apply -f ./kubernetes/mysql-db-deployment.yaml
	kubectl apply -f ./kubernetes/mysql-db-pv.yaml
	kubectl apply -f ./kubernetes/mysql-db-pvc.yaml
	kubectl apply -f ./kubernetes/mysql-db-service.yaml