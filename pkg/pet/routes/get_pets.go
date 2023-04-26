package routes

import (
	"io"
	"net/http"
)

func GetPets(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("http://localhost:9898/pets")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}
