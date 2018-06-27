package trafficlightco

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_Search(t *testing.T) {
	expectedValidCompany := Company{
		Name: "35° EAST",
		ID:   "9432824c-afc3-4f52-a04c-e4662f6e4766",
	}
	type args struct {
		apiKey string
		co     string
		na     string
		iType  string
		iValue string
	}
	tests := []struct {
		name                string
		args                args
		usedResponseStatus  int
		usedContentType     string
		usedResponseContent []byte
		want                []Company
		wantErr             bool
	}{
		{
			name: "Working search request",
			args: args{
				apiKey: "123",
				co:     "FR",
				na:     "35° EAST",
				iType:  "",
				iValue: "",
			},
			usedResponseStatus:  http.StatusOK,
			usedContentType:     "application/json",
			usedResponseContent: testGetCompanies(),
			want:                []Company{expectedValidCompany},
			wantErr:             false,
		},
		{
			name:    "Empty search params",
			args:    args{apiKey: "123"},
			wantErr: true,
		},
		{
			name: "http status not ok",
			args: args{
				apiKey: "123",
				co:     "FR",
				na:     "35° EAST",
				iType:  "",
				iValue: "",
			},
			usedResponseStatus: http.StatusInternalServerError,
			wantErr:            true,
		},
		{
			name: "wrong content-type",
			args: args{
				apiKey: "123",
				co:     "FR",
				na:     "35° EAST",
				iType:  "",
				iValue: "",
			},
			usedResponseStatus:  http.StatusOK,
			usedContentType:     "application/xml",
			usedResponseContent: []byte{},
			wantErr:             true,
		},
		{
			name: "invalid JSON content",
			args: args{
				apiKey: "123",
				co:     "FR",
				na:     "35° EAST",
				iType:  "",
				iValue: "",
			},
			usedResponseStatus:  http.StatusOK,
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
						w.Write(tt.usedResponseContent)
						w.WriteHeader(tt.usedResponseStatus)
					},
				),
			)
			defer ts.Close()
			c, errC := NewClient(ts.URL, tt.args.apiKey)
			if errC != nil {
				t.Fatalf("NewClient error = %v", errC)
				return
			}
			got, err := c.Search(tt.args.co, tt.args.na, tt.args.iType, tt.args.iValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testGetCompanies() []byte {
	return []byte(`
		{
			"results": [
				{
					"name": "35° EAST",
					"id": "9432824c-afc3-4f52-a04c-e4662f6e4766"
				}
			],
			"previous": null,
			"next": null
		}
	`)
}
