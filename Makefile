.PHONY: build

build: ## Build docker image
	docker-compose build

gen:
	protoc \
	--go_out=service/common/pb \
	--go_opt=paths=source_relative \
	--go-grpc_out=service/common/pb \
	--go-grpc_opt=paths=source_relative \
	--proto_path=service/common/protofiles \
	service/common/protofiles/*.proto

.PHONY: status logs start stop clean

status: ## Get status of containers
	docker-compose ps

logs: ## Get logs of containers
	docker-compose logs -f

start: ## Start docker containers
	docker-compose up -d

stop: ## Stop docker containers
	docker-compose stop

clean:stop ## Stop docker containers, clean data and workspace
	docker-compose down -v --remove-orphans

prune:
	docker system prune

.PHONY: test

test: ## Run tests
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes
