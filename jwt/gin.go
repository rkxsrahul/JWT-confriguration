package jwt

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func MwInitializer() *jwt.GinJWTMiddleware {

	//===================================================================================

	// intializing the gin jwt middleware by setting new configuration values
	authMiddleware := &jwt.GinJWTMiddleware{
		// name to display to the user
		Realm: "test zone",
		// passing key string by converting into bytes
		Key: []byte("1>MAxZ$%2(d_Y%]9dXLWyif?Ul47cc"),
		// Duration that a jwt token is valid
		Timeout: time.Second * 20,
		// this field allows clients to refresh their token until MaxRefresh has passed
		MaxRefresh: time.Hour * 720,
		// own Unauthorized func.
		Unauthorized: unauthorizedFunc,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		TokenLookup: "header:Authorization",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	// Initial middleware default setting.
	authMiddleware.MiddlewareInit()
	//===================================================================================

	return authMiddleware
}

func unauthorizedFunc(c *gin.Context, code int, msg string) {

	log.Println(c.Request.Header)
	log.Println(jwt.ExtractClaims(c))
	log.Println(c.Request.Header.Get("Authorization"))
	c.JSON(code, gin.H{"error": true, "code": code, "message": "Over"})

}
