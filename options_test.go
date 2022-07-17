package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRawOptions(t *testing.T) {
	tests := []struct {
		name       string
		arg        []string
		expectOpts map[string]string
		expectCode int
		wantErr    bool
	}{
		{"[0]should pass with correct raw options", []string{"-sort", "files", "-order", "desc"},
			map[string]string{"sort": "files", "order": "desc"}, ExitCodeSuccess, false},
		{"[1]should pass with correct raw options", []string{"-sort", "code"},
			map[string]string{"sort": "code"}, ExitCodeSuccess, false},
		{"[2]should pass with correct raw options", []string{"-order", "asc"},
			map[string]string{"order": "asc"}, ExitCodeSuccess, false},
		{"[3]should respond error with incorrect raw options", []string{"-error", "asc"},
			nil, ExitCodeFailed, true},
		{"[4]should respond empty value with empty raw options", []string{},
			map[string]string{}, ExitCodeSuccess, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, code, err := parseRawOptions(tt.arg)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, opts, tt.expectOpts)
			assert.Equal(t, tt.expectCode, code)
		})
	}
}
