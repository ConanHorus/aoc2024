package library

type Lazy[T any] struct {
	loaded bool
	value  T
	loader func() T
}

func NewLazy[T any](loader func() T) *Lazy[T] {
	return &Lazy[T]{
		loader: loader,
		loaded: false,
	}
}

func (this *Lazy[T]) Value() T {
	if !this.loaded {
		this.value = this.loader()
		this.loaded = true
	}

	return this.value
}
