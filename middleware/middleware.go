package middleware

import (
    "github.com/golang-jwt/jwt/v5"
    "github.com/beego/beego/v2/server/web/context"
    "os"
    "api/controllers"
    "strings"
    "api/services"
    //"fmt"
    //"net/http"
	//"github.com/golang-jwt/jwt/v5"
)



// AuthMiddleware checks the Authorization header
func AuthMiddleware(ctx *context.Context) {
    //httpMethod := ctx.Input.Method()
    // if httpMethod != "GET"{

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
    PrefixUsers := "/v1/api/users"
	PrefixPosts := "/v1/api/posts"
	if strings.HasPrefix(ctx.Input.URI(), PrefixUsers) {
		UsersApi(ctx, claims)
	}
	if strings.HasPrefix(ctx.Input.URI(), PrefixPosts) {
		PostsApi(ctx, claims)
	}

}


func UsersApi(ctx *context.Context,claims *controllers.Claims){
    // Get the HTTP method
	httpMethod := ctx.Input.Method()

		// Check for PUT or DELETE methods
		if httpMethod == "DELETE" || httpMethod == "PUT" {
			// Extract the URL ID from the path parameters
			urlID := ctx.Input.Param(":id")

			// Check if URL ID is provided and if it matches the user's ID
			if urlID != "" && urlID != claims.Id {
				// Unauthorized
				ctx.Output.SetStatus(403)
				ctx.Output.JSON(map[string]string{"error": "Forbidden for this action!"}, false, false)
				return
			}
		}
}

func PostsApi(ctx *context.Context,claims *controllers.Claims){
        
        httpMethod := ctx.Input.Method()
		// Check for PUT or DELETE methods
		if httpMethod == "DELETE" || httpMethod == "PUT"{
            urlID := ctx.Input.Param(":post_id")
            user_id,err:=services.GetPostUserId(urlID)
            if err != nil{
                ctx.Output.SetStatus(400)
				ctx.Output.JSON(map[string]string{"error": err.Error()}, false, false)
				return
            }
            if user_id != "" && user_id != claims.Id {
				ctx.Output.SetStatus(403)
				ctx.Output.JSON(map[string]string{"error": "Forbidden for this action!"}, false, false)
				return
			}
        }
}