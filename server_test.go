package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HomeHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestArticlesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/tech/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArticlesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ArticlesHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestContactHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/contact/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ContactHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ContactHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMainServerIntegration(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:42069/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)
}

func TestRenderPage(t *testing.T) {
	rr := httptest.NewRecorder()
	tmplFile := "templates/home.html"
	pageTitle := "Home"

	renderPage(rr, tmplFile, pageTitle)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("renderPage returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

