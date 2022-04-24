package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/rymiyamoto/memer-server/model"
	"github.com/rymiyamoto/memer-server/response"
	"github.com/rymiyamoto/memer-server/usecase"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld_NewHelloWorld(t *testing.T) {
	t.Parallel()

	got := NewHelloWorld()
	assert.Equal(t, got, &HelloWorld{
		helloWorldUsecase: usecase.NewHelloWorld(),
	})
}

func TestHelloWorld_Get(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	helloWorldUsecase := usecase.NewMockIHelloWorld(mockCtrl)
	h := &HelloWorld{
		helloWorldUsecase: helloWorldUsecase,
	}
	m := &model.HelloWorld{
		Text: "test",
	}

	doTest := func() (*httptest.ResponseRecorder, error) {
		c, rec := createContext(echo.GET, "/hello_world")
		err := h.Index(c)

		return rec, err
	}

	helloWorldUsecase.EXPECT().
		Get().
		Return(m)

	rec, err := doTest()
	assert.Nil(t, err)
	assert.Equal(t, rec.Code, http.StatusOK)

	var res struct {
		HelloWorld response.HelloWorld `json:"hello_world"`
	}

	parseResponseBody(rec.Body, &res)
	assert.Equal(t, res.HelloWorld.Text, m.Text)
}
