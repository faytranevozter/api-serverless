package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serverless/domain"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
)

func Json(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Content-Type", "application/json")

	res := response.Success(map[string]any{
		"user": domain.User{
			Name: "andi",
		},
		"something": "hello",
	})

	fmt.Println(w.Header())

	jsonResp, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}
