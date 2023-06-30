package handler

import "net/http"

type UserRequest struct {
	SiteAccess string
	MinUrltime string
	MaxUrltime string
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	myStats := h.ws.StatInfo.GetStatRecords()
	if r.URL.Path != "/stats" {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	err := tpl.ExecuteTemplate(w, "stats.html", myStats)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}

}
