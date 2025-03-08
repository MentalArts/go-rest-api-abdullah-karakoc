package middlewares

import (
	"mentalartsapi/internal/dto"
	"mentalartsapi/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Eğer hata varsa, dön
		for _, err := range c.Errors {
			switch err.Err {
			case utils.ErrInvalidID:
				c.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Message: err.Err.Error()})
			case utils.ErrNotFound:
				c.JSON(http.StatusNotFound, dto.ErrorResponseDTO{Message: err.Err.Error()})
			case utils.ErrBadRequest:
				c.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Message: err.Err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, dto.ErrorResponseDTO{Message: utils.ErrInternal.Error()})
			}
			return
		}
	}
}
