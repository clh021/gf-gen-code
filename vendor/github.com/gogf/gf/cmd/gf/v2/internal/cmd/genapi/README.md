## 本目录操作细节

```bash
cp -r gendao genapi
cd genapi
sed -i'' 's/package gendao/package genapi/' *.go
sed -i'' 's/gf gen dao/gf gen api/' *.go
# 参考 gendao.go 中的 Dao 函数，编写了 genapi.go 中的 Api 函数
# 参考 gendao_entity.go 文件，编写了 genapi_api.go 文件
```
