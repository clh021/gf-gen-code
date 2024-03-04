package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
)

func IsDevelop() bool {
	return gfile.IsFile("./hack/config.yaml")
}

func SureAllFilesExist(ctx context.Context, fileNames []string) bool {
	if !IsDevelop() {
		for _, fileName := range fileNames {
			fullPath := filepath.Clean(fileName)
			if !gfile.Exists(fullPath) {
				glog.Printf(ctx, "Failed to find required file: %s", fullPath)
				glog.Printf(ctx, "You can get file `%s` by Command: %s", fullPath, "compat_detect_tool init")
				os.Exit(0)
			}
		}
	}
	return true
}

func autoGenerateFile(ctx context.Context, filePath, tplPath string) {
	sDir := gfile.SelfDir()
	fFile := gfile.Join(sDir, filePath)

	// 检查文件是否存在
	if !gfile.IsFile(fFile) {
		efile := gres.Get(tplPath)
		if efile == nil {
			glog.Printf(ctx, "Failed to get %s, unable to generating: %s", tplPath, fFile)
			return
		}
		err := gfile.PutBytes(fFile, efile.Content())
		if err != nil {
			glog.Fatalf(ctx, "Error occurred while generating %s: %v", fFile, err)
			return
		}
		glog.Printf(ctx, "Generated: %s", fFile)
	} else {
		glog.Printf(ctx, "Exists: %s     --skip", fFile)
	}
}
