package scm

type IMergeRequest interface {
	SetID(int)
	SetIID(int)
	SetTargetBranch(string)
	SetSourceBranch(string)
	SetAuthors([]string)

	GetID() int
	GetIID() int
	GetTargetBranch() string
	GetSourceBranch() string
	GetAuthors() []string
}

type MergeRequest struct {
	ID           int
	IID          int
	TargetBranch string
	SourceBranch string
	Authors      []string
}

func (mr *MergeRequest) SetID(id int) {
	mr.ID = id
}

func (mr *MergeRequest) SetIID(iid int) {
	mr.IID = iid
}

func (mr *MergeRequest) SetTargetBranch(branch string) {
	mr.TargetBranch = branch
}

func (mr *MergeRequest) SetSourceBranch(branch string) {
	mr.SourceBranch = branch
}

func (mr *MergeRequest) SetAuthors(authors []string) {
	mr.Authors = authors
}

func (mr *MergeRequest) GetID() int {
	return mr.ID
}

func (mr *MergeRequest) GetIID() int {
	return mr.IID
}

func (mr *MergeRequest) GetTargetBranch() string {
	return mr.TargetBranch
}

func (mr *MergeRequest) GetSourceBranch() string {
	return mr.SourceBranch
}

func (mr *MergeRequest) GetAuthors() []string {
	return mr.Authors
}
