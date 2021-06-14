package gitcmd

import (
	"reflect"
	"testing"
)

// TODO: figure out exec comand testing
// func TestCloneCommand_Exec(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		cc      *CloneCommand
// 		want    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.cc.Exec()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CloneCommand.Exec() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("CloneCommand.Exec() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestCloneCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		cc   CloneCommandBuilder
		want []string
	}{
		{"Repo", NewGitCommand().Clone().Repo("repo/clone"), []string{"clone", "repo/clone"}},
		{"Repo.Branch", NewGitCommand().Clone().Repo("repo/clone").Branch("branch_clone"), []string{"clone", "--branch", "branch_clone", "repo/clone"}},
		{"Origin.Repo", NewGitCommand().Clone().Origin("different/origin").Repo("repo/clone"), []string{"clone", "--origin", "different/origin", "repo/clone"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cc.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
