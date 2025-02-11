s:
	swag init -g cmd/main.go -o docs
r:
	go run cmd/main.go

build:
	go build -o kpi cmd/main.go

run-build:
	./kpi
