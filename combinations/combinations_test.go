package combinations

import (
	"reflect"
	"testing"
)

func TestCombineExample1(t *testing.T) {
	n := 4
	k := 2
	ans := combine(n, k)
	rightAns := [][]int{
		{2, 4},
		{3, 4},
		{2, 3},
		{1, 2},
		{1, 3},
		{1, 4}}
	if !reflect.DeepEqual(ans, rightAns) {
		t.Error("right ans = ", rightAns, ", current ans = ", ans)
	}
}

func TestCombineExample2(t *testing.T) {
	n := 1
	k := 1
	ans := combine(n, k)
	rightAns := [][]int{
		{1}}
	if !reflect.DeepEqual(ans, rightAns) {
		t.Error("right ans = ", rightAns, ", current ans = ", ans)
	}
}

func BenchmarkCombineMyTest0(b *testing.B) {
	n := 20
	k := 10
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		combine(n, k)
	}
}
