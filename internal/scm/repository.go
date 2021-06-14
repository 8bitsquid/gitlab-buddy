package scm

type IRepoService interface {
	Get(interface{}) (IRepository, error)
	Clone(IRepository) (IRepository, error)
	Push(IRepository) error

	AddTag(IRepository, string, string, string) (ITag, error)
	ProtectTag(IRepository, string) (ITag, error)

	// TODO: Move these methods to a BrachService interface in scm/branch.go
	GetBranch(IRepository, string) (IBranch, error)
	MoveBranch(IRepository, string, string) (IBranch, error)
	SetDefaultBranch(IRepository, string) (IBranch, error)
	DeleteBranch(IRepository, string) (IResponse, error)
	ProtectBranch(IRepository, string) (IBranch, error)
	UnprotectBranch(IRepository, string) (IBranch, error)
	UpdateMergeRequestsToNewBranch(IRepository, string, string) error

	HasSubmodules(IRepository) bool
}

type IRepository interface {
	SetID(int)
	SetName(string)
	SetPath(string)
	SetURL(string)
	SetBranch(string)
	SetGroupID(int)
	SetCloneURL(string)
	SetUpstream(string)

	GetID() int
	GetName() string
	GetPath() string
	GetURL() string
	GetBranch() string
	GetGroupID() int
	GetCloneURL() string
	GetUpstream() string
}

type Repository struct {
	ID       int
	Name     string
	Path     string
	URL      string
	Branch   string
	GroupID  int
	CloneURL string
	Upstream string
}

func (r *Repository) SetID(id int) {
	r.ID = id
}

func (r *Repository) SetName(name string) {
	r.Name = name
}

func (r *Repository) SetPath(path string) {
	r.Path = path
}

func (r *Repository) SetURL(url string) {
	r.URL = url
}

func (r *Repository) SetBranch(branch string) {
	r.Branch = branch
}

func (r *Repository) SetGroupID(groupID int) {
	r.GroupID = groupID
}

func (r *Repository) SetCloneURL(cloneURL string) {
	r.CloneURL = cloneURL
}

func (r *Repository) SetUpstream(upstream string) {
	r.Upstream = upstream
}

func (r *Repository) GetID() int {
	return r.ID
}

func (r *Repository) GetName() string {
	return r.Name
}

func (r *Repository) GetPath() string {
	return r.Path
}

func (r *Repository) GetURL() string {
	return r.URL
}

func (r *Repository) GetBranch() string {
	return r.Branch
}

func (r *Repository) GetGroupID() int {
	return r.GroupID
}

func (r *Repository) GetCloneURL() string {
	return r.CloneURL
}

func (r *Repository) GetUpstream() string {
	return r.Upstream
}
