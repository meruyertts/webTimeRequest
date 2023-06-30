package handler

import "net/http"

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}
