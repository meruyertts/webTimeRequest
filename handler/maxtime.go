package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) MaxTime(w http.ResponseWriter, r *http.Request) {
	h.ws.StatInfo.UpdateMaxTimeRequests()
	if r.URL.Path != "/max-time" {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	maxWebsite := h.ws.UpdateMaxTimeInfo()
	message := ""
	if maxWebsite == "" {
		message = "Unfortunately, The info about maximum access time is unavailable"
	} else {
		message = fmt.Sprintf("The website  with the max access time is %s", maxWebsite)
	}
	result := UserRequest{MaxUrltime: message}
	err := tpl.ExecuteTemplate(w, "index.html", result)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}
