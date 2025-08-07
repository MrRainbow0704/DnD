package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/MrRainbow0704/DnD/api/middleware"
	"github.com/MrRainbow0704/DnD/internal/utils"
)

func GetThisUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CtxUserID).(int64)
	u, err := db.GetUser(r.Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusNotFound,
			E{notFoundError: fmt.Errorf("user with ID %d not found", id)},
		)
		return
	} else if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusInternalServerError,
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}

	utils.SendJSON(
		w,
		http.StatusOK,
		M{
			msgKey:  "User retrieved successfully",
			userKey: u,
		},
	)
}

func EditThisUser(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(
		w,
		http.StatusNotImplemented,
		M{msgKey: "Not implemented"},
	)
}

func DeleteThisUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CtxUserID).(int64)

	err := db.DelUser(r.Context(), id)
	if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusInternalServerError,
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}

	utils.SendJSON(
		w,
		http.StatusOK,
		M{msgKey: "User deleted successfully"},
	)
}
