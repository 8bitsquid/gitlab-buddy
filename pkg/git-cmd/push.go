package gitcmd

const (
	PUSH_CMD            = "push"
	PUSH_UPSTREAM_CMD   = "--set-upstream"
	PUSH_REPOSITORY_CMD = "--repo"
)

type PushCommandBuilder interface {
	Exec() (string, error)
	Repo(string) PushCommandBuilder
	Upstream(string) PushCommandBuilder
	UpdateRemoteRef(string) PushCommandBuilder
}

type PushCommand struct {
	args    map[string]string
	refspec string
	gc      *GitCommand
}

func NewPushCommand(gc *GitCommand) *PushCommand {
	pc := &PushCommand{gc: gc}
	pc.args = make(map[string]string)
	return pc
}

func (pc *PushCommand) Exec() (string, error) {
	commandSlice := pc.Build()
	return pc.gc.Exec(commandSlice)
}

func (pc *PushCommand) Build() []string {
	args := buildArgs(pc.args)
	commandSlice := []string{PUSH_CMD}
	commandSlice = append(commandSlice, args...)
	commandSlice = append(commandSlice, pc.refspec)
	commandSlice = filterCommandSlice(commandSlice)
	return commandSlice
}

func (pc *PushCommand) Repo(repo string) PushCommandBuilder {
	pc.args[PUSH_REPOSITORY_CMD] = repo
	return pc
}

func (pc *PushCommand) Upstream(upstream string) PushCommandBuilder {
	pc.args[PUSH_UPSTREAM_CMD] = upstream
	return pc
}

func (pc *PushCommand) UpdateRemoteRef(ref string) PushCommandBuilder {
	pc.refspec = ref
	return pc
}
