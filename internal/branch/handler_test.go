package branch

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type fakeLister struct {
	branches []Branch
	listErr  error
	branch   Branch
	findErr  error
}

func (f fakeLister) List(context.Context) ([]Branch, error) {
	return f.branches, f.listErr
}

func (f fakeLister) FindByID(context.Context, string) (Branch, error) {
	return f.branch, f.findErr
}

func TestHandlerListReturnsBranches(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(fakeLister{branches: []Branch{{
		ID:       "branch-1",
		Code:     "D1",
		Name:     "Spa District 1",
		Timezone: "Asia/Ho_Chi_Minh",
	}}}, fakeLister{})

	router := gin.New()
	router.GET("/api/v1/branches", handler.List)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/branches", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}

	want := `{"data":[{"id":"branch-1","code":"D1","name":"Spa District 1","timezone":"Asia/Ho_Chi_Minh"}]}`
	if recorder.Body.String() != want {
		t.Fatalf("body = %s, want %s", recorder.Body.String(), want)
	}
}

func TestHandlerListReturnsInternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(fakeLister{listErr: errors.New("database unavailable")}, fakeLister{})

	router := gin.New()
	router.GET("/api/v1/branches", handler.List)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/branches", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusInternalServerError)
	}

	want := `{"error":"could not load branches"}`
	if recorder.Body.String() != want {
		t.Fatalf("body = %s, want %s", recorder.Body.String(), want)
	}
}

func TestHandlerGetByIDReturnsBranch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(fakeLister{}, fakeLister{branch: Branch{
		ID:       "00000000-0000-0000-0000-000000000001",
		Code:     "D1",
		Name:     "Spa District 1",
		Timezone: "Asia/Ho_Chi_Minh",
	}})

	router := gin.New()
	router.GET("/api/v1/branches/:id", handler.GetByID)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/branches/00000000-0000-0000-0000-000000000001", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}

	want := `{"data":{"id":"00000000-0000-0000-0000-000000000001","code":"D1","name":"Spa District 1","timezone":"Asia/Ho_Chi_Minh"}}`
	if recorder.Body.String() != want {
		t.Fatalf("body = %s, want %s", recorder.Body.String(), want)
	}
}

func TestHandlerGetByIDRejectsInvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(fakeLister{}, fakeLister{})

	router := gin.New()
	router.GET("/api/v1/branches/:id", handler.GetByID)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/branches/not-a-uuid", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
	}

	want := `{"error":"invalid branch id"}`
	if recorder.Body.String() != want {
		t.Fatalf("body = %s, want %s", recorder.Body.String(), want)
	}
}

func TestHandlerGetByIDReturnsNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(fakeLister{}, fakeLister{findErr: ErrNotFound})

	router := gin.New()
	router.GET("/api/v1/branches/:id", handler.GetByID)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/branches/00000000-0000-0000-0000-000000000001", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusNotFound)
	}

	want := `{"error":"branch not found"}`
	if recorder.Body.String() != want {
		t.Fatalf("body = %s, want %s", recorder.Body.String(), want)
	}
}

func TestHandlerGetByIDReturnsInternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(fakeLister{}, fakeLister{findErr: errors.New("database unavailable")})

	router := gin.New()
	router.GET("/api/v1/branches/:id", handler.GetByID)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/branches/00000000-0000-0000-0000-000000000001", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusInternalServerError)
	}

	want := `{"error":"could not load branch"}`
	if recorder.Body.String() != want {
		t.Fatalf("body = %s, want %s", recorder.Body.String(), want)
	}
}
