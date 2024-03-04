# gf-gen-code
gen code  for goframe

## SQLite 注释的识别

[识别逻辑](https://github.com/clh021/gf-gen-code/commit/54d110e663c8eedd51f40c750be10726b2dfdb65)

[sql示例](https://github.com/clh021/gf-gen-code/blob/main/service/db/test.sql)


## 如果依赖 gf 的 internal 代码怎么办？
通过 go:generate command 获取代码和更新代码
`go build -o source-analysis gitee.com/source-analysis/testkit/cmd/source-analysis`
```go
//go:generate go run my_generator.go
//go:generate ./generate_code.sh arg1 arg2
go generate -run MyGenerator // 将只执行那些注释中包含 MyGenerator 字符串的命令
```

## TODO

- 支持自定义模板(有内置模板，同时支持识别项目目录下的 gen_templates 作为自定义模板目录)
- 支持 init 命令，生成项目，生成数据库和配置，生成 apidoc 页面(后期可不断完善初始化项目)，使可直接运行起来
- api 生成api
- logic 生成 logic(自动修改 ctrl )
- api接口
- 搜索带翻页
- 获取一个
- 增加一个
- 修改一个字段
- 修改多个字段
- 删除一个
- 增加多个
- 文件上传
- 文件下载
- 文件压缩下载
- 文件信息获取