package main

import (
	"fmt"
	"net/http"
)

type liteHandler struct {
	message string
}

func (m *liteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	hdl := &liteHandler{"Chao mung"}
	http.Handle("/welcome", hdl)
	http.ListenAndServe(":8000", nil)
}
