// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`5
1 2 3 2 1`,
			`1 2 3 2 1`,
		},
		{
			`3
10 6 8`,
			`10 6 6`,
		},
		
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}