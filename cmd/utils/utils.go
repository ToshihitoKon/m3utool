package utils

import (
	"io/ioutil"
	"log"
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
