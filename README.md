# gf-gen-code
gen code  for goframe

## SQLite 注释的识别

[识别逻辑](./commit/54d110e663c8eedd51f40c750be10726b2dfdb65)

[sql示例](./blob/main/service/db/test.sql)


## 如果依赖 gf 的 internal 代码怎么办？
通过 go:generate command 获取代码和更新代码
`go build -o source-analysis gitee.com/source-analysis/testkit/cmd/source-analysis`
```go
//go:generate gofmt -w *.go
//go:generate go run my_generator.go
//go:generate ./generate_code.sh arg1 arg2
go generate -run MyGenerator // 将只执行那些注释中包含 MyGenerator 字符串的命令
```