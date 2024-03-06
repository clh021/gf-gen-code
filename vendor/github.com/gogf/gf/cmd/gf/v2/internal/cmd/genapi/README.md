## 本目录操作细节

```bash
cp -r gendao genapi
cd genapi
sed -i'' 's/package gendao/package genapi/' *.go
sed -i'' 's/gf gen dao/gf gen api/' *.go
```

新增 gendao_api.go