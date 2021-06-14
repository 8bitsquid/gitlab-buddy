package tools

import (
	"reflect"
	"testing"
)

func TestFilterStringSlice(t *testing.T) {
	type args struct {
		slice []string
		f     func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Filter empty strings", 
			args{
				[]string{"some", "of", "", "strings", "are", "", "", "empty"},
				func(s string) bool {
					return s != ""
				},
			}, []string{"some", "of", "strings", "are", "empty"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterStringSlice(tt.args.slice, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
