package response

import (
	"github.com/gin-gonic/gin"
	"go-api-practice/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Success(c *gin.Context) {
	JSON(c, gin.H{
		"success": true,
		"message": "success",
	})
}

func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"success": true,
		"data":    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

func CreatedJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("Not Found", msg...),
	})
}

func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("Forbidden", msg...),
	})
}

func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("Internal Server Error", msg...),
	})
}

func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("Bad Request", msg...),
		"error":   err.Error(),
	})
}

func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("Internal Server Error", msg...),
		"error":   err.Error(),
	})
}

func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Validation Error",
		"errors":  errors,
	})
}

func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("Unauthorized", msg...),
	})
}

func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) == 0 {
		message = defaultMsg
	} else {
		message = msg[0]
	}
	return
}
