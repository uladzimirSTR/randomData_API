package randomDataServer

import (
	"net/http"
)

func RandomDataServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Users)
	http.ListenAndServe(":9909", mux)
}
