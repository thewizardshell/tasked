.PHONY: run build

run:
	go run .\cmd\main.go

build:
	go build -o bin\tasked.exe .\cmd\main.go

sqlc:
	$(USERPROFILE)\go\bin\sqlc.exe generate

swagger:
	$(USERPROFILE)\go\bin\swag.exe init -g cmd/main.go
