// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`2`, `5`, 
			`25`,
		},
		{
			`3`, `7`, 
			`499`,
		},
		{
			`7`, `17`, 
			`20379000`,
		},
		{
			`2`, `30`,
			`2609044274`,
		},
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, kMirror, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-268/problems/sum-of-k-mirror-numbers/