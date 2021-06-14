package scm

import (
	"time"
)

const (
	CONCURRENCY_LIMIT = 10
	TIMEOUT           = 20 * time.Second

	PAGINATOIN_PER_PAGE          = 100
	PAGINATOIN_CONCURRENCY_LIMIT = 3
)

type IClient interface {
	GroupService() IGroupService
	RepoService() IRepoService
	Cleanup() error
}
