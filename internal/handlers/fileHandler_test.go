package handlers_test

import (
	"DZ_Backend_dev_Go_level_1/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileServer(t *testing.T) {
	cases := map[string]struct {
		req      *http.Request
		rr       *httptest.ResponseRecorder
		handler  *handlers.FileHendler
		status   int
		expected string
	}{
		"Test 1": {
			req:     httptest.NewRequest("GET", "/files", nil),
			rr:      httptest.NewRecorder(),
			handler: &handlers.FileHendler{PathDir: "./test"},
			status: http.StatusOK,
			expected: `Name: testfile.txt расширение txt размер 73 байт
Name: testfile2.txt расширение txt размер 73 байт
Name: Текстовый файл.jpeg расширение jpeg размер 0 байт
Name: Туманность_Ориона.jpg расширение jpg размер 123284 байт
`,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			cs.handler.ServeHTTP(cs.rr, cs.req)
			if cs.status != cs.rr.Code {
				t.Errorf("handler returned wrong status code: got %v want %v", cs.rr.Code, cs.status)
			}

			if cs.rr.Body.String() != cs.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					cs.rr.Body.String(), cs.expected)
			}

			//client := NewClient().SetSource(NewHttpDoerMock(cs.response))
			//result, err := client.GetСategories()
			//for i := 0; i < len(cs.want); i++ {
			//	if cs.want[i] != result[i] {
			//		t.Errorf("expected: %s, actual: %s", cs.want[i], result[i])
			//	}
			//}
			//
			//if err != nil {
			//	if err.Error() != cs.errorMessage {
			//		t.Errorf("expected err: %s, actual err: %s", cs.errorMessage, err.Error())
			//	}
			//}
		})
	}

}
