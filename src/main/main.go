package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"time"
)

type FileInformation struct {
	Name             string    `json:"name"`
	Size             int64     `json:"size"`
	LastModification time.Time `json:"modTime"`
	Directory        bool      `json:"isDir"`
	Path             string    `json:"path"`
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/list", listPage).Methods("GET")
	myRouter.HandleFunc("/list/{userId}", listPageById).Methods("GET")
	fmt.Println("The server is on now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func listPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome! You have to enter your userId. Please insert /list/yourUserId")
}

func listPageById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["userId"]

	usr, err := user.Lookup(key)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	files, err := ioutil.ReadDir(usr.HomeDir)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	for _, loopFile := range files {
		file := FileInformation{
			Name:             loopFile.Name(),
			Size:             loopFile.Size(),
			LastModification: loopFile.ModTime(),
			Directory:        loopFile.IsDir(),
			Path: 			  usr.HomeDir,
		}
		json.NewEncoder(w).Encode(file)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "User / Directory not found")
	}
}

func main() {
	handleRequests()
}