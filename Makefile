init:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && \
	go get -v github.com/swaggo/swag/cmd/swag && \
	go get -v github.com/swaggo/echo-swagger && \
	go mod vendor

run:
	swag init && \
	go run main.go