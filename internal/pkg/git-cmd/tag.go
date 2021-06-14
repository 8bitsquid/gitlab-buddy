package gitcmd

const (
	TAG_CMD         = "tag"
	TAG_MESSAGE_CMD = "-m"
	TAG_LIST_CMD    = "-l"
)

type TagCommandBuilder interface {
	Exec() (string, error)
	Build() []string
	CreateTag(string) TagCommandBuilder
	GetTag(string) TagCommandBuilder
	WithBranch(string) TagCommandBuilder
	WithCommit(string) TagCommandBuilder
	WithMessage(string) TagCommandBuilder
}

type TagCommand struct {
	gc      *GitCommand
	args    map[string]string
	tagName string
	commit  string
}

func NewTagCommand(gc *GitCommand) *TagCommand {
	bc := &TagCommand{gc: gc}
	bc.args = make(map[string]string)
	return bc
}

func (tc *TagCommand) Exec() (string, error) {
	command := tc.Build()
	return tc.gc.Exec(command)
}

func (tc *TagCommand) Build() []string {
	args := buildArgs(tc.args)
	commandSlice := []string{TAG_CMD}

	if _, isListCommand := tc.args[TAG_LIST_CMD]; isListCommand {
		return append(commandSlice, args...)
	}

	commandSlice = append(commandSlice, args...)
	commandSlice = append(commandSlice, tc.tagName)
	commandSlice = append(commandSlice, tc.commit)
	commandSlice = filterCommandSlice(commandSlice)
	return commandSlice
}

func (tc *TagCommand) CreateTag(tagName string) TagCommandBuilder {
	tc.tagName = tagName
	return tc
}

func (tc *TagCommand) GetTag(tagName string) TagCommandBuilder {
	tc.args[TAG_LIST_CMD] = tagName
	return tc
}

func (tc *TagCommand) WithCommit(commit string) TagCommandBuilder {
	tc.commit = commit
	return tc
}

func (tc *TagCommand) WithBranch(branch string) TagCommandBuilder {
	return tc.WithCommit(branch)
}

func (tc *TagCommand) WithMessage(msg string) TagCommandBuilder {
	tc.args[TAG_MESSAGE_CMD] = msg
	return tc
}
