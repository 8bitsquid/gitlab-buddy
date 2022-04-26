package gitlabclient

import (
	"reflect"
	"testing"

	"github.com/xanzy/go-gitlab"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
)

func TestGroupService_Get(t *testing.T) {
	type args struct {
		groupID interface{}
	}
	tests := []struct {
		name    string
		gs      *GroupService
		args    args
		want    scm.IGroup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.gs.Get(tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupService_Create(t *testing.T) {
	type args struct {
		group scm.IGroup
	}
	tests := []struct {
		name    string
		gs      *GroupService
		args    args
		want    scm.IGroup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.gs.Create(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupService_CloneRepo(t *testing.T) {
	type args struct {
		group scm.IGroup
		repo  scm.IRepository
	}
	tests := []struct {
		name    string
		gs      *GroupService
		args    args
		want    scm.IRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.gs.CloneRepo(tt.args.group, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.CloneRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupService.CloneRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupService_GetAllRepos(t *testing.T) {
	type args struct {
		group scm.IGroup
	}
	tests := []struct {
		name string
		gs   *GroupService
		args args
		want []scm.IRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.GetAllRepos(tt.args.group); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupService.GetAllRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupService_getRepoListPage(t *testing.T) {
	type args struct {
		group interface{}
		page  int
	}
	tests := []struct {
		name    string
		gs      *GroupService
		args    args
		want    []*gitlab.Project
		want1   *gitlab.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.gs.getRepoListPage(tt.args.group, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.getRepoListPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupService.getRepoListPage() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GroupService.getRepoListPage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
