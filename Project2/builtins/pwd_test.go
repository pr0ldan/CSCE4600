// test_pwd.go
package builtins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintWorkingDirectory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []string
		wantErr bool
		wantOut string
	}{
		{
			name:    "Print working directory",
			args:    []string{},
			wantErr: false,
			// Note: Update this based on the expected output for your system.
			wantOut: "/path/to/current/directory\n",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Redirect stdout to capture the output.
			stdout := &bytes.Buffer{}

			// Run the PrintWorkingDirectory function.
			err := PrintWorkingDirectory(stdout, tt.args...)

			// Assertions
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantOut, stdout.String())
			}

			// Add additional assertions as needed.
		})
	}
}
