package pieces

const (
	P = iota
	R
	N
	B
	Q
	K
	p = iota + 6
	r
	n
	b
	q
	k
)

var Pieces = map[rune]int{
	'P': P,
	'R': R,
	'N': N,
	'B': B,
	'K': K,
	'p': p,
	'r': r,
	'n': n,
	'b': b,
	'q': q,
	'k': k,
}
