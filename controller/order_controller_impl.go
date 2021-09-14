package controller

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}


func NewProductController(ProductService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productName := params.ByName("productName")

	ProductResponses := controller.ProductService.FindByName(request.Context(), productName)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}


func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	merchantId := params.ByName("merchantId")
	id, err := strconv.Atoi(merchantId)
	helper.PanicIfError(err)

	ProductResponses := controller.ProductService.FindAll(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
