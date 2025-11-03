# two command one for run the main.go and one for build and put in the bin folder with the version
 build:
	go build -o ./bin/dotsync -ldflags "-X main.version=`git describe --tags --always`" ./cmd/dotsync/main.go

 run:
	go run ./cmd/dotsync/main.go

.PHONY: build run