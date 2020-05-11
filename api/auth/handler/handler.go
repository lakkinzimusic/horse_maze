package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lakkinzimusic/horse_maze/api/auth"
)

//Handler struct
type Handler struct {
	useCase auth.UseCase
}

//NewHandler func
func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SignUp func
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userInput signInput
	fmt.Println("1")
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.useCase.SignUp(userInput.Username, userInput.Password); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, userInput)
}

type signInResponse struct {
	Token string `json:"token"`
}

//SignIn func
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var userInput signInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := h.useCase.SignUp(userInput.Username, userInput.Password); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, userInput)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
