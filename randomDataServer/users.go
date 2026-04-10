package randomDataServer

import (
	"fmt"
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET request received")
	case "POST":
		fmt.Fprintf(w, "POST request received")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
