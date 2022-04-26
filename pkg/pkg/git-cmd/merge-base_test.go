package gitcmd

import (
	"reflect"
	"testing"
)

func TestMergeBaseCommand_Build(t *testing.T) {
	tests := []struct {
		name string
		mb   MergeBaseCommandBuilder
		want []string
	}{
		{"CheckBase.AgainstBase", NewGitCommand().MergeBase().CheckBase("some_commit").AgainstBase("some_other_commit"), []string{"merge-base", "some_commit", "some_other_commit"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mb.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeBaseCommand.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
