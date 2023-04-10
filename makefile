registry:=registry.hub.docker.com
username:=ahmadfajarislami
image:=go_todo_list
tags:=latest


.PHONY: help
help: ## Show help command
	@printf "Makefile Command\n";
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: check-swag
check-swag: ## Check if swag command already exist
	command -v swag >/dev/null 2>&1 || { go install github.com/swaggo/swag/cmd/swag@v1.8.10; }

.PHONY: docs 
docs: ## Generate Documents
	swag init -g ./internal/delivery/http/main.go --output ./docs/

.PHONY: readenv
readenv: ## not work
	export $(cat .env | xargs -L 1)

.PHONY: migrate
migrate: ## Create Migrations file, example : make migrate name="xxxx"
	@if [ -z "${name}" ]; then \
		echo "Error: name is required \t example : make migrate name="name_file_migration";" \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations ':hammer: ${name}'


migrate-up: ## Up migration, example : make migrate-up envfile=.env.test
	go run cmd/migrate/main.go -envfile=${envfile}

migrate-rollback: ## Up rollback, example : make migrate-rollback -rollback=true envfile=.env.test
	go run cmd/migrate/main.go -rollback -envfile=${envfile}

build:
	go build -o app cmd/main.go

run:
	./app

run.local:
	go run cmd/main.go

run.env:
	docker compose up -d mysql_go_todo_list

docker.build:
	docker build --rm -t ${registry}/${username}/${image}:${tags} .
	docker image prune --filter label=stage=dockerbuilder -f

docker.run:
	docker run --name ${image} -p 8080:8080 ${registry}/${username}/${image}:${tags}

dc.up: ## up compose image
	docker compose -f docker-compose.yaml up -d

dc.logs: ## logs compose image
	docker compose -f docker-compose.yaml logs -f

dc.stop: ## stop compose image
	docker compose -f docker-compose.yaml stop

dc.down: ## rm compose image
	docker compose -f docker-compose.yaml down -v


docker.rm:
	docker rm ${registry}/${username}/${image}:${tags} -f
	docker rmi ${registry}/${username}/${image}:${tags}

docker.enter:
	docker exec -it ${image} bash

docker.enterimg:
	docker run -it --entrypoint sh  ${registry}/${username}/${image}:${tags}

dc.check:
	 docker compose -f docker-compose.yaml config
	 
push-image: dockerbuild
	docker push ${registry}/${username}/${image}:${tags}

flysecret:
	flyctl secrets set $(cat .env | xargs)

flylist:
	flyctl secrets list


entermysql:
	docker exec -it mysql_go_todo_list mysql -uADMIN -pSECRET todo4

test:
	go test -run=TestHandlerUsers ./internal/delivery/http/handler -v -count=1 --cover