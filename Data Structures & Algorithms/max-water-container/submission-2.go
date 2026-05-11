func maxArea(heights []int) int {
	res := 0
	i, j := 0, len(heights) - 1
	for i < j {
		area := (j - i) * minHeight(heights[i], heights[j])
		if area > res {
			res = area
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func minHeight(i, j int) int {
	if i < j {
		return i
	}
	return j
}