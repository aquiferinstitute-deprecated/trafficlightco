package trafficlightco

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetRating(t *testing.T) {
	expectedValidTime, _ := time.Parse(
		time.RFC3339,
		"2017-10-07T16:50:06.664292Z",
	)
	expectedValidRating := Rating{
		Color:     "green",
		CompanyID: "9432824c-afc3-4f52-a04c-e4662f6e4766",
		Created:   expectedValidTime,
	}
	type args struct {
		apiKey string
		cID    string
	}
	tests := []struct {
		name                string
		args                args
		usedResponseStatus  int
		usedContentType     string
		usedResponseContent []byte
		want                *Rating
		wantErr             bool
	}{
		{
			name: "Working ratings request",
			args: args{
				apiKey: "123",
				cID:    "9432824c-afc3-4f52-a04c-e4662f6e4766",
			},
			usedResponseStatus:  http.StatusCreated,
			usedContentType:     "application/json",
			usedResponseContent: testGetRatings(),
			want:                &expectedValidRating,
			wantErr:             false,
		},
		{
			name:    "Empty companyID params",
			args:    args{apiKey: "123"},
			wantErr: true,
		},
		{
			name: "http status not ok",
			args: args{
				apiKey: "123",
				cID:    "9432824c-afc3-4f52-a04c-e4662f6e4766",
			},
			usedResponseStatus: http.StatusInternalServerError,
			wantErr:            true,
		},
		{
			name: "wrong content-type",
			args: args{
				apiKey: "123",
				cID:    "9432824c-afc3-4f52-a04c-e4662f6e4766",
			},
			usedResponseStatus:  http.StatusCreated,
			usedContentType:     "application/xml",
			usedResponseContent: []byte{},
			wantErr:             true,
		},
		{
			name: "invalid JSON content",
			args: args{
				apiKey: "123",
				cID:    "9432824c-afc3-4f52-a04c-e4662f6e4766",
			},
			usedResponseStatus:  http.StatusCreated,
			usedContentType:     "application/json",
			usedResponseContent: []byte(`{`),
			wantErr:             true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("content-type", tt.usedContentType)
						w.WriteHeader(tt.usedResponseStatus)
						w.Write(tt.usedResponseContent)
					},
				),
			)
			defer ts.Close()

			c, errC := NewClient(ts.URL, tt.args.apiKey)
			if errC != nil {
				t.Fatalf("NewClient error = %v", errC)
				return
			}
			got, err := c.GetRating(tt.args.cID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetRating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testGetRatings() []byte {
	return []byte(`
		{
			"color": "green",
			"company": "9432824c-afc3-4f52-a04c-e4662f6e4766",
			"created": "2017-10-07T16:50:06.664292Z",
			"numeric": false
		}
	`)
}
