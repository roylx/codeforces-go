package copypasta

import (
	"math"
	"sort"
)

// NOTE: Pass n+1 if you wanna search range [0,n]
// NOTE: 二分时特判下限！（例如 0）
// TIPS: 如果输出的不是二分值而是一个与之相关的值，可以在 return false/true 前记录该值

// 隐藏的二分 https://codeforces.ml/problemset/problem/1354/D

// sort.Ints 性能测试 https://codeforces.ml/contest/977/submission/75301978

// 有些 OJ 的 Go 版本过低，不支持 sort.Slice，只能用 sort.Sort
type _pair struct{ x, y int }
type pairSlice []_pair

func (p pairSlice) Len() int { return len(p) }
func (p pairSlice) Less(i, j int) bool {
	a, b := p[i], p[j]
	return a.x < b.x || a.x == b.x && a.y < b.y
}
func (p pairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func sortCollections() {
	lowerBound := sort.SearchInts
	upperBound := func(a []int, x int) int { return sort.Search(len(a), func(i int) bool { return a[i] > x }) }
	// 也可以通过 sort.SearchInts(a, x+1) 来搜索 upperBound
	// lowerBound-1 为 <x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的
	// upperBound-1 为 <=x 的最大值的下标（-1 表示不存在），存在多个最大值时下标取最大的

	// 若要二分的函数 f(x) 对于较小的 x 返回 true，较大的 x 返回 false，如何找到最大的使 f(x) == true 的 x？
	// 考虑二分 !f(x)，则二分结果是最小的使 f(x) == false 的 x，将其减一就得到了最大的使 f(x) == true 的 x
	// 由于要对结果减一，sort.Search 应传入 n+1
	// 注意判断 x 为 0 的情况，若 f(0) == false，则二分结果是 -1
	// 好题 https://atcoder.jp/contests/abc149/tasks/abc149_e
	{
		var n int
		v := sort.Search(n+1, func(x int) bool {
			// 注意判断 x 为 0 的情况
			cnt := 0
			// ...
			return !(cnt >= x)
		}) - 1
		_ = v
	}

	searchRange := func(l, r int, f func(int) bool) int {
		for l < r {
			m := (l + r) >> 1 // 注意 l+r 是否超 int
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}
	// 若 l>0，也可以这样写
	//sort.Search(r, func(x int) bool {
	//	if x < l {
	//		return false
	//	}
	//	...
	//})

	search64 := func(n int64, f func(int64) bool) int64 {
		// Define f(-1) == false and f(n) == true.
		i, j := int64(0), n
		for i < j {
			h := (i + j) >> 1
			if f(h) {
				j = h
			} else {
				i = h + 1
			}
		}
		return i
	}

	// NOTE: step 取多少合适：
	// 如果返回结果不是答案的话，注意误差对答案的影响（由于误差累加的缘故，某些题目误差对查案的影响可以达到 n=2e5 倍，见 CF578C）

	binarySearch := func(l, r float64, f func(x float64) bool) float64 {
		step := int(math.Log2((r - l) / eps)) // eps 取 1e-8 比较稳妥（一般来说是保留小数位+2）
		for i := 0; i < step; i++ {
			mid := (l + r) / 2
			if f(mid) {
				r = mid // 减小 x
			} else {
				l = mid // 增大 x
			}
		}
		return (l + r) / 2
	}

	// 实数三分
	// NOTE: 多个下凸函数的乘积仍然是下凸函数；上凸同理 AT130F
	// https://codeforces.ml/blog/entry/60702
	// 模板题 https://www.luogu.com.cn/problem/P3382
	// 题目推荐 https://cp-algorithms.com/num_methods/ternary_search.html#toc-tgt-4
	ternarySearch := func(l, r float64, f func(x float64) float64) float64 {
		step := int(math.Log((r-l)/eps) / math.Log(1.5)) // eps 取 1e-8 比较稳妥（一般来说是保留小数位+2）
		for i := 0; i < step; i++ {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2 // 若求最大值写成 l = m1
			} else {
				l = m1 // 若求最大值写成 r = m2
			}
		}
		return (l + r) / 2
	}

	// 整数三分
	// NOTE: 若有大量相同的离散点，该方法在某些数据下会失效（例如三分的时候把存在最小值的「洼地」 skip 了）
	// https://codeforces.ml/problemset/problem/1301/B (只是举例，不用三分也可做)
	ternarySearchInt := func(l, r int, f func(x int) int) int {
		for r-l > 4 { // 最小区间长度根据题目可以扩大点
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2 // 若求最大值写成 l = m1
			} else {
				l = m1 // 若求最大值写成 r = m2
			}
		}
		min, minI := f(l), l
		for i := l + 1; i <= r; i++ {
			if v := f(i); v < min {
				min, minI = v, i
			}
		}
		return minI
	}

	// 0-1 分数规划
	// 与 0-1 背包结合，即最优比率背包
	// 与生成树结合，即最优比率生成树
	// 与负环判定结合，即最优比率环
	// 与网络流结合，即最大密度子图
	// 与费用流结合，即最优比率流
	// 与其他的各种带选择的算法乱套，即最优比率啥啥的
	// https://oi-wiki.org/misc/frac-programming/
	// todo https://www.luogu.com.cn/blog/yestoday/post-01-fen-shuo-gui-hua-yang-xie

	// CDQ 分治
	// todo https://oi-wiki.org/misc/cdq-divide/
	//      https://www.bilibili.com/video/BV1mC4y1s7ic
	//      [学习笔记]CDQ分治和整体二分 https://www.luogu.com.cn/blog/Owencodeisking/post-xue-xi-bi-ji-cdq-fen-zhi-hu-zheng-ti-er-fen
	//      https://www.luogu.com.cn/blog/ljc20020730/cdq-fen-zhi-xue-xi-bi-ji

	// 整体二分
	// todo https://oi-wiki.org/misc/parallel-binsearch/
	//      https://codeforces.ml/blog/entry/45578

	_ = []interface{}{
		lowerBound, upperBound,
		searchRange, search64, binarySearch, ternarySearch, ternarySearchInt,
	}
}
