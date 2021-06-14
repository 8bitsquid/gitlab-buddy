package tools

// import (
// 	"net/url"
// 	"reflect"
// 	"testing"
// )

// func TestNewURL(t *testing.T) {
// 	type args struct {
// 		u string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *url.URL
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewURL(tt.args.u)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewURL() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewURL() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUpdateHostName(t *testing.T) {
// 	type args struct {
// 		from url.URL
// 		to   string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *url.URL
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := UpdateHostName(tt.args.from, tt.args.to)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UpdateHostName() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UpdateHostName() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
