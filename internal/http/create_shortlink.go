package http

import (
	"ab180/internal/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	shortIDMinLength = 3
	shortIDCharset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseURL          = "https://abit.ly/"
	RespTimeFormat   = "2006-01-02T15:04:05-0700"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (sh *ShortLinkHandler) CreateShortLink(c echo.Context) error {
	inputURL := c.FormValue("inputURL")
	// url 유효성 검사
	if !isValidURL(inputURL) {
		return c.String(
			http.StatusBadRequest,
			fmt.Sprintf("입력주소: %s\ninvalid url, (주소가 올바르다면 http:// 또는 https:// 를 확인해주세요.)", inputURL),
		)
	}

	var shortLinkID, createdAt string
	var charLength = shortIDMinLength
	for i := 0; ; i++ {
		shortLink := domain.ShortLink{
			ID: stringWithCharset(charLength, shortIDCharset),
			CreatedAt: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
			URL: inputURL,
		}
		err := sh.repo.CreateShortLink(context.Background(), shortLink)
		if err != nil {
			// 중복된 id 이면 재시도
			if isDuplicatedID(err) {
				// 재시도 횟수가 10의 charLength제곱만큼 반복되면 id 길이를 늘려줌
				if i > int(math.Pow(10, float64(charLength))) {
					charLength += 1
					i = 0
				}
				continue
			}
			// 내부 오류시
			return c.String(http.StatusInternalServerError, err.Error())
		}
		shortLinkID = shortLink.ID
		createdAt = shortLink.CreatedAt.Time.Format(RespTimeFormat)
		break
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": map[string]string{
			"shortId":   shortLinkID,
			"url":       inputURL,
			"createdAt": createdAt,
			"shortLink": baseURL + shortLinkID,
		},
	})
}

func isValidURL(inputURL string) bool {
	u, err := url.Parse(inputURL)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func isDuplicatedID(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
