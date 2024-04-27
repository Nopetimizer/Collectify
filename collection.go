package collectify

type Collection[K string | uint | int | float32 | float64, T any] struct {
	items map[K]T
}

func (c Collection[K, T]) NewIndexed(with []T) *Collection[uint, T] {
	var (
		items map[uint]T = make(map[uint]T)
		index uint       = 0
		size  uint       = uint(len(with))
	)

	for ; index < size; index++ {
		items[index] = with[index]
	}

	return &Collection[uint, T]{
		items: items,
	}
}

func (c Collection[string, T]) NewKeyed(items map[string]T) *Collection[string, T] {
	return &Collection[string, T]{
		items: items,
	}
}

func (c *Collection[K, T]) All() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Average() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Chunk() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) ChunkWhile() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Collapse() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Combine() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Concat() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Contains() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) ContainsOneItem() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Count() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) CountBy() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) CrossJoin() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Diff() *Collection[K, T] {
	return c
}

func (c *Collection[K, T]) Filter(callback func(item T, key K) bool) *Collection[K, T] {
	var (
		items map[K]T = make(map[K]T)
	)

	for k, v := range c.items {
		if callback(v, k) {
			items[k] = v
		}
	}

	return (Collection[K, T]{}).NewKeyed(items)
}

func (c *Collection[K, T]) Each(callback func(item T, key K) bool) *Collection[K, T] {
	for k, v := range c.items {
		if callback(v, k) {
			break
		}
	}

	return c
}
