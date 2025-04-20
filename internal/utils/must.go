package utils

func must[T any, O any](fn func(...O) (T, error), opts ...O) T {
	res, err := fn(opts...)
	if err != nil {
		panic(err)
	}
	return res
}