package handlers

import (
	"github.com/gin-gonic/gin"
	"goc/constant"
	"goc/usual_struct"
)

func GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usual_struct.BasePage
		var rep usual_struct.Response
		err := c.ShouldBind(&req)
		if err != nil {
			rep.Code = constant.INVALID_PARAMS
			rep.Msg = constant.GetMsg(constant.INVALID_PARAMS)
			rep.JsonResponse(c)
		}

	}
}
