package handlers

import (
	"fmt"
	"net/http"
)

func (h *CompanyHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Get companies")
	fmt.Fprint(w, response)
}

func (h *CompanyHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Create Company")
	fmt.Fprint(w, response)
}

// EditCompany
func (h *CompanyHandler) Edit(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Edit Company")
	fmt.Fprint(w, response)
}

// DeleteCompany
func (h *CompanyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Delete Company")
	fmt.Fprint(w, response)
}

func (h *CompanyHandler) UrlGet() string {
	company := NewCompanyHandler()
	return company.get
}

func (h *CompanyHandler) UrlCreate() string {
	company := NewCompanyHandler()
	return company.create
}

func (h *CompanyHandler) UrlEdit() string {
	company := NewCompanyHandler()
	return company.edit
}

func (h *CompanyHandler) UrlDelete() string {
	company := NewCompanyHandler()
	return company.delete
}

func NewCompanyHandler() *CompanyHandler {
	var q CompanyHandler
	q.create = "/createurl"
	q.delete = "/deleteurl"
	q.edit = "/editurl"
	q.get = "/geturls"

	return &q
}

type CompanyHandler struct {
	Handler
}
