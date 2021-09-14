package helper

import (
	"ahmadroni/test-evermos-api/model/domain"
	"ahmadroni/test-evermos-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToMerchantResponse(data domain.Merchant) web.MerchantResponse {
	return web.MerchantResponse{
		Id:   data.Id,
		Name: data.Name,
		Email: data.Email,
		Address: data.Address,
		Rating: data.Rating,
	}
}

func ToMerchantResponses(datas []domain.Merchant) []web.MerchantResponse {
	var responses []web.MerchantResponse
	for _, data := range datas {
		responses = append(responses, ToMerchantResponse(data))
	}
	return responses
}

func ToCustomerResponse(data domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:   data.Id,
		Name: data.Name,
		Email: data.Email,
		Address: data.Address,
		PhoneNumber: data.PhoneNumber,
	}
}

func ToCustomerResponses(datas []domain.Customer) []web.CustomerResponse {
	var responses []web.CustomerResponse
	for _, data := range datas {
		responses = append(responses, ToCustomerResponse(data))
	}
	return responses
}

func ToProductResponse(data domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:   data.Id,
		MerchantId: data.MerchantId,
		CategoryId: data.CategoryId,
		Name: data.Name,
		ImagePath: data.ImagePath,
		Price: data.Price,
		Stock: data.Stock,
	}
}

func ToProductResponses(datas []domain.Product) []web.ProductResponse {
	var responses []web.ProductResponse
	for _, data := range datas {
		responses = append(responses, ToProductResponse(data))
	}
	return responses
}
