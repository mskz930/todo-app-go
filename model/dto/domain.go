package dto

type Response[T interface{}] struct {
	Status  int
	Message string
	Data    T
}
