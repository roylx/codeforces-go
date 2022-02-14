// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	examples := [][]string{
		{
			`1`, `2`, `1`, `600`, 
			`6`,
		},
		{
			`0`, `1`, `2`, `76`, 
			`6`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minCostSetTime, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-71/problems/minimum-cost-to-set-cooking-time/
