package controller

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type MerchantControllerImpl struct {
	MerchantService service.MerchantService
}

func NewMerchantController(MerchantService service.MerchantService) MerchantController {
	return &MerchantControllerImpl{
		MerchantService: MerchantService,
	}
}

func (controller *MerchantControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	merchantCreateRequest := web.MerchantCreateRequest{}
	helper.ReadFromRequestBody(request, &merchantCreateRequest)

	MerchantResponse := controller.MerchantService.Create(request.Context(), merchantCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   MerchantResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	merchantUpdateRequest := web.MerchantUpdateRequest{}
	helper.ReadFromRequestBody(request, &merchantUpdateRequest)

	merchantId := params.ByName("merchantId")
	id, err := strconv.Atoi(merchantId)
	helper.PanicIfError(err)

	merchantUpdateRequest.Id = id

	MerchantResponse := controller.MerchantService.Update(request.Context(), merchantUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   MerchantResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	merchantId := params.ByName("merchantId")
	id, err := strconv.Atoi(merchantId)
	helper.PanicIfError(err)

	controller.MerchantService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	merchantId := params.ByName("merchantId")
	id, err := strconv.Atoi(merchantId)
	helper.PanicIfError(err)

	MerchantResponse := controller.MerchantService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   MerchantResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	merchantResponses := controller.MerchantService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   merchantResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
