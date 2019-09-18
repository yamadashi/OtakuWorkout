package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]

	rawdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	data := string(rawdata)
	separator := "~~~"
	splitContent := strings.SplitN(data, separator, 2)
	if len(splitContent) < 1 {
		return
	}

	fmt.Println(strings.Trim(splitContent[0], "\n 	"))

	var newContent []byte = nil
	if len(splitContent) == 2 {
		newContent = []byte(splitContent[1])
	}
	err = ioutil.WriteFile(path, newContent, 0664)
	if err != nil {
		fmt.Println(err)
	}
}
