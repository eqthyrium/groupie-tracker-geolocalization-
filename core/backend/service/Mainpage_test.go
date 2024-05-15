package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"Adlet/core/backend/service/support"
)

func TestMainPageHandler(t *testing.T) {
	TestServer := []struct {
		Method string
		Path   string
		Answer int
	}{
		{
			Method: "GET",
			Path:   "http://localhost:8080/",
			Answer: 200,
		},
		{
			Method: "POST",
			Path:   "http://localhost:8080/",
			Answer: 405,
		},
		{
			Method: "DELETE",
			Path:   "http://localhost:8080/",
			Answer: 405,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/helloAlem",
			Answer: 404,
		},
		{
			Method: "POST",
			Path:   "http://localhost:8080/HelloAlem",
			Answer: 404,
		},
	}
	err := os.Chdir("/home/student/Task/launch")
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	support.Jsonhandler()
	for i := 0; i < len(TestServer); i++ {
		request, err := http.NewRequest(TestServer[i].Method, TestServer[i].Path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(MainPageHandler)
		handler.ServeHTTP(response, request)

		if status := response.Code; status != TestServer[i].Answer {
			t.Errorf("handler returned wrong status code: got %v want %v", status, TestServer[i].Answer)
		}
	}
}
