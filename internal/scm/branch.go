package scm

type IBranch interface {
	SetName(string)
	SetDefault(bool)
	SetProtected(bool)

	GetName() string
	IsDefault() bool
	IsProtected() bool
}

type Branch struct {
	Name      string
	Default   bool
	Protected bool
}

func (b *Branch) SetName(name string) {
	b.Name = name
}

func (b *Branch) SetDefault(isDefault bool) {
	b.Default = isDefault
}

func (b *Branch) SetProtected(isProtected bool) {
	b.Protected = isProtected
}

func (b *Branch) GetName() string {
	return b.Name
}

func (b *Branch) IsDefault() bool {
	return b.Default
}

func (b *Branch) IsProtected() bool {
	return b.Protected
}
