build:
	docker-compose build app-fibonacci

run:
	docker-compose up --remove-orphans app-fibonacci

test:
	go test -v ./...