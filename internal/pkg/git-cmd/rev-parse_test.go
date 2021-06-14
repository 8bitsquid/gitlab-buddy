package gitcmd

import (
	"reflect"
	"testing"
)

func TestRevParseCommand_Exec(t *testing.T) {
	tests := []struct {
		name    string
		rp      *RevParseCommand
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rp.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("RevParseCommand.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RevParseCommand.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevParseCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		rp   RevParseCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rp.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevParseCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevParseCommand_Branch(t *testing.T) {
	type args struct {
		branch string
	}
	tests := []struct {
		name string
		rp   *RevParseCommand
		args args
		want RevParseCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rp.Branch(tt.args.branch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevParseCommand.Branch() = %v, want %v", got, tt.want)
			}
		})
	}
}
