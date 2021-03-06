/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<uint64(len(nums)))
	dfs(nums, 0, &res, []int{})
	return res
}

func dfs(nums []int, pos int, res *[][]int, tmpSlice []int) {
	if pos == len(nums) {
		*res = append(*res, append([]int{}, tmpSlice...))
		return
	}

	dfs(nums, pos+1, res, tmpSlice)
	tmpSlice = append(tmpSlice, nums[pos])
	dfs(nums, pos+1, res, tmpSlice)
}

func subsets1(nums []int) [][]int {
	resCnt := 1 << uint64(len(nums))
	res := make([][]int, resCnt)
	for i := 0; i < resCnt; i++ {
		sl := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i&(1<<uint64(j)) != 0 {
				sl = append(sl, nums[j])
			}
		}
		res[i] = sl
	}

	return res
}
// @lc code=end

