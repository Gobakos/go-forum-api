package middleware

import (
    "github.com/golang-jwt/jwt/v5"
    "github.com/beego/beego/v2/server/web/context"
    "os"
    "test/controllers"
    //"fmt"
    //"net/http"
	//"github.com/golang-jwt/jwt/v5"
)



// AuthMiddleware checks the Authorization header
func AuthMiddleware(ctx *context.Context) {
    tokenString := ctx.Input.Header("Authorization")
    if tokenString == "" {
        ctx.Output.SetStatus(401)
        ctx.Output.JSON(map[string]string{"error": "No Authorization Header"}, false, false)
        return
    }
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }
    secretKey := []byte(os.Getenv("SECRET_KEY"))
    token, err := jwt.ParseWithClaims(tokenString, &controllers.Claims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        ctx.Output.SetStatus(401)
        ctx.Output.JSON(map[string]string{"error": "Invalid or expired token"}, false, false)
        return
    }
    claims, ok := token.Claims.(*controllers.Claims)
    if !ok || !token.Valid {
        ctx.Output.SetStatus(401)
        ctx.Output.JSON(map[string]string{"error": "Invalid or expired token"}, false, false)
        return
    }
    urlID := ctx.Input.Param(":id")
    // fmt.Printf("URL ID: %s, Claim ID: %s\n", urlID, claims.Id)
    if urlID != "" && urlID != claims.Id {
        ctx.Output.SetStatus(403)
        ctx.Output.JSON(map[string]string{"error": "Forbidden for this action!"}, false, false)
        return
    }
}