package gotesla

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Errorf("failed to new client: %s", err)
	}
}
