package gitcmd

const (
	BRANCH_CMD        = "branch"
	BRANCH_LIST_CMD   = "--list"
	BRANCH_MOVE_CMD   = "--move"
	BRANCH_DELETE_CMD = "--delete"
)

type BranchCommandBuilder interface {
	Exec() (string, error)
	Build() []string
	List() BranchCommandBuilder
	Move(string) BranchCommandBuilder
	To(string) BranchCommandBuilder
	Delete(string) BranchCommandBuilder
}

type BranchCommand struct {
	gc         *GitCommand
	SubCommand string
	args       map[string]string
	branchName string
}

func NewBranchCommand(gc *GitCommand) *BranchCommand {
	bc := &BranchCommand{gc: gc}
	bc.args = make(map[string]string)
	return bc
}

func (bc *BranchCommand) Exec() (string, error) {
	command := bc.Build()
	return bc.gc.Exec(command)
}

func (bc *BranchCommand) Build() []string {
	args := buildArgs(bc.args)
	commandSlice := []string{BRANCH_CMD}
	commandSlice = append(commandSlice, bc.SubCommand)
	commandSlice = append(commandSlice, args...)
	commandSlice = filterCommandSlice(commandSlice)
	return commandSlice
}

func (bc *BranchCommand) List() BranchCommandBuilder {
	bc.SubCommand = BRANCH_LIST_CMD
	return bc
}

func (bc *BranchCommand) Move(branch string) BranchCommandBuilder {
	bc.SubCommand = BRANCH_MOVE_CMD
	bc.branchName = branch
	return bc
}

func (bc *BranchCommand) To(branch string) BranchCommandBuilder {
	bc.args[bc.branchName] = branch
	return bc
}

func (bc *BranchCommand) Delete(branch string) BranchCommandBuilder {
	bc.args[BRANCH_DELETE_CMD] = branch
	return bc
}
