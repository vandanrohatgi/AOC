package utility

func Permutate[T string | int](s []T) [][]T {
	if len(s) == 1 {
		return [][]T{s}
	}

	var result [][]T

	for i := 0; i < len(s); i++ {
		tmp := make([]T, len(s))
		copy(tmp, s)
		tmp = append(tmp[:i], tmp[i+1:]...)
		for _, k := range Permutate(tmp) {
			result = append(result, append([]T{s[i]}, k...))
		}
	}
	return result
}
