package lab2

import (
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  bool
	}{
		{
			name:           "Valid Expression",
			input:          "+ 2 2\n",
			expectedOutput: "2 + 2",
			expectedError:  false,
		},
		{
			name:           "Valid Expression 2",
			input:          "+ 5 * + 4 2 ^ 3 2\n",
			expectedOutput: "5 + (4 + 2) * (3 ^ 2)",
			expectedError:  false,
		},
		{
			name:          "Invalid Expression",
			input:         "foo bar\n",
			expectedError: true,
		},
		{
			name:          "Empty Expression",
			input:         "",
			expectedError: true,
		},
		{
			name:          "Incomplete Expression",
			input:         "+ 5\n",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			output := &strings.Builder{}
			handler := &ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			if tc.expectedError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %s", err)
				}
				actualOutput := output.String()
				if actualOutput != tc.expectedOutput {
					t.Errorf("Expected output '%s', but got '%s'", tc.expectedOutput, actualOutput)
				}
			}
		})
	}
}
