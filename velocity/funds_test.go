package velocity

import (
	"reflect"
	"testing"
)

/**
There are too many tests due to the implementation having too many edge cases. Ideally this should be turned
into a smart fuzz test to cover unknown edge cases
*/
func TestDollar_MarshalJSON(t *testing.T) {
	type fields struct {
		Amount int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		//{name: "testMarshal"},
		//{name: "testUnmarshal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dollar{
				Amount: tt.fields.Amount,
			}
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDollar_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Amount int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dollar{
				Amount: tt.fields.Amount,
			}
			if err := d.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
