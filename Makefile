export PORT = ${HISTORY_SERVER_LISTEN_ADDR}
build:
	go build 

run:
	go run main.go

test:
	go test  -cover ./...

image:
	 docker build --no-cache -t location-app .

run-docker: image
run-docker: check-env
	@echo Running on port ${PORT}
	docker run -e HISTORY_SERVER_LISTEN_ADDR=${PORT} -p ${PORT}:${PORT} -it location-app

check-env:
ifeq ($(PORT),)
PORT := 8080
endif