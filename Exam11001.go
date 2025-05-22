package resources

import "sort"

/*
在公司饭堂里新加一个卖包子的窗口，老板一笼可以蒸N个包子。有人来买包子时，如果一个人没买完一笼，剩下的老板会放冰箱里卖给下一个人。
下一个人来买包子，老板先从冰箱里拿出来卖，卖完了再去蒸。如果买到了从冰箱拿出来的，那这个人就买到了不新鲜的包子。
在窗口排队的同事，每个人都想买特定多个包子，老板为了让大部分人都买到新鲜的包子，会调换排队同事的位置
返回最多有几个人可以买到新鲜的包子。

老板一笼可以蒸n个包子；
排队的同事每人想买arr[i]个包子

输入：n = 3, arr = [1,2,3,4,5,6]
输出：4
解释：老板将同事排队顺序排为[6,2,4,5,1,3] 。那么第 1，2，4，6 个人都买到了新鲜的包子
*/
func (context ExamContext) doAnswer(param Param) Answer {
	// 将数组按购买量降序排序
	sort.Slice(param.arr, func(i, j int) bool {
		return param.arr[i] > param.arr[j]
	})

	remain := 0  // 当前冰箱剩余包子数
	cnt := 0     // 能买到新鲜包子的人数
	n := param.n // 每笼包子数

	for _, y := range param.arr {
		if remain == 0 { // 当前冰箱为空，必须蒸新笼
			cnt++
			// 计算需要蒸多少笼
			k := (y + n - 1) / n
			// 更新剩余包子数
			remain = k*n - y
		} else { // 冰箱有剩余，优先用冰箱的
			if y <= remain { // 冰箱里的足够满足需求
				remain -= y
			} else { // 需要补充蒸笼
				needed := y - remain
				k := (needed + n - 1) / n
				remain = k*n - needed
			}
		}
	}

	return Answer{rs1: cnt}
}
