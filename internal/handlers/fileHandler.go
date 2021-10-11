package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type FileHendler struct {
	PathDir string
}

func (fs *FileHendler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	ftype := r.FormValue("type")
	ftype = strings.TrimPrefix(ftype, ".")

	dir, err := ioutil.ReadDir(fs.PathDir)
	if err != nil {
		log.Printf("directory %s connection error: %v", fs.PathDir, err)
	}

	for _, file := range dir {
		if !file.IsDir() {
			ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
			if ext == ftype {
				fmt.Fprintf(w, "Name: %s расширение %s размер %d байт\n", filepath.Base(file.Name()), ext, file.Size())
			}
			if ftype == "" {
				fmt.Fprintf(w, "Name: %s расширение %s размер %d байт\n", filepath.Base(file.Name()), ext, file.Size())
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}
