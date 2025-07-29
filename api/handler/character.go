package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrRainbow0704/DnD/api/middleware"
	"github.com/MrRainbow0704/DnD/internal/utils"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	con := struct {
		Name string `json:"name" validate:"required"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{jsonDecodeError: fmt.Errorf("Error deconding JSON")},
		)
		return
	}
	if con.Name == "" {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{emptyFieldError: fmt.Errorf("name field is empty")},
		)
		return
	}
	owner := r.Context().Value(middleware.CtxUserID).(int64)
	c, err := db.NewCharacter(r.Context(), owner, con.Name)
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{alreadyExistsError: fmt.Errorf("character already exists")},
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
			msgKey:       "Character created successfully",
			characterKey: c,
		},
	)
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{invalidParamError: fmt.Errorf("invalid character ID")},
		)
		return
	}

	c, err := db.GetCharacter(r.Context(), int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusNotFound,
			E{notFoundError: fmt.Errorf("character with ID %d not found", id)},
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
		http.StatusInternalServerError,
		M{
			msgKey:       "Character retrieved successfully",
			characterKey: c,
		},
	)
}
