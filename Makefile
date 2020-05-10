build-server:
	docker-compose build

build-client:
	mkdir ./bin/; cd ./client; go build -o ../bin/client;

run-server:
	docker-compose up;

run-client: build-client
	./bin/client
	
