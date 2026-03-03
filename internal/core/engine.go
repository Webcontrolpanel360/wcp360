package core

import (
    "fmt"
    "github.com/Webcontrolpanel360/wcp360/internal/models"
    "github.com/Webcontrolpanel360/wcp360/internal/provisioner"
)

type Engine struct {
    Version string
}

func NewEngine(v string) *Engine {
    return &Engine{Version: v}
}

func (e *Engine) ProvisionTenant(t models.Tenant) error {
    fmt.Printf("[CORE] Starting orchestration for: %s\n", t.Username)
    
    // Call the Web Provisioner
    err := provisioner.SetupWebsite(t)
    if err != nil {
        return err
    }
    
    return nil
}
