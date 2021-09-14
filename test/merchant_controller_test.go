package test

import (
	"ahmadroni/test-evermos-api/app"
	"ahmadroni/test-evermos-api/controller"
	"ahmadroni/test-evermos-api/middleware"
	"ahmadroni/test-evermos-api/model/domain"
	"ahmadroni/test-evermos-api/repository"
	"ahmadroni/test-evermos-api/service"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func setupMerchantRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	MerchantRepository := repository.NewMerchantRepository()
	MerchantService := service.NewMerchantService(MerchantRepository, db, validate)
	MerchantController := controller.NewMerchantController(MerchantService)
	router := httprouter.New()
	app.NewMerchantRouter(MerchantController, router)

	return middleware.NewAuthMiddleware(router)
}

func truncateMerchant(db *sql.DB) {
	db.Exec("TRUNCATE Merchant")
}

func TestCreateMerchantSuccess(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)
	router := setupMerchantRouter(db)

	requestBody := strings.NewReader(`{"name" : "Merchant Test", "email":"email@gmail.com","address":"Bandung","rating":5}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/merchants", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Merchant Test", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateMerchantFailed(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)
	router := setupMerchantRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/merchants", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateMerchantSuccess(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)

	tx, _ := db.Begin()
	MerchantRepository := repository.NewMerchantRepository()
	Merchant := MerchantRepository.Save(context.Background(), tx, domain.Merchant{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupMerchantRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/merchants/"+strconv.Itoa(Merchant.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, Merchant.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateMerchantFailed(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)

	tx, _ := db.Begin()
	MerchantRepository := repository.NewMerchantRepository()
	Merchant := MerchantRepository.Save(context.Background(), tx, domain.Merchant{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupMerchantRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/merchants/"+strconv.Itoa(Merchant.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetMerchantSuccess(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)

	tx, _ := db.Begin()
	repository := repository.NewMerchantRepository()
	merchant := repository.Save(context.Background(), tx, domain.Merchant{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupMerchantRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/merchants/"+strconv.Itoa(merchant.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, merchant.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, merchant.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetMerchantFailed(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)
	router := setupMerchantRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/merchants/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteMerchantSuccess(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)

	tx, _ := db.Begin()
	MerchantRepository := repository.NewMerchantRepository()
	Merchant := MerchantRepository.Save(context.Background(), tx, domain.Merchant{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupMerchantRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/merchants/"+strconv.Itoa(Merchant.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteMerchantFailed(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)
	router := setupMerchantRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/merchants/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListMerchantsSuccess(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)

	tx, _ := db.Begin()
	MerchantRepository := repository.NewMerchantRepository()
	Merchant1 := MerchantRepository.Save(context.Background(), tx, domain.Merchant{
		Name: "Gadget",
	})
	Merchant2 := MerchantRepository.Save(context.Background(), tx, domain.Merchant{
		Name: "Computer",
	})
	tx.Commit()

	router := setupMerchantRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/merchants", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var Merchants = responseBody["data"].([]interface{})

	MerchantResponse1 := Merchants[0].(map[string]interface{})
	MerchantResponse2 := Merchants[1].(map[string]interface{})

	assert.Equal(t, Merchant1.Id, int(MerchantResponse1["id"].(float64)))
	assert.Equal(t, Merchant1.Name, MerchantResponse1["name"])

	assert.Equal(t, Merchant2.Id, int(MerchantResponse2["id"].(float64)))
	assert.Equal(t, Merchant2.Name, MerchantResponse2["name"])
}

func TestMerchantUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateMerchant(db)
	router := setupMerchantRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/merchants", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
