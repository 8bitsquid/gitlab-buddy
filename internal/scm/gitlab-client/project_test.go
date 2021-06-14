package gitlabclient

import (
	"reflect"
	"testing"

	"github.com/xanzy/go-gitlab"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/scm"
)

func TestNewProject(t *testing.T) {
	type args struct {
		proj *gitlab.Project
	}
	tests := []struct {
		name string
		args args
		want scm.IRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProject(tt.args.proj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_Get(t *testing.T) {
	type args struct {
		projectID interface{}
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.Get(tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_Clone(t *testing.T) {
	type args struct {
		repo scm.IRepository
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.Clone(tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.Clone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_Push(t *testing.T) {
	type args struct {
		repo scm.IRepository
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ps.Push(tt.args.repo); (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectService_AddTag(t *testing.T) {
	type args struct {
		repo    scm.IRepository
		tagName string
		branch  string
		message string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.ITag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.AddTag(tt.args.repo, tt.args.tagName, tt.args.branch, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.AddTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.AddTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_ProtectTag(t *testing.T) {
	type args struct {
		repo    scm.IRepository
		tagName string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.ITag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.ProtectTag(tt.args.repo, tt.args.tagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.ProtectTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.ProtectTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_MoveBranch(t *testing.T) {
	type args struct {
		repo      scm.IRepository
		oldBranch string
		newBranch string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IBranch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.MoveBranch(tt.args.repo, tt.args.oldBranch, tt.args.newBranch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.MoveBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.MoveBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_SetDefaultBranch(t *testing.T) {
	type args struct {
		repo   scm.IRepository
		branch string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IBranch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.SetDefaultBranch(tt.args.repo, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.SetDefaultBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.SetDefaultBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_DeleteBranch(t *testing.T) {
	type args struct {
		repo   scm.IRepository
		branch string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.DeleteBranch(tt.args.repo, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.DeleteBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.DeleteBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_ProtectBranch(t *testing.T) {
	type args struct {
		repo   scm.IRepository
		branch string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IBranch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.ProtectBranch(tt.args.repo, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.ProtectBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.ProtectBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_UnprotectBranch(t *testing.T) {
	type args struct {
		repo   scm.IRepository
		branch string
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IBranch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.UnprotectBranch(tt.args.repo, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.UnprotectBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.UnprotectBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_createProject(t *testing.T) {
	type args struct {
		repo scm.IRepository
	}
	tests := []struct {
		name    string
		ps      *ProjectService
		args    args
		want    scm.IRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.createProject(tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.createProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.createProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectService_getCreatePayload(t *testing.T) {
	type args struct {
		p *Project
	}
	tests := []struct {
		name string
		ps   *ProjectService
		args args
		want *gitlab.CreateProjectOptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.getCreatePayload(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectService.getCreatePayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
