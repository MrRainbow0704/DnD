package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrRainbow0704/DnD/internal/config"
	database "github.com/MrRainbow0704/DnD/internal/db"
	"github.com/MrRainbow0704/DnD/internal/utils"
)

type M map[string]any   // Alias for map[string]any
type E map[string]error // Alias for map[string]error
var db = database.Get() // Database instance
var cnf = config.Get()  // Database instance
const (
	msgKey             = "message" // Key for the message in the response map
	tokenKey           = "token"   // Key for the token in the response map
	userKey            = "user"    // Key for the user in the response map
	characterKey       = "character"
	jsonDecodeError    = "JSON_DECODE"
	emptyFieldError    = "EMPTY_FIELD"
	invalidParamError  = "INVALID_URL_PARAMETER"
	notFoundError      = "USER_NOT_FOUND"
	internalError      = "INTERNAL"
	credentialsError   = "INVALID_CREDENTIALS"
	alreadyExistsError = "USER_ALREADY_EXISTS"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Read and decode the JSON body into the [con] struct.
	con := struct {
		Username string `json:"username" validate:"required"`
		Password []byte `json:"password" validate:"required"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Error decoding JSON"},
			E{jsonDecodeError: err},
		)
		return
	}
	if con.Username == "" || len(con.Password) == 0 {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Username and password are required"},
			E{emptyFieldError: fmt.Errorf("username or password fields are empty")},
		)
		return
	}

	// Fetch the user from the database using the provided username, then compare
	// the provided password with the one stored in the databas
	u, err := db.GetUserFromName(r.Context(), con.Username)
	if errors.Is(err, sql.ErrNoRows) {
		utils.SendJSON(
			w,
			http.StatusNotFound,
			M{msgKey: "User not found"},
			E{notFoundError: fmt.Errorf("user with username \"%s\" not found", con.Username)},
		)
		return
	} else if err != nil {
		utils.SendJSON(
			w,
			http.StatusInternalServerError,
			M{msgKey: "Internal server error"},
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}
	if !utils.ComparePassword(u.Passwd, con.Password, u.Salt) {
		utils.SendJSON(
			w,
			http.StatusUnauthorized,
			M{msgKey: "Internal server error"},
			E{credentialsError: fmt.Errorf("access credentials don't match")},
		)
		return
	}

	// Encode the user's ID into the JWT and return it to the user.
	_, tokenString, err := utils.TokenAuth.Encode(map[string]any{
		utils.JWTIDKey:   u.ID,
		utils.JWTRoleKey: u.Role,
	})
	if err != nil {
		utils.SendJSON(
			w,
			http.StatusInternalServerError,
			M{msgKey: "Internal server error"},
			E{credentialsError: fmt.Errorf("internal server error")},
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "jwt",
		MaxAge: cnf.JWTMaxAge,
		Value:  tokenString,
	})
	utils.SendJSON(
		w,
		http.StatusOK,
		M{
			msgKey:   "Login successful",
			tokenKey: tokenString,
		},
		nil,
	)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "jwt",
		MaxAge: -1,
	})
	utils.SendJSON(
		w,
		http.StatusOK,
		M{msgKey: "Logout successful"},
		nil,
	)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	con := struct {
		Username string `json:"username" validate:"required"`
		Password []byte `json:"password" validate:"required"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Error decoding JSON"},
			E{jsonDecodeError: err},
		)
		return
	}
	if con.Username == "" || len(con.Password) == 0 {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Username and password are required"},
			E{emptyFieldError: fmt.Errorf("username or password fields are empty")},
		)
		return
	}

	salt, err := utils.RandomSecret(cnf.PasswdSaltLen)
	if err != nil {
		utils.SendJSON(
			w,
			http.StatusInternalServerError,
			M{msgKey: "Failed to generate salt"},
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}
	u, err := db.NewUser(r.Context(), con.Username, con.Password, salt)
	if errors.Is(err, sql.ErrNoRows) {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "User already exists"},
			E{alreadyExistsError: fmt.Errorf("user already exists")},
		)
		return
	} else if err != nil {
		utils.SendJSON(
			w,
			http.StatusInternalServerError,
			M{msgKey: "Internal server error"},
			E{internalError: fmt.Errorf("internal server error")},
		)
		return
	}

	utils.SendJSON(
		w,
		http.StatusOK,
		M{
			msgKey:  "User created successfully",
			userKey: u,
		},
		nil,
	)
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	con := struct {
		Name string `json:"name" validate:"required"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Error decoding JSON"},
			E{jsonDecodeError: err},
		)
		return
	}
	if con.Name == "" {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Name is required"},
			E{emptyFieldError: fmt.Errorf("name field is empty")},
		)
		return
	}
	owner := r.Context().Value("USER_ID").(int64)
	c, err := db.NewCharacter(r.Context(), owner, con.Name)
	if errors.Is(err, sql.ErrNoRows) {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Character already exists"},
			E{alreadyExistsError: fmt.Errorf("character already exists")},
		)
		return
	} else if err != nil {
		utils.SendJSON(
			w,
			http.StatusInternalServerError,
			M{msgKey: "Internal server error"},
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
		nil,
	)
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendJSON(
			w,
			http.StatusBadRequest,
			M{msgKey: "Invalid character ID"},
			E{invalidParamError: fmt.Errorf("invalid character ID")},
		)
		return
	}

	c, err := db.GetCharacter(r.Context(), int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		utils.SendJSON(
			w,
			http.StatusNotFound,
			M{msgKey: "Character not found"},
			E{notFoundError: fmt.Errorf("character with ID %d not found", id)},
		)
		return
	} else if err != nil {
		utils.SendJSON(
			w,
			http.StatusInternalServerError,
			M{msgKey: "Internal server error"},
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
		nil,
	)
}
