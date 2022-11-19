package main

//输入：nums = [3,6,2,7,1], k = 6
//输出：4
//解释：以 6 为最小公倍数的子数组是：
//[3, 6, 2]
//[3, 6]
//[2, 6]
//[6]
//[2, 3]

func subarrayLCM(nums []int, k int) int {
	res := [][]int{}

	for _, num := range nums {
		ret := []int{}
		if isCM(num, k) {
			ret = append(ret, num)
		}

		res = append(res, ret)
	}

}

// isCM 是否是公倍数
func isCM(input int, k int) bool {
	if input > k {
		return false
	}

	if k%input == 0 {
		return true
	}

	return false
}
