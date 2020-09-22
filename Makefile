default: build

app = ics
src = main.go internal
ver = $(shell date +"%Y%m%d%H%M%S")

test: ${src}
	@gofmt -w -s ${src}
	@goimports -w ${src}
	@go vet ./...
	@go test ./...
# 	@golint -set_exit_status ./...

build: test
	@go build -o ${app} -ldflags "-X 'main.Version=${ver}'"
