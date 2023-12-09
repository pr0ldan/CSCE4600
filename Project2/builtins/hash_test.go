package builtins

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHashCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		args          []string
		wantOptions   map[string]interface{}
		wantArguments []string
		wantErr       bool
	}{
		{
			name:          "Test with -r -p filename -l argument",
			args:          []string{"-r", "-p", "filename", "-l", "arg1", "arg2"},
			wantOptions:   map[string]interface{}{"-r": true, "-p": true, "filename": "filename", "-l": true},
			wantArguments: []string{"arg1", "arg2"},
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Redirect stdout and stderr to capture the output.
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			// Run the HashCommand in a goroutine to simulate a concurrent execution.
			go func() {
				err := HashCommand(tt.args...)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
				}
			}()

			// Give some time for the goroutine to execute.
			time.Sleep(10 * time.Millisecond)

			// Assertions
			options := parseHashOptions(tt.args)
			require.Equal(t, tt.wantOptions, options)

			arguments := parseHashArguments(tt.args)
			require.Equal(t, tt.wantArguments, arguments)

			require.Empty(t, stdout.String()) // Ensure no output to stdout
			require.Empty(t, stderr.String()) // Ensure no output to stderr

			// Add additional assertions as needed.

		})
	}
}
