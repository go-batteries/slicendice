package slicendice

type IntTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Take[E any, T IntTypes](elements []E, takeN int) []E {
	n := len(elements)

	if n < 0 {
		return []E{}
	}

	if n < takeN {
		return elements
	}

	return elements[:takeN]
}
