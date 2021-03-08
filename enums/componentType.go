package enums

type ComponentType int

const (
	SpriteRenderer     ComponentType = iota
	SpriteAnimator     ComponentType = iota
	BoxCollider        ComponentType = iota
	LineCollider       ComponentType = iota
	KeyboardController ComponentType = iota
)

func (c ComponentType) String() string {
	return [...]string{"SpriteRenderer", "SpriteAnimator", "BoxCollider", "LineCollider", "KeyboardController"}[c]
}
