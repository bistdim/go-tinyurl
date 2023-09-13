package main

import (
	"net/http"

	"math/rand"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type urlResource struct{}

type hashResult struct {
	TinyUrl string `json:"tinyUrl"`
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (rs urlResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create) // POST /v1/url - create a new hash by url

	return r
}

func (rs urlResource) Create(w http.ResponseWriter, r *http.Request) {
	tinyUrl := r.Host + "/" + generateRandomString(6)

	render.JSON(w, r, &hashResult{TinyUrl: tinyUrl})
}
