package errors

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

func RederError(r *http.Request, w http.ResponseWriter, err error) {
	errs := err.(Errors)
	if len(errs) == 0 {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, NewInternalSystemError())
		return
	}

	statusCode, _ := strconv.Atoi(errs["code"].Error())
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	render.Status(r, statusCode)
	render.JSON(w, r, errs)
}
