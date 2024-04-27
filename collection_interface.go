package collectify

type CollectionInterface[K string | uint | int, T any] interface {
	All() map[K]T

	Average()

	Chunk()
}
