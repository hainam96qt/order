package common

import (
	"context"
	"order-gokomodo/pkg/entities/common"
)

func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID := ctx.Value("user_id")
	switch userID.(type) {
	case int:
		return userID.(int), nil
	}
	return 0, common.NewErr("Can't get user_id from context")
}

func GetUserRoleFromContext(ctx context.Context) (string, error) {
	role := ctx.Value("role")
	switch role.(type) {
	case string:
		return role.(string), nil
	}
	return "", common.NewErr("Can't get role from context")
}
