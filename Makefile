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
	@rm -f service/tpl/build_pack_data.go
	gf pack service/tpl/gen_templates service/tpl/build_pack_data.go -p gen_templates
	@docker run -it --rm -v `pwd`:/app -w /app -e CGO_ENABLED=1 -u ${UID}:${GID} leehom/detect:centos7.go1.19 go build -mod vendor -ldflags "-s -w -X github.com/clh021/gf-gen-code/cmd/v1/cmd.BuiltGit=${gitCID} -X github.com/clh021/gf-gen-code/cmd/v1/cmd.BuiltTime=${gitTime}" -o tmp/gf_gen cmd/v1/*.go
	@rm -f service/tpl/build_pack_data.go
#	cd tmp/; zip -r -q "gf_gen.${fileTime}.zip" gf_gen

.PHONY: format
# go install golang.org/x/tools/cmd/goimports@latest
format:
	find . -name "*.go" -not -path "./vendor/*" | xargs goimports -w

.PHONY: gen
gen:
	go generate ./...

.PHONY: preproj
preproj:
	@./scripts/preproj.sh

.PHONY: test
test:
	@./scripts/test.sh

.PHONY: h
h:
	./tmp/gf_gen -h