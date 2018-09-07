# Go パラメータ
GOOS=linux
GOARCH=amd64
GOCMD=go


run_local_agent:
	docker-compose up --build api etcd
	go run cmd/agent/main.go

build:
	go build cmd/agent/*.go -o bin/agent -v -x

dep:
	dep ensure
	dep status

dep_update:
	dep ensure -update
	dep prune
	dep status

analysis:
	gofmt -d -s `find ./ -name "*.go" | grep -v vendor`
	golint ./... | grep -v vendor # https://github.com/golang/lint/issues/320

test-small:
	go vet ./...
	go test -v ./...

test-medium: # with root, having dependency for outside
	go test -tags=medium -v ./...

clean:
	go clean
	rm -rf bin
