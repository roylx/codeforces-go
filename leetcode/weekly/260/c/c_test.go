// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[["#", " ", "#"], [" ", " ", "#"], ["#", "c", " "]]`, `"abc"`, 
			`true`,
		},
		{
			`[[" ", "#", "a"], [" ", "#", "c"], [" ", "#", "a"]]`, `"ac"`, 
			`false`,
		},
		{
			`[["#", " ", "#"], [" ", " ", "#"], ["#", " ", "c"]]`, `"ca"`, 
			`true`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, placeWordInCrossword, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-260/problems/check-if-word-can-be-placed-in-crossword/
