package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Nama   string
	Nim    string
	Alamat string
}

var Data = []Student{
	Student{"Yoga", "123", "Araya"},
	Student{"Pras", "234", "Malang"},
}

func getuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

// func createuser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "POST" {
// 		var newData = Student{}
// 		Data = append(Data, newData)

// 	}

// 	http.Error(w, "Bad Request", http.StatusBadRequest)
// }

func main() {
	http.HandleFunc("/getuser", getuser)
	// http.HandleFunc("/create", createuser)

	fmt.Println("starting web server at http://localhost:5000/")
	http.ListenAndServe(":5000", nil)
}
