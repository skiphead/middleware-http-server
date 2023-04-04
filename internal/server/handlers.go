package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {

	Data := "Middleware layer"
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("", "")

		ctx := context.WithValue(r.Context(), "data", Data)

		next(w, r.WithContext(ctx))

	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	contextData := r.Context().Value("data").(string)

	err := json.NewEncoder(w).Encode(map[string]interface{}{"middleware": contextData})
	if err != nil {
		log.Println(err)
	}

}
