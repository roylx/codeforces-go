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
			`"7+3*1*2"`, `[20,13,42]`, 
			`7`,
		},
		{
			`"3+5*2"`, `[13,0,10,13,13,16,16]`, 
			`19`,
		},
		{
			`"6+0*1"`, `[12,9,6,4,8,6]`, 
			`10`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, scoreOfStudents, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-260/problems/the-score-of-students-solving-math-expression/
