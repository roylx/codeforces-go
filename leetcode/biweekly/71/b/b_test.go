// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	examples := [][]string{
		{
			`[9,12,5,10,14,3,10]`, `10`, 
			`[9,5,3,10,10,12,14]`,
		},
		{
			`[-3,4,3,2]`, `2`, 
			`[-3,2,4,3]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, pivotArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-71/problems/partition-array-according-to-given-pivot/
