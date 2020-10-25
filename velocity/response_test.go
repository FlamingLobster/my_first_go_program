package velocity

import (
	"reflect"
	"testing"
)

func TestAccepted(t *testing.T) {
	type args struct {
		loadFund *Funds
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		//{name: "testAccept"},
		//{name: "testReject"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Accepted(tt.args.loadFund); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Accepted() = %v, want %v", got, tt.want)
			}
		})
	}
}
