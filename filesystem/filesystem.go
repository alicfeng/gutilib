package filesystem

import "os"

// PathExist 路径是否存在 涵盖文件与文件夹
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
