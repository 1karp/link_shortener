package handlers

import (
	"encoding/json"
	"errors"
	"io"
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

func ApiShortenHandler(rw http.ResponseWriter, req *http.Request, storage storage.Storage, cfg config.Config) {
	requestBody, err := parseShortenURLRequest(req)
	if err != nil {
		http.Error(rw, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	shortCode := generateShortURL()
	fullShortURL := generateFullShortURL(cfg.BaseShortURLAddress, shortCode)

	responseBody := shortenURLResponse{
		ShortURL: fullShortURL,
	}

	data, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(rw, "Failed to create response body", http.StatusInternalServerError)
		return
	}

	storage.Set(shortCode, requestBody.URL)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	_, err = rw.Write(data)
	if err != nil {
		http.Error(rw, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func parseShortenURLRequest(req *http.Request) (*shortenURLRequest, error) {
	rawRequestBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("unable to read request body")
	}

	if !json.Valid(rawRequestBody) {
		return nil, errors.New("invalid JSON body")
	}

	var reqBody shortenURLRequest
	if err := json.Unmarshal(rawRequestBody, &reqBody); err != nil {
		return nil, errors.New("failed to unmarshal JSON body")
	}

	if len(reqBody.URL) == 0 {
		return nil, errors.New("URL missing in the request body")
	}

	return &reqBody, nil
}

func generateFullShortURL(baseAddress, code string) string {
	return baseAddress + "/" + code
}
