git_short_hash=`git rev-parse --short HEAD`

build:
	GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -o harpokratos cmd/main.go
	mkdir ./bin
	mv harpokratos ./bin/

docker-build:
	docker build -t harpokratos:$(git_short_hash) .

docker-build-stack: docker-build
	cd ../samplestack/golang-example && docker build -t golangexampleapi:$(git_short_hash) .
	cd ../..