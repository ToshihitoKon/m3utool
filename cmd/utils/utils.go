package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	rawPaths := strings.Split(string(b), "\n")
	paths := []string{}
	for _, path := range rawPaths {
		if len(path) == 0 {
			continue
		}
		paths = append(paths, path)
	}
	return paths
}

func CheckFileExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		// 見つからない場合ここ
		// なんか重いエラーはFatalにしたい気持ちもある
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}
