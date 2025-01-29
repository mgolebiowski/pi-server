.PHONY: run
run: env
	go run cmd/pi-server/main.go

.PHONY: env
env:
	@export $(cat .env | xargs) > /dev/null