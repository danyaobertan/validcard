package validcard

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/danyaobertan/validcard/pkg/errorops"
)

func TestValidate(t *testing.T) {
	type fields struct {
		CardNumber      string
		ExpirationMonth int
		ExpirationYear  int
	}

	tests := []struct { //nolint:govet // it's a test
		name   string
		fields fields
		want   *errorops.Error
	}{
		{
			name: "empty card number",
			fields: fields{
				CardNumber:      "",
				ExpirationMonth: 1,
				ExpirationYear:  2024,
			},
			want: errorops.NewError(
				http.StatusBadRequest,
				"card number is empty",
				nil,
			),
		},
		{
			name: "empty expiration month",
			fields: fields{
				CardNumber:      "12345678901234",
				ExpirationMonth: 0,
				ExpirationYear:  2024,
			},
			want: errorops.NewError(
				http.StatusBadRequest,
				"expiration month is empty",
				nil,
			),
		},
		{
			name: "empty expiration year",
			fields: fields{
				CardNumber:      "12345678901234",
				ExpirationMonth: 1,
				ExpirationYear:  0,
			},
			want: errorops.NewError(
				http.StatusBadRequest,
				"expiration year is empty",
				nil,
			),
		},
		{
			name: "valid request",
			fields: fields{
				CardNumber:      "12345678901234",
				ExpirationMonth: 1,
				ExpirationYear:  2024,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RequestBody{
				CardNumber:      tt.fields.CardNumber,
				ExpirationMonth: tt.fields.ExpirationMonth,
				ExpirationYear:  tt.fields.ExpirationYear,
			}
			if got := r.validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestBody.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
