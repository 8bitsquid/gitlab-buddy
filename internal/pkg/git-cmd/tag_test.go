package gitcmd

import (
	"reflect"
	"testing"
)

func TestTagCommand_Exec(t *testing.T) {
	tests := []struct {
		name    string
		tc      *TagCommand
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tc.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("TagCommand.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TagCommand.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		tc   *TagCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCommand_CreateTag(t *testing.T) {
	type args struct {
		tagName string
	}
	tests := []struct {
		name string
		tc   *TagCommand
		args args
		want TagCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.CreateTag(tt.args.tagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCommand.CreateTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCommand_GetTag(t *testing.T) {
	type args struct {
		tagName string
	}
	tests := []struct {
		name string
		tc   *TagCommand
		args args
		want TagCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.GetTag(tt.args.tagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCommand.GetTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCommand_WithCommit(t *testing.T) {
	type args struct {
		commit string
	}
	tests := []struct {
		name string
		tc   *TagCommand
		args args
		want TagCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.WithCommit(tt.args.commit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCommand.WithCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCommand_WithBranch(t *testing.T) {
	type args struct {
		branch string
	}
	tests := []struct {
		name string
		tc   *TagCommand
		args args
		want TagCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.WithBranch(tt.args.branch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCommand.WithBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCommand_WithMessage(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		tc   *TagCommand
		args args
		want TagCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tc.WithMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCommand.WithMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
