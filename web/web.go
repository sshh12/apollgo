package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/mobile/asset"
)

// SpaHandler for SPA
type SpaHandler struct{}

func serveAsset(w http.ResponseWriter, r *http.Request, file asset.File, ext string) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if ext == "html" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	} else if ext == "js" {
		w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
	} else if ext == "svg" {
		w.Header().Set("Content-Type", "image/svg+xml")
	} else if ext == "css" {
		w.Header().Set("Content-Type", "text/css; charset=UTF-8")
	} else if ext == "json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else if ext == "png" {
		w.Header().Set("Content-Type", "image/png")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = strings.ReplaceAll(path[1:], "/", "__")

	if file, err := asset.Open(path); err == nil {
		parts := strings.Split(path, ".")
		serveAsset(w, r, file, parts[len(parts)-1])
		return
	}
	indexFile, err := asset.Open("index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		serveAsset(w, r, indexFile, "html")
	}
}

// ServeSPA starts the react server
func ServeSPA() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.PathPrefix("/").Handler(SpaHandler{})

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
