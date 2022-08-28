APP_NAME=user-auth-service

build: dep
	env CGO_ENABLED=0 GOARCH=amd64 go build -ldflags -installsuffix -o ${APP_NAME} github.com/mariojuzar/go-user-auth

dep:
	@echo "# Downloading Dependencies"
	@go mod download

run-api: dep
	@echo ">> Running API Server"
	@env $$(cat .env | xargs) go run github.com/mariojuzar/go-user-auth server

migrate: dep
	@echo ">> Running API Server"
	@env $$(cat .env | xargs) go run github.com/mariojuzar/go-user-auth migrate
