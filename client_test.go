package trafficlightco

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	u, errP := url.Parse("http://test.com")
	if errP != nil {
		t.Fatalf("TestNewClient: %v", errP)
	}
	tests := []struct {
		name    string
		u       string
		key     string
		want    *Client
		wantErr bool
	}{
		{
			name: "Working client",
			u:    "http://test.com",
			key:  "123",
			want: &Client{
				apiURL: u,
				apiKey: "123",
				httpClient: &http.Client{
					Timeout: apiTimeout * time.Second,
				},
			},
			wantErr: false,
		},
		{
			name:    "Empty API key",
			u:       "http://test.com",
			key:     "",
			wantErr: true,
		},
		{
			name:    "Non working URL",
			u:       ":",
			key:     "123",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.u, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
