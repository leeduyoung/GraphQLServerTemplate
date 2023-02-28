.DEFAULT_GOAL := start

.PHONY: gqlgen-install
gqlgen-install:
	go get github.com/99designs/gqlgen

.PHONY: gqlgen
gqlgen: gqlgen-install
	cd pkg/graph && go run github.com/99designs/gqlgen generate

.PHONY: entgo-install
entgo-install:
	go get entgo.io/ent/cmd/ent

.PHONY: entgo_init
entgo_init: entgo-install
	go run entgo.io/ent/cmd/ent new --target ./internal/schema $(name)

.PHONY: entgo
entgo: entgo-install
	rm -rf ./ent && go run entgo.io/ent/cmd/ent generate --target ./ent ./internal/schema --feature sql/upsert

.PHONY: setup
setup:
	go mod tidy

.PHONY: start
start: entgo gqlgen setup
	go run server.go

.PHONY: fstart
fstart:
	go run server.go

.PHONY: mock
mock:
	mockery --dir ./internal/repository/ --all --keeptree

.PHONY: test
test: test_setup test_run test_clean

.PHONY: test_setup
test_setup:
	@echo "building the environment.."
	docker-compose -f docker/docker-compose.yml up --build -d
	@sleep 10
	@echo "environment build is done."

.PHONY: test_run
test_run:
	@echo "started run the all tests."
	go test -v -count 1 -p 1 -timeout 120s ./utils/... -cover
	go test -v -count 1 -p 1 -timeout 120s ./repository/... -cover
	@echo "all tests were completed."

.PHONY: test_clean
test_clean:
	@echo "cleaning the environment.."
	docker-compose -f docker/docker-compose.yml down
	@echo "environment cleaned up."
