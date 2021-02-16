BINARY=ad-storage-service
test: 
	go test -v -cover -covermode=atomic ./...

ad:
	go build -o ${BINARY} cmd/ad/storage/main.go


unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

remove:
	docker container rm postgres

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint