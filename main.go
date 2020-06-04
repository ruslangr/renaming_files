package main

import (
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	var files []string

	root := "/home/ruslan/Downloads/scrn/" //указываем папку
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	newName := []string{}
	os.Chdir(root)
	for i, _ := range files {
		newName = append(newName, "Udemy_Go_p3_"+strconv.Itoa(i)) //указываем новое имя
		if files[i] != "main.go" && files[i] != "main" && files[i] != root {
			os.Rename(files[i], newName[i])
		}
	}
}
