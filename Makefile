launch_args=
test_args=-coverprofile cover.out && go tool cover -func cover.out
cover_args=-cover -coverprofile=cover.out `go list ./...` && go tool cover -html=cover.out

VERSION?= $(shell git describe --match 'v[0-9]*' --tags --always)

# make tidy
tidy:
	go mod tidy

# make clean-up-mock
clean-up-mock:
	rm -rf ./mock

# make generate
generate: clean-up-mock
	go generate ./...


# make lint
lint:
	@golangci-lint run

# make coverage
coverage:
	@echo "total code coverage : "
	@go tool cover -func cover.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'

# make test
test:
ifeq (, $(shell which richgo))
	go test ./... $(test_args)
else
	richgo test ./... $(test_args)
endif

# make cover
cover:
ifeq (, $(shell which richgo))
	go test $(cover_args)
else
	richgo test $(cover_args)
endif

# make changelog VERSION=vx.x.x
changelog:
	git-chglog -o CHANGELOG.md --next-tag $(VERSION)

%:
	@:
