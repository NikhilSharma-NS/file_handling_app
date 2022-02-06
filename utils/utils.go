package utils

import (
	"bufio"
	"io"
	"io/ioutil"
)

var (
	BasePath string = "./store/"
)

func GetFileList(path string) (fileList []string, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fileList, err
	}
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList, nil
}

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func WordCount(rdr io.Reader) int {
	counts := 0
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts++
	}
	return counts
}
