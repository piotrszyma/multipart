package main

import (
	"fmt"
	"net/http"

	_ "embed"

	"github.com/piotrszyma/multipart"
)

//go:embed index.html
var indexPage []byte

type formData struct {
	Name string `multipart:"name"`
	Optional *string `multipart:"optional"`
}

func BuildMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleIndexPage))
	mux.Handle("/form", http.HandlerFunc(handleForm))
	return mux
}

func main() {
	http.ListenAndServe(":8080", BuildMux())
}

func handleIndexPage(w http.ResponseWriter, request *http.Request) {
	w.Write(indexPage)
}

func handleForm(w http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var formData formData

	err = multipart.Bind(request.MultipartForm, &formData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(formData.Name)

	http.Redirect(w, request, "/", http.StatusSeeOther)
}

// func createEmployee(w http.ResponseWriter, request *http.Request) {
// 	err := request.ParseMultipartForm(32 << 20) // maxMemory 32MB
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	//Access the name key - First Approach
// 	fmt.Println(request.Form["name"])
// 	//Access the name key - Second Approach
// 	fmt.Println(request.PostForm["name"])
// 	//Access the name key - Third Approach
// 	fmt.Println(request.MultipartForm.Value["name"])
// 	//Access the name key - Fourth Approach
// 	fmt.Println(request.PostFormValue("name"))
// 	//Access the age key - First Approach
// 	fmt.Println(request.Form["age"])
// 	//Access the age key - Second Approach
// 	fmt.Println(request.PostForm["age"])
// 	//Access the age key - Third Approach
// 	fmt.Println(request.MultipartForm.Value["age"])
// 	//Access the age key - Fourth Approach
// 	fmt.Println(request.PostFormValue("age"))

// 	//Access the photo key - First Approach
// 	_, h, err := request.FormFile("photo")
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	//Access the photo key - Second Approach
	// for _, h := range request.MultipartForm.File["photo"] {
// 		err := saveFile(h, "mapaccess")
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 	}
// 	w.WriteHeader(200)
// 	return
// }
