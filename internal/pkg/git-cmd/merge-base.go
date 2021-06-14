package gitcmd

const (
	MERGE_BASE_CMD = "merge-base"
)

type MergeBaseCommandBuilder interface {
	Exec() (string, error)
	Build() []string
	CheckBase(string) MergeBaseCommandBuilder
	AgainstBase(string) MergeBaseCommandBuilder
}

type MergeBaseCommand struct {
	args    map[string]string
	commitA string
	commitB string
	gc      *GitCommand
}

func NewMergeBaseCommand(gc *GitCommand) *MergeBaseCommand {
	mbc := &MergeBaseCommand{gc: gc}
	return mbc
}

func (mb *MergeBaseCommand) Exec() (string, error) {
	commandSlice := mb.Build()
	return mb.gc.Exec(commandSlice)

}

func (mb *MergeBaseCommand) Build() []string {
	commandSlice := []string{MERGE_BASE_CMD, mb.commitA, mb.commitB}
	commandSlice = filterCommandSlice(commandSlice)
	return commandSlice
}

func (mb *MergeBaseCommand) CheckBase(commit string) MergeBaseCommandBuilder {
	mb.commitA = commit
	return mb
}

func (mb *MergeBaseCommand) AgainstBase(commit string) MergeBaseCommandBuilder {
	mb.commitB = commit
	return mb
}
