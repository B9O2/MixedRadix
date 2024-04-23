package mixedradix

import (
	"fmt"
	"testing"
)

func TestMR(t *testing.T) {
	n := NewMixedRadix(
		//MakeNumberList(0, 9),
		MakeNumberList(0, 9),
		MakeNumberList(0, 9),
	)
	for i := 0; i <= 105; i++ {
		//	fmt.Println(n.Dump())
		fmt.Println(n.Decimal(), n.Format("%s:%s"))
		n.Add(1)
	}

}
