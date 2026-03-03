// internal/models/tenant.go
// Defines the structure of a hosting account (Tenant).
package models

type Tenant struct {
    ID       string `json:"id"`       // Unique ID for the tenant
    Username string `json:"username"` // Linux system username
    Domain   string `json:"domain"`   // Main domain associated
    Status   string `json:"status"`   // Current state (active/suspended)
}
