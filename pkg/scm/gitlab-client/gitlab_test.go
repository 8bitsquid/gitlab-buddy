package gitlabclient

import (
	"reflect"
	"testing"

	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
)

func TestNewClient(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    scm.IClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitlabClient_GroupService(t *testing.T) {
	tests := []struct {
		name string
		g    GitlabClient
		want scm.IGroupService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.GroupService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitlabClient.GroupService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitlabClient_RepoService(t *testing.T) {
	tests := []struct {
		name string
		g    GitlabClient
		want scm.IRepoService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RepoService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitlabClient.RepoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitlabClient_Cleanup(t *testing.T) {
	tests := []struct {
		name    string
		g       *GitlabClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.Cleanup(); (err != nil) != tt.wantErr {
				t.Errorf("GitlabClient.Cleanup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
