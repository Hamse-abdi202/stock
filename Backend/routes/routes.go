package routes

import (
	handlers "Github.com/Hams/abdi/InventoryManagmentSystem/handler"
	"Github.com/Hams/abdi/InventoryManagmentSystem/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	userHandler := handlers.RegisterUserHandler()
	categoryHandler := handlers.RegisterCategoryHandler()
	productHandler := handlers.RegisterProductHandler()
	stockMovementHandler := handlers.RegisterStockMovementHandler()

	usersGroup := apiGroup.Group("/users")
	{
		usersGroup.POST("/create", userHandler.CreateUser)
		usersGroup.POST("/login", userHandler.LoginUser)
		usersGroup.GET("/whoami", userHandler.WhoAmI)
		usersGroup.POST("/refresh-token", userHandler.RefreshToken)
	}

	categoriesGroup := apiGroup.Group("/categories")
	{
		categoriesGroup.POST("/create", categoryHandler.CreateCategory)
		categoriesGroup.GET("/list", categoryHandler.GetAllCategories)
		categoriesGroup.GET("/list/:id", categoryHandler.GetCategoryByID)
		categoriesGroup.PUT("/update/:id", categoryHandler.UpdateCategory)
		categoriesGroup.DELETE("/delete/:id", categoryHandler.DeleteCategory)
	}

	productsGroup := apiGroup.Group("/products")
	{
		productsGroup.POST("/create", productHandler.CreateProduct)
		productsGroup.GET("/list", productHandler.GetAllProducts)
		productsGroup.GET("/list/:id", productHandler.GetProductByID)
		productsGroup.PUT("/update/:id", productHandler.UpdateProduct)
		productsGroup.DELETE("/delete/:id", productHandler.DeleteProduct)
	}

	stockGroup := apiGroup.Group("/stock")
	{
		stockGroup.POST("/create",middleware.Authenticate(), stockMovementHandler.CreateStockMovement)
		stockGroup.GET("/list", stockMovementHandler.GetAllStockMovements)
		stockGroup.GET("/list/:id", stockMovementHandler.GetStockMovementByID)
		stockGroup.GET("/product/:productID", stockMovementHandler.GetProductStockMovements)
	}
}
