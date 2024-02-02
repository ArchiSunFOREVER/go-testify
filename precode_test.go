package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/?city=moscow&count=10", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка статус кода
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверка, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())

	// Преобразование ответа в массив строк
	cafes := strings.Split(responseRecorder.Body.String(), ",")

	// Проверка, что количество элементов в ответе равно totalCount
	assert.Len(t, cafes, totalCount)
}
