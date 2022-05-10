package gitcmd

import (
	"reflect"
	"testing"
)

func TestPushCommand_Exec(t *testing.T) {
	tests := []struct {
		name    string
		pc      *PushCommand
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pc.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("PushCommand.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PushCommand.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		pc   *PushCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushCommand_Repo(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name string
		pc   *PushCommand
		args args
		want PushCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.Repo(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushCommand.Repo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushCommand_Upstream(t *testing.T) {
	type args struct {
		upstream string
	}
	tests := []struct {
		name string
		pc   *PushCommand
		args args
		want PushCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.Upstream(tt.args.upstream); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushCommand.Upstream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushCommand_UpdateRemoteRef(t *testing.T) {
	type args struct {
		ref string
	}
	tests := []struct {
		name string
		pc   *PushCommand
		args args
		want PushCommandBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.UpdateRemoteRef(tt.args.ref); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushCommand.UpdateRemoteRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
