package http

import (
	"ab180/internal/domain"
	mock_http "ab180/internal/http/mock"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateShortLink(t *testing.T) {
	var inputURL = "http://google.com"
	form := make(map[string][]string)
	form["inputURL"] = []string{inputURL}

	ctrl := gomock.NewController(t)
	repo := mock_http.NewMockShortLinkRepo(ctrl)
	repo.EXPECT().CreateShortLink(context.Background(), domain.ShortLink{}).Return(nil)
	h := NewShortLinkHandler(repo, nil, "")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/short-links", nil)
	req.Form = form
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if err := h.CreateShortLink(ctx); err != nil {
		t.Fatalf("Expected not error")
	} else {
		if ctx.Response().Status != http.StatusCreated {
			t.Fatalf("Expected http status code is %d, got %d", http.StatusCreated, ctx.Response().Status)
		}
	}
}

func TestCreateShortLink_EmptyURL(t *testing.T) {
	var inputURL = ""
	form := make(map[string][]string)
	form["inputURL"] = []string{inputURL}

	h := NewShortLinkHandler(nil, nil, "")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/short-links", nil)
	req.Form = form
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if err := h.CreateShortLink(ctx); err != nil {
		t.Fatalf("Expected not error")
	} else {
		if ctx.Response().Status != http.StatusBadRequest {
			t.Fatalf("Expected http status code is %d, got %d", http.StatusBadRequest, ctx.Response().Status)
		}
	}
}

func TestCreateShortLink_InvalidURL(t *testing.T) {
	var inputURL = "ht://www.google.com"
	form := make(map[string][]string)
	form["inputURL"] = []string{inputURL}

	h := NewShortLinkHandler(nil, nil, "")

	req := httptest.NewRequest(http.MethodPost, "/short-links", nil)
	req.Form = form
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)

	if err := h.CreateShortLink(ctx); err != nil {
		t.Fatalf("Expected not error")
	} else {
		if ctx.Response().Status != http.StatusBadRequest {
			t.Fatalf("Expected http status code is %d, got %d", http.StatusBadRequest, ctx.Response().Status)
		}
	}
}
