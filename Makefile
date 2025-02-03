.PHONY: run
run: env
	go run cmd/pi-server/main.go

.PHONY: env
env:
	@export $(cat .env | xargs) > /dev/null

.PHONY: rpi_build
rpi_build:
	GOOS=linux GOARCH=arm go build cmd/pi-server/main.go
