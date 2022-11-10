package routes

import (
	"e-commerce/handler"
	"e-commerce/token"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProductRoutes(g *echo.Group, productHandler *handler.ProductHandler) {
	g.GET("/products", productHandler.GetAllProducts)
	g.GET("/products/:id", productHandler.GetProductDetail)

	user := g.Group("/users/:id")
	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &token.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}))
	user.GET("/products", productHandler.GetAllProducts)
	user.GET("/products/:id", productHandler.GetProductDetail)

	admin := g.Group("/admin")
	admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &token.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}))
	admin.POST("/product", productHandler.CreateProduct)
	admin.GET("/products", productHandler.GetAllProducts)
	admin.GET("/products/:id", productHandler.GetProductDetail)
}

func CartRoutes(g *echo.Group, cartHandler *handler.CartHandler) {
	user := g.Group("/users/:id")
	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &token.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}))
	user.POST("/products/:id", cartHandler.AddToCart)
	user.GET("/carts", cartHandler.GetCart)
}

func CheckoutRoutes(g *echo.Group, checkoutHandler *handler.CheckoutHandler) {
	user := g.Group("/users/:id")
	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &token.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}))
	user.POST("/carts", checkoutHandler.CheckoutProduct)
}

func UserRoutes(g *echo.Group, userHandler *handler.UserHandler) {
	g.POST("/register", userHandler.RegisterNewUser)
	g.POST("/admin", userHandler.LoginAdmin)
	g.POST("/login", userHandler.LoginUser)
}
