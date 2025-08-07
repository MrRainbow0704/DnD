package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MrRainbow0704/DnD/api/middleware"
	"github.com/MrRainbow0704/DnD/internal/utils"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	con := struct {
		Name string `json:"name" validate:"required"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{jsonDecodeError: fmt.Errorf("error deconding JSON")},
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
	master := r.Context().Value(middleware.CtxUserID).(int64)
	c, err := db.NewCampaign(r.Context(), con.Name, master)
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusBadRequest,
			E{alreadyExistsError: fmt.Errorf("campaign already exists")},
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
			msgKey:      "Campaign created successfully",
			campaignKey: c,
		},
	)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CtxPathID).(int64)

	c, err := db.GetCampaign(r.Context(), int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusNotFound,
			E{notFoundError: fmt.Errorf("campaign with ID %d not found", id)},
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
			msgKey:      "Character retrieved successfully",
			campaignKey: c,
		},
	)
}
