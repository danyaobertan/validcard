package validcard

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/danyaobertan/validcard/internal/api/domain"
	"github.com/danyaobertan/validcard/pkg/errorops"
)

func TestValidateTransform(t *testing.T) {
	type args struct {
		number string
	}

	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "valid card number",
			args: args{number: "5131685483769062"},
			want: []int{5, 1, 3, 1, 6, 8, 5, 4, 8, 3, 7, 6, 9, 0, 6, 2},
		},
		{
			name: "valid card number with spaces",
			args: args{number: "5131 6854 8376 9062"},
			want: []int{5, 1, 3, 1, 6, 8, 5, 4, 8, 3, 7, 6, 9, 0, 6, 2},
		},
		{
			name:    "invalid card number with dashes",
			args:    args{number: "5131-6854-8376-9062"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid card number with unexpected character",
			args:    args{number: "5131-6854-8376-9062A"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateTransform(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateTransform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidCardNumber(t *testing.T) {
	type args struct {
		cardNumber string
	}

	tests := []struct { //nolint:govet // it's a test
		name  string
		args  args
		want  bool
		want1 *errorops.Error
	}{
		{
			name: "length less than minimum",
			args: args{cardNumber: "1234567890"},
			want: false,
			want1: errorops.NewError(
				400,
				"invalid card number length",
				"card number length must be between 13 and 19, but current length is 10",
			),
		},
		{
			name: "length more than maximum",
			args: args{cardNumber: "12345678901234567901"},
			want: false,
			want1: errorops.NewError(
				400,
				"invalid card number length",
				"card number length must be between 13 and 19, but current length is 20",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isValidCardNumber(tt.args.cardNumber)
			if got != tt.want {
				t.Errorf("isValidCardNumber() got = _%v_, want _%v_", got, tt.want)
			}

			if !errorsEqual(got1, tt.want1) {
				t.Errorf("isValidCardNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsValidExpDate(t *testing.T) {
	type args struct {
		month int
		year  int
	}

	tests := []struct { //nolint:govet // it's a test
		name  string
		args  args
		want  bool
		want1 *errorops.Error
	}{
		{
			name: "invalid month",
			args: args{month: 0, year: 2021},
			want: false,
			want1: errorops.NewError(
				400,
				"invalid month",
				"month must be between 1 and 12",
			),
		},
		{
			name: "invalid year",
			args: args{month: 1, year: 9999},
			want: false,
			want1: errorops.NewError(
				400,
				"invalid year",
				fmt.Sprintf("year must be between %d and %d", time.Now().Year(), time.Now().Year()+domain.MaxYearsInFuture),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isValidExpDate(tt.args.month, tt.args.year)
			if got != tt.want {
				t.Errorf("isValidCardNumber() got = _%v_, want _%v_", got, tt.want)
			}

			if !errorsEqual(got1, tt.want1) {
				t.Errorf("isValidCardNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func errorsEqual(e1, e2 *errorops.Error) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	if e1 == nil || e2 == nil {
		return false
	}

	return e1.Code == e2.Code && e1.Message == e2.Message
}
