package scm

type ISubmodule interface {
	SetSHA1(string)
	SetPath(string)
	SetBranch(string)

	GetSHA1() string
	GetPath() string
	GetBranch() string
}

type Submodule struct {
	SHA1 string
	Path string
	Branch string
}

func (sm *Submodule) SetSHA1(sha1 string) {
	sm.SHA1 = sha1
}

func (sm *Submodule) SetPath(path string) {
	sm.Path = path
}

func (sm *Submodule) SetBranch(msg string) {
	sm.Branch = msg
}

func (sm *Submodule) GetSHA1() string {
	return sm.SHA1
}

func (sm *Submodule) GetPath() string {
	return sm.Path
}

func (sm *Submodule) GetBranch() string {
	return sm.Branch
}
