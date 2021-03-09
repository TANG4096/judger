package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	pos := strings.LastIndex(dir, "judger/")
	return dir[:pos+len("judger")]
}
