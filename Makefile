ENV?=dev
AWS_ACCOUNT=${AWS_ACCOUNT_ENV}
AWS_PROFILE:=${AWS_PROFILE_ENV}
AWS_REGION?=ap-northeast-1
AWS_REPOSITORY:=modev-backend
TAG:=$(shell git rev-parse --short HEAD)
PWD:=$(shell pwd)
DATABASE=mysql://root@tcp(mysql:3306)/modev?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true


.PHONY: $(shell egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')
setup:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go mod tidy

gen:
	oapi-codegen -config ./gen/model.config.yaml openapi.yaml
	oapi-codegen -config ./gen/server.config.yaml openapi.yaml

migrate-create:
	docker run \
		-v ${PWD}/mysql/migrate:/migrations \
		migrate/migrate \
		create -ext sql -dir /migrations -seq initialize

migrate-up:
	docker run \
		-v ${PWD}/mysql/migrate:/migrations \
		--network modev-backend-network \
		migrate/migrate \
		-path=/migrations/ \
		-database "${DATABASE}" \
		up

migrate-drop:
	docker run \
		--network modev-backend-network \
		migrate/migrate \
		-database "${DATABASE}" \
		drop -f

swagger:
	docker run --rm --name openapi -d \
		-p 8081:8080 \
		-v $(PWD):/tmp \
		-e SWAGGER_FILE=/tmp/openapi.yaml \
		--platform=linux/amd64 \
		--name swagger \
		swaggerapi/swagger-editor
	open http://localhost:8081

build:
	docker build -t $(AWS_REPOSITORY) . --platform=linux/amd64

tag:
	echo $(TAG)
	docker tag $(AWS_REPOSITORY) $(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(AWS_REPOSITORY):$(TAG)
	docker tag $(AWS_REPOSITORY) $(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(AWS_REPOSITORY):latest

login:
	aws ecr get-login-password --region $(AWS_REGION) --profile $(AWS_PROFILE) | docker login --username AWS --password-stdin $(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(AWS_REPOSITORY)

push:
	docker push $(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(AWS_REPOSITORY):$(TAG)
	docker push $(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(AWS_REPOSITORY):latest

run:
	docker run -it $(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(AWS_REPOSITORY):$(TAG)
