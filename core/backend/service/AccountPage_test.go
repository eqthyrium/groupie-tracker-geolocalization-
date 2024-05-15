package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"Adlet/core/backend/service/support"
)

func TestAccountPage(t *testing.T) {
	TestServer := []struct {
		Method string
		Path   string
		Answer int
	}{
		{
			Method: "GET",
			Path:   "http://localhost:8080/account",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=-1",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=53",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=%",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=A",
			Answer: 400,
		},
		{
			Method: "POST",
			Path:   "http://localhost:8080/account",
			Answer: 405,
		},
		{
			Method: "DELETE",
			Path:   "http://localhost:8080/account",
			Answer: 405,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=53",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=0",
			Answer: 200,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/account?Image_Index_Value=20",
			Answer: 200,
		},
	}

	err := os.Chdir("/home/student/Task/launch")
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	support.Jsonhandler()
	// Here we are checking the wanted and got status code
	for i := 0; i < len(TestServer); i++ {
		request, err := http.NewRequest(TestServer[i].Method, TestServer[i].Path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(AccountPageHandler)
		handler.ServeHTTP(response, request)

		if status := response.Code; status != TestServer[i].Answer {
			t.Errorf("handler returned wrong status code: got %v want %v", status, TestServer[i].Answer)
		}
	}

	// We are going to compare the content of the repsonding html, we are going to take under consideration only last two cases which are suppoerted in the above struct
	var file1, file2 string
	for i := len(TestServer) - 2; i < len(TestServer); i++ {
		request, err := http.NewRequest(TestServer[i].Method, TestServer[i].Path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(AccountPageHandler)
		handler.ServeHTTP(response, request)
		if i == len(TestServer)-2 {
			file1 = response.Body.String()
		} else {
			file2 = response.Body.String()
		}

	}

	filePath1 := "../core/backend/service/file/file1.txt"
	filePath2 := "../core/backend/service/file/file2.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileContent1 := string(data)
	data, err = ioutil.ReadFile(filePath2)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileContent2 := string(data)

	if !(fileContent1 == file1 && fileContent2 == file2) {
		t.Errorf("There is wrong output")
	}
}
