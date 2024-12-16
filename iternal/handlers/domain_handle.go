package handlers

import (
	"fmt"
	"net/http"
)

// GetDomains
func (h *DomainHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Get Domains")
	fmt.Fprint(w, response)
}

// CreateDomain
func (h *DomainHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Create Domain")
	fmt.Fprint(w, response)
}

// EditDomain
func (h *DomainHandler) Edit(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Edit Domain")
	fmt.Fprint(w, response)
}

// DeleteDomain
func (h *DomainHandler) Delete(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Delete domain")
	fmt.Fprint(w, response)
}

func (h *DomainHandler) UrlGet() string {
	url := NewDomainHandler()
	return url.get
}

func (h *DomainHandler) UrlCreate() string {
	url := NewDomainHandler()
	return url.create
}

func (h *DomainHandler) UrlEdit() string {
	url := NewDomainHandler()
	return url.edit
}

func (h *DomainHandler) UrlDelete() string {
	url := NewDomainHandler()
	return url.delete
}

func NewDomainHandler() *DomainHandler {
	var q DomainHandler
	q.create = "/createdomain"
	q.delete = "/deletedomain"
	q.edit = "/editdomain"
	q.get = "/getdomain"

	return &q
}

type DomainHandler struct {
	Handler
}
