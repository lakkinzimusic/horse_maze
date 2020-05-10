PROJECT?=horse_maze
APP?=horsemaze
PORT?=8000

RELEASE?=0.0.2
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=docker.io/lakkinzimusic/${APP}


GOOS?=linux
GOARCH?=amd64

commit: 
	git add .
	git commit -m "message"
	git push "https://$(shell git config user.name):$(shell git config user.password)@github.com/lakkinzimusic/horse_maze.git"

clean: commit
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
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

test:
	go test -v -race ./...

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

minikube: push
	for t in $(shell find ./ -type f -name "*.yaml"); do \
        cat $$t | \
        	sed -E "s/\{\{(\s*)\.Release(\s*)\}\}/$(RELEASE)/g" | \
        	sed -E "s/\{\{(\s*)\.ServiceName(\s*)\}\}/$(APP)/g"; \
        echo; \
    done > tmp.yaml
	kubectl apply -f tmp.yaml