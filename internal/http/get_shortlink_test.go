package http

import (
	"ab180/internal/domain"
	mock_http "ab180/internal/http/mock"
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetShortLink(t *testing.T) {
	var (
		inputShortLinkID = "aaa"
		testURL          = "http://google.com"
	)

	form := make(map[string][]string)
	form["inputShortLinkID"] = []string{inputShortLinkID}

	ctrl := gomock.NewController(t)
	repo := mock_http.NewMockShortLinkRepo(ctrl)
	repo.
		EXPECT().
		GetURLByShortLink(context.Background(), domain.ShortLink{ID: inputShortLinkID}).
		Return(nil, domain.ShortLink{
			ID:        inputShortLinkID,
			CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			URL:       testURL,
		})
	h := NewShortLinkHandler(repo, nil, "")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/short-links", nil)
	req.Form = form
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if err := h.GetShortLink(ctx); err != nil {
		t.Fatalf("Expected not error")
	} else {
		if ctx.Response().Status != http.StatusOK {
			t.Fatalf("Expected http status code is %d, got %d", http.StatusOK, ctx.Response().Status)
		}
	}
}
