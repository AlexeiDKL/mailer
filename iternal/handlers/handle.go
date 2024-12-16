package handlers

import (
	"fmt"
	"net/http"
)

type Handler struct {
	get    string
	create string
	edit   string
	delete string
}

type Handlers interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	UrlGet() string
	UrlCreate() string
	UrlEdit() string
	UrlDelete() string
}

/*	Заполнение бд /createcompany
	компании

	Нужно получить: название компании
	Возврат удача, или не удача

	Название компании, пытаемся получить её id в "company"
		если не получилось, то пытаемся добавить в бд "company"
			если не удалось, возвращаем ошибку
		    если удалось, то получаем id
*/

// Create company information
// urls, urlsType and domains
// Create company information
// urls, urlsType and domains
func CreateCompanyInfo(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Create Company Info")
	fmt.Fprint(w, response)
}

func GetPin(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Get Pin")
	fmt.Fprint(w, response)
}

func ValidPin(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Valid Pin")
	fmt.Fprint(w, response)
}
