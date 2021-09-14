package controller

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(CustomerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: CustomerService,
	}
}

func (controller *CustomerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerCreateRequest := web.CustomerCreateRequest{}
	helper.ReadFromRequestBody(request, &customerCreateRequest)

	customerResponse := controller.CustomerService.Create(request.Context(), customerCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerUpdateRequest := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(request, &customerUpdateRequest)

	CustomerId := params.ByName("CustomerId")
	id, err := strconv.Atoi(CustomerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	customerResponse := controller.CustomerService.Update(request.Context(), customerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse := controller.CustomerService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responses := controller.CustomerService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
