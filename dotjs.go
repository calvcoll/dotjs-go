package main

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	url = strings.TrimSuffix(url, "/")
	url = strings.TrimSuffix(url, "www.")
	/*fmt.Println("Wanted: " + r.URL.Path)*/

	usr, _ := user.Current()
	dir := usr.HomeDir

	if url == "favicon.ico" {
		w.WriteHeader(http.StatusNotFound) //normally untrue http.StatusOkay
	} else {
		data, err := ioutil.ReadFile(dir + "/.js/" + url)
		if err != nil {
			w.Write([]byte("// dotjs is working - no file found, or file error //"))
		} else {
			w.Write(data)
		}
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3131", nil)
}
