package core

import "testing"

func TestNewEngine(t *testing.T) {
    version := "1.0.0"
    engine := NewEngine(version)
    
    if engine.Version != version {
        t.Errorf("Expected version %s, got %s", version, engine.Version)
    }
}
