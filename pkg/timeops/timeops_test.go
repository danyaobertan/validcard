package timeops

import (
	"testing"
	"time"
)

func TestIsFutureDate(t *testing.T) {
	currentYear, currentMonth, _ := time.Now().Date()
	nextMonth := time.Now().AddDate(0, 1, 0).Month()
	prevMonth := time.Now().AddDate(0, -1, 0).Month()
	nextYear := time.Now().AddDate(1, 0, 0).Year()
	prevYear := time.Now().AddDate(-1, 0, 0).Year()

	type args struct {
		month int
		year  int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "current month and year",
			args: args{month: int(currentMonth), year: currentYear},
			want: true,
		},
		{
			name: "next month",
			args: args{month: int(nextMonth), year: currentYear},
			want: true,
		},
		{
			name: "previous month",
			args: args{month: int(prevMonth), year: currentYear},
			want: false,
		},
		{
			name: "next year",
			args: args{month: int(currentMonth), year: nextYear},
			want: true,
		},
		{
			name: "previous year",
			args: args{month: int(currentMonth), year: prevYear},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFutureDate(tt.args.month, tt.args.year); got != tt.want {
				t.Errorf("IsFutureDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsYearInRange(t *testing.T) {
	type args struct {
		year          int
		maxValidYears int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "current year",
			args: args{year: time.Now().Year(), maxValidYears: 20},
			want: true,
		},
		{
			name: "next year",
			args: args{year: time.Now().Year() + 1, maxValidYears: 20},
			want: true,
		},
		{
			name: "20 years in the future",
			args: args{year: time.Now().Year() + 20, maxValidYears: 20},
			want: true,
		},
		{
			name: "21 years in the future",
			args: args{year: time.Now().Year() + 21, maxValidYears: 20},
			want: false,
		},
		{
			name: "previous year",
			args: args{year: time.Now().Year() - 1, maxValidYears: 20},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsYearInRange(tt.args.year, tt.args.maxValidYears); got != tt.want {
				t.Errorf("IsYearInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
