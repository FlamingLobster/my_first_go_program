package velocity

import (
	"reflect"
	"testing"
	"time"
)

func TestToStartOfDay(t *testing.T) {
	type args struct {
		unrounded time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		//{name: "testRandomDate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStartOfDay(tt.args.unrounded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStartOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStartOfWeek(t *testing.T) {
	type args struct {
		unrounded time.Time
	}
	tests := []struct {
		name string
		args args
		want UniqueTransactionKey
	}{
		//{name: "testRandomWeek"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStartOfWeek(tt.args.unrounded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStartOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
