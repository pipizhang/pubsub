package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestIPWhitelist(t *testing.T) {
	handle := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	h := IPWhitelist([]string{"127.0.0.1"})(handle)
	he := h(c).(*echo.HTTPError)

	// Check Whitelist
	assert.Error(t, h(c))
	assert.Equal(t, he.Code, http.StatusForbidden)

	// Check Enable
	m := IPWhitelistWithConfig(IPWhitelistConfig{
		Enable: false,
		List:   []string{"127.0.0.1"},
	})(handle)
	assert.NoError(t, m(c))
}
