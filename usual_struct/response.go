package usual_struct

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Data interface{}
	Msg  string
}

func (re Response) JsonResponse(c *gin.Context) {
	c.JSON(http.StatusOK, re)
}

func (re Response) BadResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, re)
}
