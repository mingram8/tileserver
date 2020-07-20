package middleware

import (
	"log"
	"net/http"
	"path/filepath"
)

type HttpHandler struct {
	RootDirectory string
	Handler       http.Handler
}

func (httpHandler *HttpHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
func (httpHandler *HttpHandler) Tiles(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Origin, Accept,X-Requested-With, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	z, _ := r.URL.Query()["z"]
	x, _ := r.URL.Query()["x"]
	y, _ := r.URL.Query()["y"]
	layer, _ := r.URL.Query()["layer"]
	tmpDir := filepath.FromSlash(httpHandler.RootDirectory)
	tmpDir = filepath.Join(tmpDir,layer[0])
	tmpDir = filepath.Join(tmpDir,x[0])
	tmpDir = filepath.Join(tmpDir,y[0])
	tmpDir = filepath.Join(tmpDir,z[0])
	log.Print(tmpDir)
	http.ServeFile(w, r, tmpDir)

}
func (httpHandler *HttpHandler) NewLogger(next http.Handler) *HttpHandler {
	httpHandler.Handler = next
	return httpHandler
}


func (httpHandler HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpHandler.Handler.ServeHTTP(w, r)

}