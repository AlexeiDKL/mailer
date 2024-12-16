package handlers

import (
	"fmt"
	"net/http"
)

// CreateUrl
func (h *UrlHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Create Url")
	fmt.Fprint(w, response)
}

// GetUrls
func (h *UrlHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Get Urls")
	fmt.Fprint(w, response)
}

// EditUrl
func (h *UrlHandler) Edit(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Edit Url")
	fmt.Fprint(w, response)
}

// DeleteUrl
func (h *UrlHandler) Delete(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Delete Url")
	fmt.Fprint(w, response)
}

func (h *UrlHandler) UrlGet() string {
	url := NewUrlHandler()
	return url.get
}

func (h *UrlHandler) UrlCreate() string {
	url := NewUrlHandler()
	return url.create
}

func (h *UrlHandler) UrlEdit() string {
	url := NewUrlHandler()
	return url.edit
}

func (h *UrlHandler) UrlDelete() string {
	url := NewUrlHandler()
	return url.delete
}

func NewUrlHandler() *UrlHandler {
	var q UrlHandler
	q.create = "/createurl"
	q.delete = "/deleteurl"
	q.edit = "/editurl"
	q.get = "/geturls"

	return &q
}

type UrlHandler struct {
	Handler
}
