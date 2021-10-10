package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UploadHandler struct {
	HostAddr  string
	UploadDir string
}

// Здесь мы реализовали загрузку файла на сервер.
func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	filePath := h.UploadDir + "/" + header.Filename
	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "File %s has been successfully uploaded\n", header.Filename)

	fileLink := h.HostAddr + "/files/" + header.Filename
	//fmt.Println("fileLink", fileLink, "filePath", filePath)
	// fmt.Fprintln(w, fileLink)

	req, err := http.NewRequest(http.MethodHead, fileLink, nil)
	if err != nil {
		http.Error(w, "Unable to check file", http.StatusInternalServerError)
		return
	}

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to check file", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintln(w, fileLink); err != nil {
		log.Printf("error writing to response body: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
}
