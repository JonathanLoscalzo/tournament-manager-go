package interfaces

type Repository[T any] interface {
	Create(dto T) (T, error)
	Get(dto T) (T, error)
	Delete(id string) error
	Update(id string, dto T) (T, error)
}
