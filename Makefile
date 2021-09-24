BINARY=rps-web

build:
	@-go build -o dist/${BINARY} .
	@echo "Binary ${BINARY} built in ./dist"

start: build
	@-./dist/rps-web &
	@echo "Application started on port 8080"


stop:
	@-pkill -SIGTERM -f "./dist/${BINARY}" 
	@echo "Application stopped"