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
		{"should pass with raw options [-sort files -order desc]", []string{"-sort", "files", "-order", "desc"},
			map[string]string{"sort": "files", "order": "desc"}, ExitCodeSuccess, false},
		{"should pass with raw options [-sort code]", []string{"-sort", "code"},
			map[string]string{"sort": "code"}, ExitCodeSuccess, false},
		{"should pass with raw options [-order asc]", []string{"-order", "asc"},
			map[string]string{"order": "asc"}, ExitCodeSuccess, false},
		{"should respond error with raw options [-error asc]", []string{"-error", "asc"},
			nil, ExitCodeFailed, true},
		{"should respond empty value with empty raw options", []string{},
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
