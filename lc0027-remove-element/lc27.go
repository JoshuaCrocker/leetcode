package lc0027_remove_element

func removeElement(nums []int, val int) int {
	front, back := 0, len(nums)-1

	k := 0
	for front <= back {
		if nums[front] == val {
			tmp := nums[back]
			nums[back] = nums[front]
			nums[front] = tmp
			back--
		} else {
			front++
			k++
		}
	}

	return k
}
