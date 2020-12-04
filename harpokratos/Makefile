build:
	GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -o main cmd/main.go
	mkdir ./bin
	mv main ./bin/harpokratos

docker-build:
	docker build -t harpokratos:latest .