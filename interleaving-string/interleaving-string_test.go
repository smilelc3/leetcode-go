package interleaving_string

import (
	"testing"
)

func TestIsInterleaveExample1(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	ans := isInterleave(s1, s2, s3)
	rightAns := true
	if ans != rightAns {
		t.Error("right ans = ", rightAns, ", current ans = ", ans)
	}
}

func TestIsInterleaveExample2(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	ans := isInterleave(s1, s2, s3)
	rightAns := false
	if ans != rightAns {
		t.Error("right ans = ", rightAns, ", current ans = ", ans)
	}
}

func TestIsInterleaveExample3(t *testing.T) {
	s1 := ""
	s2 := ""
	s3 := ""
	ans := isInterleave(s1, s2, s3)
	rightAns := true
	if ans != rightAns {
		t.Error("right ans = ", rightAns, ", current ans = ", ans)
	}
}

func BenchmarkIisInterleaveDPWithRollArrayMyTest0(b *testing.B) {
	s1 := "aabccaabccaabccaabcc"
	s2 := "dbbcadbbcadbbcadbbca"
	s3 := "aadbbbacccaadbbbacccaadbbbacccaadbbbaccc"
	for idx := 0; idx < b.N; idx++ {
		_ = isInterleaveDPWithRollArray(s1, s2, s3)
	}
}

func BenchmarkIisInterleaveDPWithRollArrayAndStaCompMyTest0(b *testing.B) {
	s1 := "aabccaabccaabccaabcc"
	s2 := "dbbcadbbcadbbcadbbca"
	s3 := "aadbbbacccaadbbbacccaadbbbacccaadbbbaccc"
	for idx := 0; idx < b.N; idx++ {
		_ = isInterleaveDPWithRollArrayAndStaComp(s1, s2, s3)
	}
}
