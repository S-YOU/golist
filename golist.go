package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

var (
	ignoreSuffixes = flag.String("exclude-suffixes", "", "")
)

func main() {
	flag.Parse()
	walk(".")
}

func walk(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if name := f.Name(); matcher(name) {
			if f.IsDir() {
				walk(path.Join(dir, name))
				continue
			}
			if path.Ext(name) == ".go" {
				fmt.Println(path.Join(dir, name))
			}
		}
	}
}

func matcher(name string) bool {
	for _, x := range strings.Split(*ignoreSuffixes, ",") {
		if strings.HasSuffix(name, x) {
			return false
		}
	}
	return true
}
