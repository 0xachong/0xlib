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
