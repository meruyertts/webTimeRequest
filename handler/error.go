package handler

import "net/http"

type ErrorBody struct {
	Status  int
	Message string
}

func errorHeader(w http.ResponseWriter, status int) {
	errH := setError(status)
	err := tpl.ExecuteTemplate(w, "error.html", errH)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
}

func setError(status int) *ErrorBody {
	return &ErrorBody{
		Status:  status,
		Message: http.StatusText(status),
	}
}
