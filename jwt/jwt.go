package jwt

import (
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

func GinJwtToken() map[string]interface{} {

	// intializing middleware
	mw := MwInitializer()

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	// extracting claims in form of map
	claims := token.Claims.(jwt.MapClaims)

	// extracting expire time
	expire := mw.TimeFunc().Add(mw.Timeout)

	// setting claims
	claims["id"] = 1
	claims["name"] = "rahul"
	claims["email"] = "rahulkumarrvc"
	claims["sys_role"] = "user"
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()

	mapd := map[string]interface{}{"token": "", "expire": ""}

	// signing token
	tokenString, err := token.SignedString(mw.Key)
	if err != nil {
		return mapd
	}

	// passing map eith all information
	mapd = map[string]interface{}{"error": false,
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339)}

	return mapd
}
