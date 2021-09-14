package controller

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(OrderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: OrderService,
	}
}


func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(request, &OrderCreateRequest)

	OrderResponse := controller.OrderService.Create(request.Context(), OrderCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   OrderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
