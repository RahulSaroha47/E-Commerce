package routes

import (
	"e-commerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRequests *gin.Engine) {
    incomingRequests.POST("/users/signup",controllers.SignUp())
	incomingRequests.POST("/users/login",controllers.LogIn())
    incomingRequests.POST("/admin/addproduct",controllers.ProductViewerAdmin())
    incomingRequests.GET("/users/productview",controllers.SearchProduct())
    incomingRequests.GET("/users/search",controllers.SearchProductByQuery())
}
