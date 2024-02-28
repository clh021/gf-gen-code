# 获取当前用户的 UID 和 GID
UID=$(shell id -u)
GID=$(shell id -g)
fileTime=$(shell date +%Y%m%d%H%M)

gitTime=$(shell date +"%Y%m%d.%H%M%S")
gitCID=$(shell git rev-parse HEAD | cut -c1-15)
gitTag=$(shell git tag --list --sort=version:refname 'v*' | tail -1)
gitCount=$(shell git log --pretty=format:'' | wc -l)/$(shell git rev-list --all --count)
buildStr=${gitTime}.${gitCID}.${gitTag}.${gitCount}
# CGO_ENABLED=1: go-sqlite3 requires cgo to work

# go mod vendor
# go mod tidy
# go generate
# CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -ldflags "-X main.build=${buildStr}" -o tmp/app cmd/main.go
.PHONY: build
build:
	@docker run -it --rm -v `pwd`:/app -w /app -e CGO_ENABLED=1 -u ${UID}:${GID} leehom/detect:centos7.go1.19 go build -mod vendor -ldflags "-s -w -X github.com/clh021/gf-gen-code/cmd/v1/cmd.BuiltGit=${gitCID} -X github.com/clh021/gf-gen-code/cmd/v1/cmd.BuiltTime=${gitTime}" -o tmp/gf_gen cmd/v1/*.go
#	cd tmp/; zip -r -q "gf_gen.${fileTime}.zip" gf_gen

.PHONY: gen
gen:
	go generate ./...

.PHONY: test
test:
	./tmp/gf_gen -c ./tmp/config.yaml

.PHONY: h
h:
	./tmp/gf_gen -h