package link

import (
	"fmt"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/resp"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /link/{id}", handler.GoTo())
	router.HandleFunc("PATH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/link")
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp.SetJson(w, createdLink, 201)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/{alias}")
		body, err := req.HandleBody[LinkReadRequest](&w, r)
		if err != nil {
			return
		}
		// res := NewLink(body.Url)
		resp.SetJson(w, body, 201)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/link/{id}")
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		resp.SetJson(w, body, 201)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/link/{id}")
		id := r.PathValue("id")
		fmt.Println(id)
		// body, err := req.HandleBody[LinkReadRequest](&w, r)
		// if err != nil {
		// 	return
		// }
		// resp.SetJson(w, body, 201)
	}
}
