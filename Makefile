.DEFAULT_GOAL = help

MAIN_API = "cmd/api/main.go"
MAIN_API_BIN = "bin/api"

ent-init: ## Inits Ent schemas.
	go run -mod=mod entgo.io/ent/cmd/ent new \
	--target entity/schema \
	$(schema)

go-generate: ## Generates all scenarios.
	go generate ./...

encore-conn-uri: ## Get encore-managed DB URI.
	encore db conn-uri -v billing

help: ## Prints this message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build help ent-init encore-conn-uri encore-run
