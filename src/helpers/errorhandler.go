package helpers

import (
	"EspressoTracking/database/errorlogging"
	"github.com/gin-gonic/gin"
	"runtime"
)

type ErrorHandler struct{}

func (eh *ErrorHandler) handleError(c *gin.Context, err error) {
	ec := errorlogging.NewErrorLogRepository()
	_, file, line, _ := runtime.Caller(0)
	errRecord := errorlogging.NewErrorRecord(c.ClientIP(), file+" "+string(line), err.Error(), err.Error())
	ec.CreateErrorLogRecord(*errRecord)
	c.JSON(500, gin.H{"error": "Internal Error"})
}
