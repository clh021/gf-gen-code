# 方案一说明

本方案是考虑 `fork` 一份 官方 `gf/cmd/gf` 代码，来维护一个定制化的 `gf` 命令行工具。

通过扩展`crud`命令，调用内置代码的同时，增加自定义逻辑。

考虑到 `crud` 的实际需要并非只是 修改下 `ctrl` 增加一个 `logic` 那么简单。

应该要根据 数据库表，生成 `crud` 的 `api` 然后，生成后续一系列操作。

所以，此方案暂时搁置，考虑单独维护一个工具项目来实现定制化代码生成的需求。
新工具仍然会尽可能调用 `gf/cmd/gf` 中的代码，来完成后续一系列操作。