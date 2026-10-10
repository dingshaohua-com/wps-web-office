package bootstrap

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterAPIServesStaticFiles(t *testing.T) {
	router := http.NewServeMux()
	RegisterAPI(router)

	request := httptest.NewRequest(http.MethodGet, "/static/jl.docx", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", response.Code, http.StatusOK)
	}
	if response.Body.Len() != 118221 {
		t.Errorf("response body size = %d, want %d", response.Body.Len(), 118221)
	}
}

func TestFileDownloadUsesMockFileURL(t *testing.T) {
	router := http.NewServeMux()
	RegisterAPI(router)

	request := httptest.NewRequest(http.MethodGet, "/office/v3/3rd/files/jl_docx/download", nil)
	request.Host = "office.example.com:9443"
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d; body = %s", response.Code, http.StatusOK, response.Body.String())
	}

	var body struct {
		Code int `json:"code"`
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	want := "http://office.example.com:9443/static/jl.docx"
	if body.Code != 0 {
		t.Errorf("body.Code = %d, want 0", body.Code)
	}
	if body.Data.URL != want {
		t.Errorf("body.Data.URL = %q, want %q", body.Data.URL, want)
	}
}
