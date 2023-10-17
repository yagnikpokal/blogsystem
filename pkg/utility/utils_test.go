package utility

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Create a test JSONResponse for testing
var testResponse = JSONResponse{
	Error:   false,
	Message: "Success",
	Data:    "Some data",
}

func TestWriteJSON(t *testing.T) {
	// Create an HTTP response recorder
	w := httptest.NewRecorder()

	// Call WriteJSON to write the test JSONResponse
	err := WriteJSON(w, http.StatusOK, testResponse)

	// Check for errors
	if err != nil {
		t.Errorf("WriteJSON returned an error: %v", err)
	}

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedJSON, _ := json.Marshal(testResponse)
	if strings.TrimSpace(w.Body.String()) != string(expectedJSON) {
		t.Errorf("Response body does not match the expected JSON")
	}
}

func TestReadJSON(t *testing.T) {
	// Create a JSON string
	jsonStr := `{"error":false,"message":"Success","data":"Some data"}`

	// Create an HTTP request with the JSON string
	r, _ := http.NewRequest("POST", "/test", strings.NewReader(jsonStr))

	// Create a test data structure to decode into
	var testData JSONResponse

	// Call ReadJSON to decode the request
	err := ReadJSON(nil, r, &testData)

	// Check for errors
	if err != nil {
		t.Errorf("ReadJSON returned an error: %v", err)
	}

	// Check the decoded data
	if testData.Error || testData.Message != "Success" || testData.Data != "Some data" {
		t.Errorf("Decoded data does not match the expected JSONResponse")
	}
}

func TestErrorJSON(t *testing.T) {
	// Create an HTTP response recorder
	w := httptest.NewRecorder()

	// Create an error to send in the response
	testError := errors.New("Test error")

	// Call ErrorJSON to write an error response
	err := ErrorJSON(w, testError, http.StatusNotFound)

	// Check for errors
	if err != nil {
		t.Errorf("ErrorJSON returned an error: %v", err)
	}

	// Check the response status code
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	// Check the response body
	expectedJSON, _ := json.Marshal(JSONResponse{Error: true, Message: "Test error"})
	if strings.TrimSpace(w.Body.String()) != string(expectedJSON) {
		t.Errorf("Response body does not match the expected JSON")
	}
}
func TestWriteJSONWithHeaders(t *testing.T) {
	// Create an HTTP response recorder
	w := httptest.NewRecorder()

	// Define custom headers
	customHeaders := http.Header{
		"X-Custom-Header": []string{"Value1"},
		"Another-Header":  []string{"Value2"},
	}

	// Call WriteJSON with custom headers
	data := JSONResponse{
		Error:   false,
		Message: "Success",
		Data:    "Some data",
	}

	err := WriteJSON(w, http.StatusOK, data, customHeaders)

	// Check for errors
	if err != nil {
		t.Errorf("WriteJSON returned an error: %v", err)
	}

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response headers
	for key, values := range customHeaders {
		expectedValues := values[0] // For simplicity, assume only one value per header
		if w.Header().Get(key) != expectedValues {
			t.Errorf("Header %s does not match the expected value %s", key, expectedValues)
		}
	}
}
