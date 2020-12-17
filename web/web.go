package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sshh12/apollgo/app"
	"golang.org/x/mobile/asset"
)

// SpaHandler for SPA
type SpaHandler struct{}

var extToContentType = map[string]string{
	"html": "text/html; charset=utf-8",
	"js":   "text/javascript; charset=UTF-8",
	"svg":  "image/svg+xml",
	"css":  "text/css; charset=UTF-8",
	"json": "application/json; charset=UTF-8",
	"png":  "image/png",
}

func serveAsset(w http.ResponseWriter, r *http.Request, file asset.File, ext string) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", extToContentType[ext])
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

// ServeWebApp starts http server
func ServeWebApp(state *app.State) {

	initCfg := state.GetCfg()
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	router.HandleFunc("/api/config", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(state.GetCfg())
	})
	router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(state.GetStatus())
	})
	router.PathPrefix("/").Handler(SpaHandler{})

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0:%d", initCfg.ApollgoPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
