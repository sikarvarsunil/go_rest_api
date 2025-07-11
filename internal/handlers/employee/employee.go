package employee

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/sikarvarsunil/go_rest_api/internal/types"
	"github.com/sikarvarsunil/go_rest_api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee types.Employee

		err := json.NewDecoder(r.Body).Decode(&employee)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		slog.Info("Creating a employee")
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
