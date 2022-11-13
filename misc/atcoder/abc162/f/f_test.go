// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc162/submit?taskScreenName=abc162_f
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`6
1 2 3 4 5 6`,
			`12`,
		},
		{
			`5
-1000 -100 -10 0 10`,
			`0`,
		},
		{
			`10
1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000`,
			`5000000000`,
		},
		{
			`27
18 -28 18 28 -45 90 -45 23 -53 60 28 -74 -71 35 -26 -62 49 -77 57 24 -70 -93 69 -99 59 57 -49`,
			`295`,
		},
		{
			`199999
`,
			`99998995014475`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc162/tasks/abc162_f

func TestCompare(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 3)
		rg.NewLine()
		rg.IntSlice(n, -3, 3)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n+1)
		f := make([]int, n+1)
		sum := 0
		for i := 1; i < n+1; i++ {
			fmt.Fscan(in, &a[i])
			if i&1 != 0 {
				sum += a[i]
			}
			if i == 1 {
				continue
			}
			if i&1 != 0 {
				f[i] = max(f[i-2]+a[i], f[i-1])
			} else {
				f[i] = max(f[i-2]+a[i], sum)
			}
		}

		fmt.Fprintln(out , f[n])
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}