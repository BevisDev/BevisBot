package middleware

import (
	"github.com/BevisDev/BevisBot/internal/app/dto/response"
	"github.com/BevisDev/BevisBot/internal/app/enums"
	"github.com/BevisDev/godev/consts"
	"github.com/gin-gonic/gin"
	"strings"
)

func authHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(consts.Authorization)
		if accessToken == "" {
			response.Unauthorized(c, enums.InvalidCredentials)
			c.Abort()
			return
		}

		var token = strings.TrimPrefix(accessToken, consts.Bearer_)
		if token == "" {
			response.Unauthorized(c, enums.InvalidCredentials)
			c.Abort()
			return
		}

		//// validate token admin
		//cf := config.GetInstance().GetAppConfig()
		//if cf.TokenAdmin != "" && cf.TokenAdmin == token {
		//	c.Next()
		//	return
		//}
		//
		//// validate other token
		//userInfo, err := lib.KeyCloak.GetUserInfo(c.Request.Context(), token)
		//if err != nil || userInfo == nil {
		//	lib.Logger.Error(state, "error VerifyToken {}", err)
		//	response.Unauthorized(c, enums.InvalidCredentials)
		//	c.Abort()
		//	return
		//}
		//
		//partner := memory.GetInstance().Partners.GetByPreferredUsername(*userInfo.PreferredUsername)
		//if partner == nil || enums.IsSuspendedOrInactive(partner.Status) {
		//	response.Unauthorized(c, enums.InvalidClient)
		//	c.Abort()
		//	return
		//}
		//
		//ctx := context.WithValue(c.Request.Context(), constants.PartnerCtx, partner.Id)
		//c.Request = c.Request.WithContext(ctx)
		//
		//c.Next()
	}
}
