package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) SiteAccess(w http.ResponseWriter, r *http.Request) {
	h.ws.StatInfo.UpdateTimeRequests()
	if r.URL.Path != "/site-access" {
		errorHeader(w, http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		errorHeader(w, http.StatusMethodNotAllowed)
		return
	}
	websiteName := r.FormValue("websiteName")
	websiteTime := h.ws.UpdateTimeInfo(websiteName)
	message := fmt.Sprintf("the website %s access duration time is %s", websiteName, websiteTime.String())
	result := UserRequest{SiteAccess: message}
	err := tpl.ExecuteTemplate(w, "index.html", result)
	if err != nil {
		fmt.Println("it's here")
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}
