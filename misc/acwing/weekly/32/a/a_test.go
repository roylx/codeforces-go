// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`001001`,
			`NO`,
		},
		{
			`1000000001`,
			`YES`,
		},
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
