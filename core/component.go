package core

type Component interface {
	OnUpdate() error
	OnDraw() error
}

type ComponentAttributes interface {
	IsUpdatable() bool
	IsDrawable() bool
}

type componentAttributes struct {
	updatable bool
	drawable  bool
}

func (c componentAttributes) IsUpdatable() bool {
	return c.updatable
}

func (c componentAttributes) IsDrawable() bool {
	return c.drawable
}

func NewComponentAttributes(updatable bool, drawable bool) ComponentAttributes {
	return &componentAttributes{
		updatable:     updatable,
		drawable:      drawable,
	}
}