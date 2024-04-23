package mixedradix

import "strconv"

func MakeNumberList(from, to uint) []string {
	if to < from {
		from, to = to, from
	}
	var result []string

	for i := from; i <= to; i++ {
		result = append(result, strconv.Itoa(int(i)))
	}
	return result
}
