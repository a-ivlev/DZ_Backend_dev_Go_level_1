package handlers_test

import (
	"DZ_Backend_dev_Go_level_1/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	cases := map[string]struct {
		req      *http.Request
		rr       *httptest.ResponseRecorder
		handler  *handlers.Handler
		status   int
		expected string
		//errorMessage string
	}{
		"Test 1": {
			// Создаем запрос с указанием нашего хендлера. Так как мы тестируем GET-эндпоинт
			// то нам не нужно передавать тело, поэтому третьим аргументом передаем nil
			req: httptest.NewRequest("GET", "/home/?name=John", nil),
			// Мы создаем ResponseRecorder(реализует интерфейс http.ResponseWriter)
			// и используем его для получения ответа
			rr: httptest.NewRecorder(),
			// Указываем какой хендлер будем тестировать
			handler: &handlers.Handler{},
			// Прогнозируемы код ответа
			status: http.StatusOK,
			// тело ответа
			expected: `Parsed query-param with key "name": John`,
		},
		"Test 2": {
			// Создаем запрос с указанием нашего хендлера. Так как мы тестируем GET-эндпоинт
			// то нам не нужно передавать тело, поэтому третьим аргументом передаем nil
			req: httptest.NewRequest("GET", "/home/?name=Вано", nil),
			// Мы создаем ResponseRecorder(реализует интерфейс http.ResponseWriter)
			// и используем его для получения ответа
			rr: httptest.NewRecorder(),
			// Указываем какой хендлер будем тестировать
			handler: &handlers.Handler{},
			// Прогнозируемы код ответа
			status: http.StatusOK,
			// тело ответа
			expected: `Parsed query-param with key "name": Вано`,
		},
		"Test 3": {
			// Создаем запрос с указанием нашего хендлера. Так как мы тестируем GET-эндпоинт
			// то нам не нужно передавать тело, поэтому третьим аргументом передаем nil
			req: httptest.NewRequest("GET", "/home", nil),
			// Мы создаем ResponseRecorder(реализует интерфейс http.ResponseWriter)
			// и используем его для получения ответа
			rr: httptest.NewRecorder(),
			// Указываем какой хендлер будем тестировать
			handler: &handlers.Handler{},
			// Прогнозируемы код ответа
			status: http.StatusOK,
			// тело ответа
			expected: `Parsed query-param with key "name": `,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			// Наш хендлер соответствует интерфейсу http.Handler, а значит
			// мы можем использовать ServeHTTP и напрямую указать
			// Request и ResponseRecorder
			cs.handler.ServeHTTP(cs.rr, cs.req)
			// Проверяем статус-код ответа
			if http.StatusOK != cs.rr.Code {
				t.Errorf("handler returned wrong status code: got %v want %v", cs.rr.Code, http.StatusOK)
			}
			// Проверяем тело ответа
			if cs.rr.Body.String() != cs.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					cs.rr.Body.String(), cs.expected)
			}
		})
	}
}

//func TestGetHandler(t *testing.T) {
//	// Создаем запрос с указанием нашего хендлера. Так как мы тестируем GET-эндпоинт
//	// то нам не нужно передавать тело, поэтому третьим аргументом передаем nil
//	req, err := http.NewRequest("GET", "/?name=John", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Мы создаем ResponseRecorder(реализует интерфейс http.ResponseWriter)
//	// и используем его для получения ответа
//	rr := httptest.NewRecorder()
//	handler := &handlers.Handler{}
//
//	// Наш хендлер соответствует интерфейсу http.Handler, а значит
//	// мы можем использовать ServeHTTP и напрямую указать
//	// Request и ResponseRecorder
//	handler.ServeHTTP(rr, req)
//
//	// Проверяем статус-код ответа
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	// Проверяем тело ответа
//	expected := `Parsed query-param with key "name": John`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//}
