package gitcmd

import (
	"reflect"
	"testing"
)

// TODO: figure out exec command testing
// func TestBranchCommand_Exec(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		bc      *BranchCommand
// 		want    string
// 		wantErr bool
// 	}{
// 		{"Branch List", NewGitCommand().Branch.List().Build()}
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.bc.Exec()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("BranchCommand.Exec() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("BranchCommand.Exec() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestBranchCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		bc   BranchCommandBuilder
		want []string
	}{
		{"List", NewGitCommand().Branch().List(), []string{"branch", "--list"}},
		{"Move.To", NewGitCommand().Branch().Move("old_branch").To("new_branch"), []string{"branch", "--move", "old_branch", "new_branch"}},
		{"Delete", NewGitCommand().Branch().Delete("branch_name"), []string{"branch", "--delete", "branch_name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BranchCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
