package handlers

import (
	"fmt"
	"net/http"
)

// GetUrlTypes
func (h *UrlTypeHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Get Url Types")
	fmt.Fprint(w, response)
}

// CreateUrlType
func (h *UrlTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Create Url Type")
	fmt.Fprint(w, response)
}

// EditUrlType
func (h *UrlTypeHandler) Edit(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Edit Url Type")
	fmt.Fprint(w, response)
}

// DeleteUrlType
func (h *UrlTypeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Delete Url Type")
	fmt.Fprint(w, response)
}

func (h *UrlTypeHandler) UrlGet() string {
	url := NewUrlTypeHandler()
	return url.get
}

func (h *UrlTypeHandler) UrlCreate() string {
	url := NewUrlTypeHandler()
	return url.create
}

func (h *UrlTypeHandler) UrlEdit() string {
	url := NewUrlTypeHandler()
	return url.edit
}

func (h *UrlTypeHandler) UrlDelete() string {
	url := NewUrlTypeHandler()
	return url.delete
}

func NewUrlTypeHandler() *UrlTypeHandler {
	var q UrlTypeHandler
	q.create = "/createurltype"
	q.delete = "/deleteurltype"
	q.edit = "/editurltype"
	q.get = "/geturltypes"

	return &q
}

type UrlTypeHandler struct {
	Handler
}
