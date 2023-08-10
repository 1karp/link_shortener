package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/1karp/link_shortener/internal/app/config"
	"github.com/1karp/link_shortener/internal/app/storage"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	ShortURL string `json:"short_url"`
}

func APIShortenHandler(rw http.ResponseWriter, req *http.Request, storage storage.Storage, cfg config.Config) {
	var reqModel shortenURLRequest

	jsonDecoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	if err := jsonDecoder.Decode(&reqModel); err != nil {
		http.Error(rw, "Failed to decode request JSON body", http.StatusInternalServerError)
		return
	}

	if reqModel.URL == "" {
		http.Error(rw, "Invalid request: URL missing", http.StatusBadRequest)
		return
	}

	shortCode := generateShortURL()
	fullShortURL := generateFullShortURL(cfg.BaseShortURLAddress, shortCode)

	respModel := shortenURLResponse{
		ShortURL: fullShortURL,
	}

	data, err := json.Marshal(respModel)
	if err != nil {
		http.Error(rw, "Failed to create response body", http.StatusInternalServerError)
		return
	}

	storage.Set(shortCode, reqModel.URL)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	_, err = rw.Write(data)
	if err != nil {
		http.Error(rw, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func generateFullShortURL(baseAddress, code string) string {
	return baseAddress + code
}
