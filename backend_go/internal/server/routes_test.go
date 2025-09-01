package server

import (
	"backend_go/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	s := &Server{}
	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", expected, string(body))
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}

func TestUserSignupHandler(t *testing.T) {
	// Create a mock server with proper initialization
	s := &Server{}
	server := httptest.NewServer(http.HandlerFunc(s.userSignupHandler))
	defer server.Close()

	// Test data
	testData := `{"username": "testuser", "email": "test@example.com", "password": "password", "first_name": "John", "last_name": "Doe"}`

	resp, err := http.Post(server.URL, "application/json", strings.NewReader(testData))
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status created; got %v", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}

	// Parse response
	var response models.SignupResponse
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("error unmarshalling response. Err: %v", err)
	}

	// Basic validation of response structure
	if response.Message == "" {
		t.Error("expected message to be present")
	}

	if response.UserID == "" {
		t.Error("expected user_id to be present")
	}

	// Check that message matches expected
	expectedMessage := "User Created succesfully"
	if response.Message != expectedMessage {
		t.Errorf("expected message to be %v; got %v", expectedMessage, response.Message)
	}

	// Note: We don't test database operations in unit tests
	// The actual database integration should be tested separately
}
