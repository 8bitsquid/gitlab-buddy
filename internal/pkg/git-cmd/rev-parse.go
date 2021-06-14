package gitcmd

const (
	REVPARSE_CMD = "rev-parse"
)

type RevParseCommandBuilder interface {
	Exec() (string, error)
	Branch(string) RevParseCommandBuilder
}

type RevParseCommand struct {
	branch string
	gc     *GitCommand
}

func NewRevParseCommand(gc *GitCommand) *RevParseCommand {
	rpc := &RevParseCommand{gc: gc}
	return rpc
}

func (rp *RevParseCommand) Exec() (string, error) {
	command := rp.Build()
	return rp.gc.Exec(command)
}

func (rp RevParseCommand) Build() []string {
	return []string{REVPARSE_CMD, rp.branch}
}

func (rp *RevParseCommand) Branch(branch string) RevParseCommandBuilder {
	rp.branch = branch
	return rp
}
