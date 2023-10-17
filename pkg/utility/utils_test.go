package utility

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test case using the table driven test
func TestUtilityFunctions(t *testing.T) {
	tests := []struct {
		name              string
		functionUnderTest func(w http.ResponseWriter, r *http.Request) error
		inputRequest      *http.Request
		expectedStatus    int
		expectedResponse  string
	}{
		{
			name: "Test WriteJSON",
			functionUnderTest: func(w http.ResponseWriter, r *http.Request) error {
				data := map[string]string{"message": "Hello, World!"}
				return WriteJSON(w, http.StatusOK, data)
			},
			inputRequest:     httptest.NewRequest("GET", "/", nil),
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"Hello, World!"}`,
		},
		{
			name: "Test ReadJSON",
			functionUnderTest: func(w http.ResponseWriter, r *http.Request) error {
				var data map[string]string
				return ReadJSON(w, r, &data)
			},
			inputRequest: func() *http.Request {
				data := `{"message":"Hello, World!"}`
				return httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(data)))
			}(),
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"Hello, World!"}`,
		},
		{
			name: "Test ErrorJSON",
			functionUnderTest: func(w http.ResponseWriter, r *http.Request) error {
				err := errors.New("Something went wrong")
				return ErrorJSON(w, err, http.StatusNotFound)
			},
			inputRequest:     httptest.NewRequest("GET", "/", nil),
			expectedStatus:   http.StatusNotFound,
			expectedResponse: `{"error":true,"message":"Something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			err := test.functionUnderTest(recorder, test.inputRequest)
			if err != nil {
				t.Errorf("Function returned an error: %v", err)
			}

			if recorder.Code != test.expectedStatus {
				t.Errorf("Expected status %d, but got %d", test.expectedStatus, recorder.Code)
			}
		})
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

// Negative test cases
