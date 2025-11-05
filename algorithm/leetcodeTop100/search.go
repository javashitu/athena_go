package leetcodetop100

type Search struct{}

func (s Search) search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		//这个解法这里非常丑陋，因为我在下面的判断时是根据mid和target的大小关系来决定收缩左边界还是右边界，但是因为我收缩时是没有比较high的大小的，所以这里必须有这么一个判断
		//更加优雅的解法是在下面的收缩边界时判断mid处在的位置是左边是有序的还是右边是有序的，然后决定收缩左边界还是右边界，因为这样可以在收缩时就比较左右边界和target的大小来决定要不要返回
		if nums[high] == target {
			return high
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//如果mid比target大，那么需要找到比mid小的范围
			//如果处在交接点右侧，那么只有mid的左侧有比mid更小的数
			if nums[mid] < nums[high] {
				high = mid - 1
				continue
			}
			//如果mid处在交界点左侧，那么mid左侧和mid右侧均有比mid小的数
			//那么这个时候就需要比较target和high的大小，如果target比high位置大，那么要找的范围在mid左侧，如果target比high位置小，则说明在mid右侧
			if target > nums[high] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			//如果mid比target小，那么需要找到比mid更大的范围
			//如果mid处在交界处左侧，那么需要在mid右侧范围找，且只有一段mid到交界为止的范围是大于target的
			if nums[mid] > nums[low] {
				low = mid + 1
				continue
			}
			//如果mid处在交界处右侧，那么在mid的左侧范围内有比mid更大的，且右侧范围也有比mid更大的数
			if nums[high] > target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
