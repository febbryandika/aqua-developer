package main

import (
	"e-commerce/config"
	"e-commerce/handler"
	"e-commerce/repository"
	"e-commerce/routes"
	"e-commerce/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Database()
	config.AutoMigrate()

	e := echo.New()

	e.Use(middleware.Recover())

	g := e.Group("/api/v1")

	productRepository := repository.NewProductRepository(config.DB)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	cartRepository := repository.NewCartRepository(config.DB)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)

	checkoutRepository := repository.NewCheckoutRepository(config.DB)
	checkoutUseCase := usecase.NewCheckoutUseCase(checkoutRepository)
	checkoutHandler := handler.NewCheckoutHandler(checkoutUseCase)

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.ProductRoutes(g, productHandler)
	routes.CartRoutes(g, cartHandler)
	routes.CheckoutRoutes(g, checkoutHandler)
	routes.UserRoutes(g, userHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
