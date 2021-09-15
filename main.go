package main

import (
	"ahmadroni/test-evermos-api/app"
	"ahmadroni/test-evermos-api/controller"
	"ahmadroni/test-evermos-api/exception"
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/middleware"
	"ahmadroni/test-evermos-api/repository"
	"ahmadroni/test-evermos-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	app.NewCategoryRouter(categoryController, router)

	merchantRepository := repository.NewMerchantRepository()
	merchantService := service.NewMerchantService(merchantRepository, db, validate)
	merchantController := controller.NewMerchantController(merchantService)
	app.NewMerchantRouter(merchantController, router)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)
	app.NewCustomerRouter(customerController, router)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	app.NewProductRouter(productController, router)

	orderRepository := repository.NewOrderRepository()
	orderDetailRepository := repository.NewOrderDetailRepository()
	orderService := service.NewOrderService(orderRepository, orderDetailRepository, productRepository, db, validate)
	orderController := controller.NewOrderController(orderService)
	app.NewOrderRouter(orderController, router)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8084",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
