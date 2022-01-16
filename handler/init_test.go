package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	m.Run()
}

func createContext(method, uri string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	var req *http.Request
	req, _ = http.NewRequestWithContext(context.TODO(), method, uri, strings.NewReader(""))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}

func parseResponseBody(recBody io.Reader, res interface{}) interface{} {
	body, _ := io.ReadAll(recBody)

	var objmap *json.RawMessage

	json.Unmarshal(body, &objmap)

	return json.Unmarshal(*objmap, res)
}
