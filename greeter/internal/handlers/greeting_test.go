package handlers

import (
	"context"
	"testing"

	"greeter/internal/gen"
)

func TestGetGreetingWithName(t *testing.T) {
	s := NewGreetingServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{Name: "Alice"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	g, ok := resp.(gen.GetGreeting200JSONResponse)
	if !ok {
		t.Fatalf("unexpected response type: %T", resp)
	}
	if g.Name != "Alice" {
		t.Errorf("name = %q, want %q", g.Name, "Alice")
	}
	if g.Message == "" {
		t.Error("message is empty")
	}
}

func TestGetGreetingDefaultsWhenNameMissing(t *testing.T) {
	s := NewGreetingServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	g := resp.(gen.GetGreeting200JSONResponse)
	if g.Name != defaultName {
		t.Errorf("name = %q, want %q", g.Name, defaultName)
	}
}

func TestGetGreetingDefaultsWhenNameEmpty(t *testing.T) {
	s := NewGreetingServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{Name: ""},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	g := resp.(gen.GetGreeting200JSONResponse)
	if g.Name != defaultName {
		t.Errorf("name = %q, want %q", g.Name, defaultName)
	}
}
