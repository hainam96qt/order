package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"order-gokomodo/pkg/usecase/service/auth"
	"strings"
)

func Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO validate token
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		log.Println(splitToken)
		decodeValue, err := auth.DecodeToken(reqToken)
		if err != nil {
			WriteError(w, r, err, 500)
		}
		ctx := context.WithValue(r.Context(), "user_id", decodeValue)
		ctx = context.WithValue(ctx, "role", decodeValue)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// handle Error
func WriteError(w http.ResponseWriter, r *http.Request, err error, code int) {
	if code != 200 {
		log.Println("Error:", err)
	}
	if code == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	type ErrorResponse struct {
		MessageError string `json:"message_error"`
		Code         int    `json:"code"`
	}
	b, err := json.Marshal(ErrorResponse{
		Code:         code,
		MessageError: err.Error(),
	})
	if err != nil {
		code = http.StatusInternalServerError
	}
	if _, err := w.Write(b); err != nil {
		log.Println("writing response to response writer")
	}
}

func WriteResponse(tx context.Context, w http.ResponseWriter, r *http.Request, bodyRespnse interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(bodyRespnse)
	if err != nil {
		log.Print("Encode Repose have same error")
	}
}
