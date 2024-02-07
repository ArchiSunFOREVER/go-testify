package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerValidRequest(t *testing.T) {
	// Создаем GET-запрос с корректными параметрами
	req, err := http.NewRequest("GET", "/?city=moscow&count=3", nil)
	require.NoError(t, err)

	// Создаем записывающее устройство для записи ответа сервера
	responseRecorder := httptest.NewRecorder()

	// Создаем обработчик для тестирования
	handler := http.HandlerFunc(mainHandle)

	// Вызываем обработчик с созданным запросом и записываем ответ
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус код ответа равен ожидаемому http.StatusOK (200)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWrongCityValue(t *testing.T) {
	// Создаем GET-запрос с несуществующим городом
	req, err := http.NewRequest("GET", "/?city=unknowncity&count=3", nil)
	require.NoError(t, err)

	// Создаем записывающее устройство для записи ответа сервера
	responseRecorder := httptest.NewRecorder()

	// Создаем обработчик для тестирования
	handler := http.HandlerFunc(mainHandle)

	// Вызываем обработчик с созданным запросом и записываем ответ
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус код ответа равен ожидаемому http.StatusBadRequest (400)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем, что тело ответа содержит ожидаемое сообщение об ошибке
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerCountGreaterThanTotal(t *testing.T) {
	// Создаем GET-запрос с параметром count больше, чем общее количество кафе
	req, err := http.NewRequest("GET", "/?city=moscow&count=10", nil)
	require.NoError(t, err)

	// Создаем записывающее устройство для записи ответа сервера
	responseRecorder := httptest.NewRecorder()

	// Создаем обработчик для тестирования
	handler := http.HandlerFunc(mainHandle)

	// Вызываем обработчик с созданным запросом и записываем ответ
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус код ответа равен ожидаемому http.StatusOK (200)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Получаем список всех кафе для города "moscow"
	cafes := cafeList["moscow"]

	// Проверяем, что количество возвращенных кафе равно общему количеству кафе для данного города
	assert.Equal(t, len(cafes), strings.Count(responseRecorder.Body.String(), ",")+1)
}
