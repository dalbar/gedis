generate-grpc:
	protoc --go_out=./pkg/grpc --go_opt=paths=source_relative \
    --go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative \
    gedis.proto

REPO ?= dalbar
containerize:
	KO_DOCKER_REPO=$(REPO) ko build ./cmd/server

deploy-k8s:
	KO_DOCKER_REPO=$(REPO) ko apply -f deployment/deployment.yaml

deploy-migration-job:
	KO_DOCKER_REPO=$(REPO) ko apply -f deployment/migration.yaml