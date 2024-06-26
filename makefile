include .env

run:
	go run .

build:
	GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

build-image:
	docker build -f ./deploy/Dockerfile -t agusheryanto182/go-health-record . --no-cache

run-image:
	docker run -e DB_NAME=$(DB_NAME)  -e DB_PORT=$(DB_PORT) -e DB_HOST=$(DB_HOST) \
	-e DB_USERNAME=$(DB_USERNAME) \
	-e DB_PASSWORD=$(DB_PASSWORD) \
	-e DB_PARAMS=$(DB_PARAMS) \
	-e JWT_SECRET=$(JWT_SECRET) \
	-e BCRYPT_SALT=$(BCRYPT_SALT) \
	--network health-record \
	-p 8080:8080 agusheryanto182/go-health-record

migrate:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=verify-full&rootcert=ap-southeast-1-bundle.pem" -path db/migrations -verbose up

rollback:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=verify-full&rootcert=ap-southeast-1-bundle.pem" -path db/migrations -verbose drop

migrate-dev:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?$(DB_PARAMS)" -path db/migrations -verbose up

rollback-dev:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?$(DB_PARAMS)" -path db/migrations -verbose down

drop-dev:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?$(DB_PARAMS)" -path db/migrations -verbose drop

gen-swagger:
	swag init -g cmd/main.go -output cmd/docs

db-up:
	docker compose up -d

db-down:
	docker compose down