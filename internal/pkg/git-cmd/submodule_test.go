package gitcmd

import (
	"reflect"
	"testing"
)

func TestSubmoduleCommand_Exec(t *testing.T) {
	tests := []struct {
		name    string
		sc      *SubmoduleCommand
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sc.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("SubmoduleCommand.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubmoduleCommand.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmoduleCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		sc   SubmoduleCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubmoduleCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmoduleCommand_Status(t *testing.T) {
	tests := []struct {
		name string
		sc   *SubmoduleCommand
		want SubmoduleCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.Status(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubmoduleCommand.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmoduleCommand_Foreach(t *testing.T) {
	type args struct {
		subcmd string
	}
	tests := []struct {
		name string
		sc   *SubmoduleCommand
		args args
		want SubmoduleCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.Foreach(tt.args.subcmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubmoduleCommand.Foreach() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmoduleCommand_SetBranch(t *testing.T) {
	type args struct {
		branch string
	}
	tests := []struct {
		name string
		sc   *SubmoduleCommand
		args args
		want SubmoduleCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.SetBranch(tt.args.branch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubmoduleCommand.SetBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmoduleCommand_SetBranchToDefault(t *testing.T) {
	tests := []struct {
		name string
		sc   *SubmoduleCommand
		want SubmoduleCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.SetBranchToDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubmoduleCommand.SetBranchToDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
