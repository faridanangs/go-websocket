package internaluser

import "github.com/gofiber/fiber/v2"

func UserResponseWeb(req User) UserResponse {
	return UserResponse{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: req.CreatedAt,
	}
}
func UserResponsesWeb(req []User) []UserResponse {
	var responses []UserResponse
	for _, u := range req {
		responses = append(responses, UserResponseWeb(u))
	}

	return responses
}

func WebResponseJosn(ctx *fiber.Ctx, code int, sttus string, data any) error {
	return ctx.Status(code).JSON(WebResponse{
		Code:   code,
		Status: sttus,
		Data:   data,
	})
}
