package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/res"
	"net/http"
	"net/mail"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		if payload.Email == "" {
			res.Json(w, "Email required", http.StatusBadRequest)
			return
		}

		_, err = mail.ParseAddress(payload.Email)
		if err != nil {
			res.Json(w, "Wrong Email", http.StatusBadRequest)
			return
		}

		if payload.Password == "" {
			res.Json(w, "Password required", http.StatusBadRequest)
			return
		}

		fmt.Println(payload)
		data := LoginResponse{
			Token: "1234567890",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
