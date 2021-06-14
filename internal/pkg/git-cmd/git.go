package gitcmd

import (
	"bytes"
	"os/exec"

	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/tools"
	"go.uber.org/zap"
)

const (
	GIT_COMMAND = "git"
)

type GitCommandBuilder interface {
	build([]string) []string

	Exec([]string) (string, error)
	ExecInDir(string, []string) (string, error)
	FromBaseDir(string) *GitCommand
	Branch() BranchCommandBuilder
	Clone() CloneCommandBuilder
	MergeBase() MergeBaseCommandBuilder
	Push() PushCommandBuilder
	RevParse() RevParseCommandBuilder
	Submodule() SubmoduleCommandBuilder
	Tag() TagCommandBuilder
}

type GitCommand struct {
	baseDir string
}

func NewGitCommand() *GitCommand {
	gc := &GitCommand{}
	return gc
}

func (gc *GitCommand) build(command []string) []string {
	zap.S().Debugw("Building command", "command_slice", command)
	cmd := []string{"-C", gc.baseDir}
	cmd = append(cmd, command...)
	return cmd
}

func (gc *GitCommand) Exec(command []string) (string, error) {
	cmdSlice := gc.build(command)

	zap.S().Debugw("Attempting exec of git", "Command", cmdSlice)
	cmd := exec.Command(GIT_COMMAND, cmdSlice...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (gc *GitCommand) ExecInDir(dir string, command []string) (string, error) {
	inDirCmd := []string{"-C", dir}
	inDirCmd = append(inDirCmd, command...)
	return gc.Exec(inDirCmd)
}

func (gc *GitCommand) FromBaseDir(baseDir string) *GitCommand {
	gc.baseDir = baseDir
	return gc
}

func (gc *GitCommand) Branch() BranchCommandBuilder {
	return NewBranchCommand(gc)
}

func (gc *GitCommand) Clone() CloneCommandBuilder {
	return NewCloneCommand(gc)
}

func (gc *GitCommand) MergeBase() MergeBaseCommandBuilder {
	return NewMergeBaseCommand(gc)
}

func (gc *GitCommand) Push() PushCommandBuilder {
	return NewPushCommand(gc)
}

func (gc *GitCommand) RevParse() RevParseCommandBuilder {
	return NewRevParseCommand(gc)
}

func (gc *GitCommand) Submodule() SubmoduleCommandBuilder {
	return NewSubmoduleCommand(gc)
}

func (gc *GitCommand) Tag() TagCommandBuilder {
	return NewTagCommand(gc)
}

// Helper commands
func buildArgs(args map[string]string) []string {

	argSlice := make([]string, len(args))
	for a, v := range args {
		argSlice = append(argSlice, a)
		if v != "" {
			argSlice = append(argSlice, v)
		}
	}
	return argSlice
}

func filterCommandSlice(cs []string) []string {
	return tools.FilterStringSlice(cs, func(s string) bool {
		return s != ""
	})
}
