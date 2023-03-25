package controllers

import (
	"io"
	"net/http"
)

func ReverseProxyAPIPost(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	req.Header.Set("Authorization", r.Header.Get("Channel"))

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
