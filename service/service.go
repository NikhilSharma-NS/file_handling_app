package service

import (
	"encoding/json"
	"filestoreapp/utils"
	"io"
	"net/http"
	"os"
	"strconv"
)

// store the file into file store
func StoreFileHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(15 * 1024 * 1024)

	file, fileheader, err := req.FormFile("file")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	f, err := os.OpenFile(utils.BasePath+fileheader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	defer f.Close()
	io.Copy(f, file)
	res.Write([]byte(fileheader.Filename))

}

// update the file into file store
func UpdateFileHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(15 * 1024 * 1024)

	file, fileheader, err := req.FormFile("file")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileLists, err := utils.GetFileList(utils.BasePath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var osfile *os.File
	defer osfile.Close()
	if utils.IsStringInSlice(fileheader.Filename, fileLists) {
		osfile, err = os.OpenFile(utils.BasePath+fileheader.Filename, os.O_RDWR|os.O_APPEND, 0666)
	} else {
		osfile, err = os.OpenFile(utils.BasePath+fileheader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(osfile, file)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write([]byte(fileheader.Filename))
}

//  Get all  the files from file store
func ListFilesHandler(res http.ResponseWriter, req *http.Request) {
	fileLists, err := utils.GetFileList(utils.BasePath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(fileLists)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write(js)
}

// Delete the file from file store
func DeleteFileHandler(res http.ResponseWriter, req *http.Request) {
	filename := req.URL.Query().Get("filename")
	err := os.Remove(utils.BasePath + filename)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write([]byte(filename))

}

// find word count of all files from filestore
func FindWordCountHandler(res http.ResponseWriter, req *http.Request) {
	wordcount := 0
	fileLists, err := utils.GetFileList(utils.BasePath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, file := range fileLists {
		file, err := os.Open(utils.BasePath + file)
		if err != nil {
			continue
		}
		counts := utils.WordCount(file)
		defer file.Close()
		wordcount = wordcount + counts
	}
	res.Write([]byte(strconv.Itoa(wordcount)))
}

// check readyness
func CheckReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// check health
func CheckHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
