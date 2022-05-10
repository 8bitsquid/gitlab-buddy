package gitcmd

const (
	CLONE_CMD        = "clone"
	CLONE_BRANCH_CMD = "--branch"
	CLONE_ORIGIN_CMD = "--origin"
)

type CloneCommandBuilder interface {
	Exec() (string, error)
	Build() []string
	Branch(string) CloneCommandBuilder
	Origin(string) CloneCommandBuilder
	Repo(string) CloneCommandBuilder
}

type CloneCommand struct {
	args map[string]string
	repo string
	gc   *GitCommand
}

func NewCloneCommand(gc *GitCommand) *CloneCommand {
	cc := &CloneCommand{gc: gc}
	cc.args = make(map[string]string)
	return cc
}

func (cc *CloneCommand) Exec() (string, error) {
	command := cc.Build()
	return cc.gc.Exec(command)
}

func (cc *CloneCommand) Build() []string {
	args := buildArgs(cc.args)
	commandSlice := []string{CLONE_CMD}
	commandSlice = append(commandSlice, args...)
	commandSlice = append(commandSlice, cc.repo)
	commandSlice = filterCommandSlice(commandSlice)
	return commandSlice
}

func (cc *CloneCommand) Branch(branch string) CloneCommandBuilder {
	cc.args[CLONE_BRANCH_CMD] = branch
	return cc
}

func (cc *CloneCommand) Origin(origin string) CloneCommandBuilder {
	cc.args[CLONE_ORIGIN_CMD] = origin
	return cc
}

func (cc *CloneCommand) Repo(repo string) CloneCommandBuilder {
	cc.repo = repo
	return cc
}
