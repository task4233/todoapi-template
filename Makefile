CLOUDRUN_APP_NAME=todoapi-template
DOCKER_IMAGE_NAME=todoapi-template
DOCKER_CONTAINER_NAME=todoapi-template

ENV_FILE := .env
ENV := $(shell cat $(ENV_FILE))

ENV_TEST_FILE := .env.test
ENV_TEST := $(shell cat $(ENV_TEST_FILE))

.PHONY: build
build:
	go build -o server

run: build
	$(ENV) ./server

test:
	$(ENV_TEST) go test -covermode=atomic -count=1 ./...

test-with-coverage:
	$(ENV_TEST) go test -covermode=atomic -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o ./.github/cover.html

.PHONY: docker-build
docker-build:
	sudo docker build . -t ${DOCKER_IMAGE_NAME}

docker-run: docker-build
	sudo docker run -it --rm -p ${PORT}:${PORT} --env-file ${ENV_FILE} --name ${DOCKER_CONTAINER_NAME} ${DOCKER_IMAGE_NAME}

docker-exec:
	sudo docker exec -it ${DOCKER_CONTAINER_NAME} /bin/sh

cloudrun-build:
	gcloud builds submit --tag gcr.io/`gcloud config get-value project`/${CLOUDRUN_APP_NAME}

cloudrun-deploy:
	gcloud run deploy --image gcr.io/`gcloud config get-value project`/${CLOUDRUN_APP_NAME} --platform managed