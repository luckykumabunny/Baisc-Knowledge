func findPeakElement(nums []int) int {
    n := len(nums)
    if n <=1 {
        return 0
    }

    // 辅助函数便于处理边界值
    get := func(i int) int{
        if i == -1 || i == n {
            return math.MinInt64
        }
        return nums[i]
    }

    low := 0
    high := n-1
    for high > low {
        mid := low + (high-low)/2
        if get(mid) > get(mid-1 ) && get(mid) > get(mid+1) {
            return mid
        }
        if get(mid-1) > get(mid){
            high = mid-1
        }else{
            low = mid+1
        }
    }
    return low
}