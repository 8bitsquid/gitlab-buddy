package gitcmd

import (
	"reflect"
	"testing"
)

func TestNewGitCommand(t *testing.T) {
	tests := []struct {
		name string
		want *GitCommand
	}{
		{"NewGitCommand", NewGitCommand()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitCommand_build(t *testing.T) {
	type args struct {
		command []string
	}
	tests := []struct {
		name string
		fbdc GitCommandBuilder
		args args
		want []string
	}{
		{
			name: "FromBaseDir",
			fbdc: NewGitCommand().FromBaseDir("/some/dir/some/place"),
			args: args{
				command: []string{"some", "git", "command"},
			},
			want: []string{"-C", "/some/dir/some/place", "some", "git", "command"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fbdc.build(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitCommand.build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildArgs(t *testing.T) {
	type args struct {
		args map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArgs(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
