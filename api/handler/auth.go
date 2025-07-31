package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MrRainbow0704/DnD/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Read and decode the JSON body into the [con] struct.
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

	// Fetch the user from the database using the provided username, then compare
	// the provided password with the one stored in the databas
	u, err := db.GetUserFromName(r.Context(), con.Username)
	if errors.Is(err, sql.ErrNoRows) {
		utils.ErrorJSON(
			w,
			http.StatusNotFound,
			E{notFoundError: fmt.Errorf("user with username \"%s\" not found", con.Username)},
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
	if !utils.ComparePassword(u.Passwd, []byte(con.Password), u.Salt) {
		utils.ErrorJSON(
			w,
			http.StatusUnauthorized,
			E{credentialsError: fmt.Errorf("access credentials don't match")},
		)
		return
	}
	userString, err := json.Marshal(u)
	if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusInternalServerError,
			E{internalError: fmt.Errorf("error encoding user data: %w", err)},
		)
		return
	}

	// Encode the user's ID into the JWT and return it to the user.
	_, tokenString, err := utils.TokenAuth.Encode(map[string]any{
		"sub":            u.ID,                           // Subject
		"exp":            time.Now().Add(time.Hour * 72), // Expiration time
		"nbf":            time.Now(),                     // Not before
		"iat":            time.Now(),                     // Issued at
		utils.JWTRoleKey: u.Role,                         // User role	
	})
	if err != nil {
		utils.ErrorJSON(
			w,
			http.StatusInternalServerError,
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		MaxAge:   cnf.JWTMaxAge,
		Expires:  time.Now().Add(time.Hour * 72),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "user",
		Value:    string(userString),
		MaxAge:   cnf.JWTMaxAge,
		Expires:  time.Now().Add(time.Hour * 72),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	utils.SendJSON(
		w,
		http.StatusOK,
		M{
			msgKey:   "Login successful",
			tokenKey: tokenString,
			userKey:  u,
		},
	)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "user",
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	utils.SendJSON(
		w,
		http.StatusOK,
		M{msgKey: "Logout successful"},
	)
}
