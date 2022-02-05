package main

import (
	"filestoreapp/controller"
	"net/http"
)

func main() {
	r := controller.FileAppRouter()
	http.ListenAndServe(":8080", r)
}
