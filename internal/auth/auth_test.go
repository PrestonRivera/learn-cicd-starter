package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := make(http.Header)
	header1.Add("Authorization", "ApiKey token1")

	header2 := make(http.Header)
	header2.Add("Authorization", "ApiKeys token2")

	header3 := make(http.Header)
	header3.Add("Authorization", "")

	header4 := make(http.Header)
	header4.Add("Authorization", " token4")

	type test struct {
		name string
		input http.Header
		want string
		wantErr bool
	}

	tests := []test{
		{
			name: "valid api key",
			input: header1,
			want: "token1",
			wantErr: false,
		},
		{
			name: "Invaldid format in value",
			input: header2,
			want: "",
			wantErr: true,
		},
		{
			name: "valid key, no value",
			input: header3,
			want: "",
			wantErr: true,
		},
		{
			name: "Invalid format, missing ApiKey",
			input: header4,
			want: "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		
		if tc.wantErr && err == nil {
			t.Errorf("%s: expected an error and got none", tc.name)
			continue
		}
		
		if !tc.wantErr && err != nil {
			t.Errorf("%s: unexpected error: %v", tc.name, err)
			continue
		}

		if got != tc.want {
			t.Errorf("%s: expected %q, got %q", tc.name, tc.want, got)
		}
	}
}