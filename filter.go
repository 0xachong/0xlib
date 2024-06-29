package xlib

func Filter[T any](data []T, fn func(t T) bool) []T {
	var res []T
	for _, d := range data {
		if fn(d) {
			res = append(res, d)
		}
	}
	return res
}

func Map[T any, U any](f func(T) U, s []T) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
