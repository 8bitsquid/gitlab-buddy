package scm

type IGroupService interface {
	Get(interface{}) (IGroup, error)
	Create(IGroup) (IGroup, error)
	CloneRepo(IGroup, IRepository) (IRepository, error)
	GetAllRepos(IGroup) []IRepository
}

type IGroup interface {
	SetID(int)
	SetName(string)

	GetID() int
	GetName() string
}

type Group struct {
	ID   int
	Name string
}

func (g *Group) SetID(id int) {
	g.ID = id
}

func (g *Group) SetName(name string) {
	g.Name = name
}

func (g *Group) GetID() int {
	return g.ID
}

func (g *Group) GetName() string {
	return g.Name
}
