package scm

type ITag interface {
	SetCommit(string)
	SetName(string)
	SetMessage(string)
	SetDescription(string)
	SetProtected(bool)

	GetCommit() string
	GetName() string
	GetMessage() string
	GetDescription() string
	GetProtected() bool
}

type Tag struct {
	Commit      string
	Name        string
	Message     string
	Description string
	Protected   bool
}

func (t *Tag) SetCommit(commit string) {
	t.Commit = commit
}

func (t *Tag) SetName(name string) {
	t.Name = name
}

func (t *Tag) SetMessage(msg string) {
	t.Message = msg
}

func (t *Tag) SetDescription(description string) {
	t.Description = description
}

func (t *Tag) SetProtected(protected bool) {
	t.Protected = protected
}

func (t *Tag) GetCommit() string {
	return t.Commit
}

func (t *Tag) GetName() string {
	return t.Name
}

func (t *Tag) GetMessage() string {
	return t.Message
}

func (t *Tag) GetDescription() string {
	return t.Description
}

func (t *Tag) GetProtected() bool {
	return t.Protected
}
