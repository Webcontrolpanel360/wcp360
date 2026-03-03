package waf

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func TestWAFBlocking(t *testing.T) {
    // Setup
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/?search=DROP TABLE", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    
    handler := Middleware()(func(c echo.Context) error {
        return c.String(http.StatusOK, "Passed")
    })

    // Execute
    err := handler(c)

    // Assertions
    if err == nil {
        assert.Equal(t, http.StatusForbidden, rec.Code)
    }
}
