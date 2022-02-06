package service

import (
	"encoding/json"
	"filestoreapp/utils"
	"io"
	"net/http"
	"os"
)

func StoreFile(res http.ResponseWriter, req *http.Request) {
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
func UpdateFile(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(15 * 1024 * 1024)

	file, fileheader, err := req.FormFile("file")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileLists, err := utils.FileList(utils.BasePath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var osfile *os.File
	defer osfile.Close()
	if utils.StringInSlice(fileheader.Filename, fileLists) {
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
func ListFiles(res http.ResponseWriter, req *http.Request) {
	fileLists, err := utils.FileList(utils.BasePath)
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

func DeleteFile(res http.ResponseWriter, req *http.Request) {
	filename := req.URL.Query().Get("filename")
	err := os.Remove(utils.BasePath + filename)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write([]byte(filename))

}

func Readynesscheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
