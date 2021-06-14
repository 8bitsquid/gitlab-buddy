package gitcmd

const (
	SUBMODULE_CMD                        = "submodule"
	SUBMODULE_STATUS_CMD                 = "status"
	SUBMODULE_FOREACH_CMD                = "foreach"
	SUBMODULE_SET_BRANCH_CMD             = "set-branch"
	SUBMODULE_SET_BRANCH_CMD_BRANCH_OPT  = "--branch"
	SUBMODULE_SET_BRANCH_CMD_DEFAULT_OPT = "--default"
)

type SubmoduleCommandBuilder interface {
	Exec() (string, error)
	Status() SubmoduleCommandBuilder
	Foreach(string) SubmoduleCommandBuilder
	SetBranch(string) SubmoduleCommandBuilder
	SetBranchToDefault() SubmoduleCommandBuilder
}

type SubmoduleCommand struct {
	subCommand string
	args       map[string]string
	gc         *GitCommand
}

func NewSubmoduleCommand(gc *GitCommand) *SubmoduleCommand {
	smc := &SubmoduleCommand{gc: gc}
	smc.args = make(map[string]string)
	return smc
}

func (sc *SubmoduleCommand) Exec() (string, error) {
	command := sc.Build()
	return sc.gc.Exec(command)
}

func (sc SubmoduleCommand) Build() []string {
	args := buildArgs(sc.args)
	commandSlice := []string{SUBMODULE_CMD}
	commandSlice = append(commandSlice, sc.subCommand)
	commandSlice = append(commandSlice, args...)
	commandSlice = filterCommandSlice(commandSlice)
	return commandSlice
}

func (sc *SubmoduleCommand) Status() SubmoduleCommandBuilder {
	sc.subCommand = SUBMODULE_STATUS_CMD
	return sc
}

func (sc *SubmoduleCommand) Foreach(subcmd string) SubmoduleCommandBuilder {
	// clear command, as no command optinos are supported
	// Currently does not support `--recursive` foreach option
	sc.subCommand = ""
	sc.args[SUBMODULE_FOREACH_CMD] = subcmd
	return sc
}

func (sc *SubmoduleCommand) SetBranch(branch string) SubmoduleCommandBuilder {
	sc.subCommand = SUBMODULE_SET_BRANCH_CMD
	sc.args[SUBMODULE_SET_BRANCH_CMD_BRANCH_OPT] = branch
	return sc
}

func (sc *SubmoduleCommand) SetBranchToDefault() SubmoduleCommandBuilder {
	sc.subCommand = SUBMODULE_SET_BRANCH_CMD
	sc.args[SUBMODULE_SET_BRANCH_CMD_DEFAULT_OPT] = ""
	return sc
}