package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetBooksHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/livros", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBooksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status retornado está incorreto: obtido %v, esperado %v", status, http.StatusOK)
	}

	expectedBooks := books
	var responseBooks []Book
	err = json.NewDecoder(rr.Body).Decode(&responseBooks)
	if err != nil {
		t.Fatal("Falha ao decodificar resposta")
	}

	if len(responseBooks) != len(expectedBooks) {
		t.Errorf("número incorreto de livros retornados: obtido %v, esperado %v", len(responseBooks), len(expectedBooks))
	}
}

func TestGetBookByIDHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/livros/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookByIDHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status retornado está incorreto: obtido %v, esperado %v", status, http.StatusOK)
	}

	var responseBook Book
	err = json.NewDecoder(rr.Body).Decode(&responseBook)
	if err != nil {
		t.Fatal("Falha ao decodificar resposta")
	}

	expectedBook := books[0]
	if responseBook.ID != expectedBook.ID || responseBook.Title != expectedBook.Title {
		t.Errorf("livro incorreto retornado: obtido %+v, esperado %+v", responseBook, expectedBook)
	}

	req, err = http.NewRequest("GET", "/livros/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("status incorreto para livro não encontrado: obtido %v, esperado %v", status, http.StatusNotFound)
	}
}

func TestGetBookByIDHandlerInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/livros/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookByIDHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("status incorreto para ID inválido: obtido %v, esperado %v", status, http.StatusBadRequest)
	}
}

func TestGetBookByNegativeID(t *testing.T) {
	req, err := http.NewRequest("GET", "/livros/-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookByIDHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("status incorreto para ID negativo: obtido %v, esperado %v", status, http.StatusNotFound)
	}
}

func TestGetBookByIDHandlerBoundary(t *testing.T) {
	lastBook := books[len(books)-1]
	req, err := http.NewRequest("GET", "/livros/"+strconv.Itoa(lastBook.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookByIDHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status retornado está incorreto: obtido %v, esperado %v", status, http.StatusOK)
	}

	var responseBook Book
	err = json.NewDecoder(rr.Body).Decode(&responseBook)
	if err != nil {
		t.Fatal("Falha ao decodificar resposta")
	}

	if responseBook.ID != lastBook.ID || responseBook.Title != lastBook.Title {
		t.Errorf("livro incorreto retornado: obtido %+v, esperado %+v", responseBook, lastBook)
	}
}
