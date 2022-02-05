package utils

import "io/ioutil"

var (
	BasePath string = "./store/"
)

func FileList(path string) (fileList []string, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fileList, err
	}
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList, nil
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
