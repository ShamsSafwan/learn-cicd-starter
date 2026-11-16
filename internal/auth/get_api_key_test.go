package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "missing authorization header",
			headers: http.Header{},
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer token"},
			},
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "valid authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			want: "test-api-key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() key = %q, want %q", got, tt.want)
			}
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("GetAPIKey() error = %v, want nil", err)
				}
				return
			}
			if err == nil || err.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
