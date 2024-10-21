package controller

import (
	"context"
	"serverless/domain"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
)

func Login(ctx context.Context, payload domain.LoginRequest) response.Base {
	errVal := map[string]string{}
	if payload.Username == "" {
		errVal["username"] = "username is required"
	}

	if payload.Password == "" {
		errVal["password"] = "password is required"
	}

	if len(errVal) > 0 {
		return response.ErrorValidation(errVal, "validation error")
	}

	if payload.Username != "admin" || payload.Password != "rahasia" {
		return response.Error(401, "invalid username or password")
	}

	return response.Success(map[string]any{
		"message": "login success",
	})
}
