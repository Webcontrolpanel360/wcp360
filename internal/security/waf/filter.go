// internal/security/waf/filter.go
// This module inspects incoming requests for malicious patterns.
package waf

import (
    "strings"
    "github.com/labstack/echo/v4"
    "net/http"
)

// Middleware returns an Echo middleware that blocks suspicious traffic.
func Middleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // Check query parameters for common SQL Injection patterns
            query := c.QueryString()
            maliciousPatterns := []string{"UNION SELECT", "DROP TABLE", "--", "OR 1=1"}
            
            for _, pattern := range maliciousPatterns {
                if strings.Contains(strings.ToUpper(query), pattern) {
                    return c.JSON(http.StatusForbidden, map[string]string{
                        "error": "Security Breach: Malicious pattern detected by WCP360 WAF",
                    })
                }
            }
            return next(c)
        }
    }
}
