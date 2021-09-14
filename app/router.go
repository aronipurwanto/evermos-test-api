package app

import (
	"ahmadroni/test-evermos-api/controller"
	"github.com/julienschmidt/httprouter"
)

func NewCategoryRouter(categoryController controller.CategoryController, router *httprouter.Router) {
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}

func NewMerchantRouter(controller controller.MerchantController, router *httprouter.Router) {
	router.GET("/api/merchants", controller.FindAll)
	router.GET("/api/merchants/:merchantId", controller.FindById)
	router.POST("/api/merchants", controller.Create)
	router.PUT("/api/merchants/:merchantId", controller.Update)
	router.DELETE("/api/merchants/:merchantId", controller.Delete)
}

func NewCustomerRouter(controller controller.CustomerController, router *httprouter.Router) {
	router.GET("/api/customers", controller.FindAll)
	router.GET("/api/customers/:customerId", controller.FindById)
	router.POST("/api/customers", controller.Create)
	router.PUT("/api/customers/:customerId", controller.Update)
	router.DELETE("/api/customers/:customerId", controller.Delete)
}

func NewProductRouter(controller controller.ProductController, router *httprouter.Router) {
	router.GET("/api/products/merchant/:merchantId", controller.FindAll)
	router.GET("/api/products/search/:productName", controller.FindByName)
	router.GET("/api/products/:customerId", controller.FindById)
	router.POST("/api/products", controller.Create)
	router.PUT("/api/products/:customerId", controller.Update)
	router.DELETE("/api/products/:customerId", controller.Delete)
}