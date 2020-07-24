compose-up: ## Starts a local instance of the service with instances of the event bus and ArangoDB in docker
	@if [ -f docker-compose.yml ]; then\
		docker-compose -p ent_graphql -f docker-compose.yml up -d ;\
	else\
		echo "No compose file.";\
	fi

compose-down: ## Shutdowns the local containers
	@docker-compose -p ent_graphql -f docker-compose.yml down

entc:
	entc generate ./internal/provider/entc/schema

gqlgen:
	gqlgen generate --config ./config/gqlgen.yml

sync:
	make entc
	make gqlgen