package web

import (
	"fmt"
	"github.com/thedevsaddam/renderer"
	"io/ioutil"
	"log"
	"net/http"
)

var Rnd *renderer.Render

const filenameInput = "json-data"

// Render index page.
// Just for fun and learn golang.
func Home(w http.ResponseWriter, req *http.Request) {

	message := struct {
		UploadSuccess bool
	}{false}

	if req.Method == "POST" {
		_, _, err := req.FormFile(filenameInput)
		if err != http.ErrMissingFile {
			err := uploadFile(w, req)
			if err != nil {
				message.UploadSuccess = true
			}
		}
	}

	Rnd.HTML(w, http.StatusOK, "home", message)
}

func uploadFile(w http.ResponseWriter, req *http.Request) error {
	req.ParseMultipartForm(0)
	file, handler, err := req.FormFile(filenameInput)

	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return nil
	}

	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	tempFile, err := ioutil.TempFile("resources/data", "data-*.json")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	log.Printf("Successfully Uploaded File")

	return err
}
