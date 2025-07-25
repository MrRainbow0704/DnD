package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrRainbow0704/DnD/internal/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	con := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{jsonDecodeError: fmt.Errorf("Error deconding JSON")},
		)
		return
	}
	if con.Username == "" || len(con.Password) == 0 {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{emptyFieldError: fmt.Errorf("username or password fields are empty")},
		)
		return
	}

	salt, err := utils.RandomSecret(cnf.PasswdSaltLen)
	if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusInternalServerError,
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}

	hash := utils.HashPassword([]byte(con.Password), salt)
	u, err := db.NewUser(r.Context(), con.Username, hash, salt)
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{alreadyExistsError: fmt.Errorf("user already exists")},
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
		http.StatusCreated,
		M{
			msgKey:  "User created successfully",
			userKey: u,
		},
	)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{invalidParamError: fmt.Errorf("invalid user ID")},
		)
		return
	}

	u, err := db.GetUser(r.Context(), int64(id))
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
		http.StatusInternalServerError,
		M{
			msgKey:  "Character retrieved successfully",
			userKey: u,
		},
	)
}
