package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) MinTime(w http.ResponseWriter, r *http.Request) {
	h.ws.StatInfo.UpdateMinTimeRequests()
	if r.URL.Path != "/min-time" {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	minWebsite := h.ws.UpdateMinTimeInfo()
	message := ""
	if minWebsite == "" {
		message = "Unfortunately, The info about minimum access time is unavailable"
	} else {
		message = fmt.Sprintf("The website  with the minimum access time is %s", minWebsite)
	}
	result := UserRequest{MinUrltime: message}
	err := tpl.ExecuteTemplate(w, "index.html", result)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}
