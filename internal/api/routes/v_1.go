package routes

import (
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/middleware"
)

func RouteV1(av *APIVersion) {
	userController := controller.NewUserController(av.contract.Service.User)
	authController := controller.NewAuthController(av.contract.Service.Auth)
	roleController := controller.NewRoleController(av.contract.Service.Role)
	supplierController := controller.NewSupplierController(av.contract.Service.Supplier)
	itemController := controller.NewItemController(av.contract.Service.Item)

	av.api.Use(middleware.LogMiddleware) // use middleware logger

	// auth routes
	auth := av.api.Group("/auth")

	auth.POST("/login", authController.Login, middleware.RateLimitMiddleware(5, 300)) // limit to 5 requests per 5 minutes
	auth.POST("/refresh", authController.RefreshAccessToken)
	auth.POST("/logout", authController.Logout)

	// user routes
	user := av.api.Group("/users")
	user.Use(middleware.JWTMiddleware()) // use middleware jwt general on user routes

	user.GET("", userController.Get)
	user.GET("/:id", userController.GetById)
	user.POST("", userController.Create)
	user.PUT("/:id", userController.Update)
	user.DELETE("/:id", userController.Delete)

	// role routes
	role := av.api.Group("/roles")
	role.Use(middleware.JWTMiddleware()) // use middleware jwt general on role routes

	role.GET("", roleController.Get)
	role.GET("/:id", roleController.GetById)
	role.POST("", roleController.Create)
	role.PUT("/:id", roleController.Update)
	role.DELETE("/:id", roleController.Delete)

	// supplier routes
	supplier := av.api.Group("/suppliers")
	supplier.Use(middleware.JWTMiddleware()) // use middleware jwt general on supplier routes

	supplier.GET("", supplierController.Get)
	supplier.GET("/:id", supplierController.GetById)
	supplier.POST("", supplierController.Create)
	supplier.PUT("/:id", supplierController.Update)
	supplier.DELETE("/:id", supplierController.Delete)

	// item routes
	item := av.api.Group("/items")
	item.Use(middleware.JWTMiddleware()) // use middleware jwt general on item routes

	item.GET("", itemController.Get)
	item.GET("/:id", itemController.GetById)
	item.POST("", itemController.Create)
	item.PUT("/:id", itemController.Update)
	item.DELETE("/:id", itemController.Delete)
}
