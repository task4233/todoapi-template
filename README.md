# todoapi-template
![CI](https://github.com/task4233/todoapi-template/workflows/CI%20for%20codecov/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/task4233/todoapi-template)](https://goreportcard.com/report/github.com/task4233/todoapi-template)
[![codecov](https://codecov.io/gh/task4233/todoapi-template/branch/master/graph/badge.svg?token=FZOJ3072P6)](https://codecov.io/gh/task4233/todoapi-template)

Web Api Template based on [Merpay Youtube live](https://www.youtube.com/watch?v=cWvAhmfZJZg).
Tests, Makefile, GitHub Actions are original.

## QuickStart
```
$ make run
./server
$ curl -X POST -d `jo title=test` localhost:8080/create
{"id":"aef968c6-ad30-45b0-aa1c-3b73eef48977","title":"test"}
$ curl localhost:8080/list | jq
[
  {
    "id": "aef968c6-ad30-45b0-aa1c-3b73eef48977",
    "title": "test"
  }
]
```

## Structure
```
.
├── .github
|   ├── codecov.yml
|   ├── cover.html
|   ├── labeler.yml
|   └── workflows
|       ├── CI.yml
|       └── labeler.yml
├── go.mod
├── go.sum
├── internal
│   ├── db
│   │   ├── db.go
│   │   ├── memory.go
│   │   └── memory_test.go
│   ├── http
│   │   ├── handler.go
│   │   ├── handler_test.go
│   │   └── server.go
│   ├── server
│   │   ├── server.go
│   │   └── server_test.go
│   └── todo
│       ├── todo.go
│       └── todo_test.go
├── main.go
└── Makefile
```

## Makefile Commands
```
build: build gocode

run: run gocode with $(ENV_FILE)

test: test gocode with $(ENV_TEST_FILE)

test-with-coverage: test gocode with coverage with $(ENV_TEST_FILE)

docker-build: build docker image

docker-run: run docker container

docker-exec: exec in docker container

cloudrun-build: building for cloudrun

cloudrun-deploy: deploying into cloudrun
```
