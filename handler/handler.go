package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"webTimeRequest/websitestats"
)

type Handler struct {
	ws *websitestats.WebsiteStats
}

func NewHandler(ws websitestats.WebsiteStats) *Handler {
	return &Handler{
		ws: &ws,
	}
}

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("./ui/templates/index.html", "./ui/templates/stats.html", "./ui/templates/error.html")
	if err != nil {
		log.Fatal(err)
	}
}
func (h *Handler) InitRoutes() {
	router := http.NewServeMux()
	router.HandleFunc("/", h.IndexHandler)
	router.HandleFunc("/site-access", h.SiteAccess)
	router.HandleFunc("/min-time", h.MinTime)
	router.HandleFunc("/max-time", h.MaxTime)
	router.HandleFunc("/stats", h.Stats)
	fmt.Println("Server launched at http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("This host is already run")
	}
}
