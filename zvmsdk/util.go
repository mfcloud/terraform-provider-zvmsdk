package zvmsdk


func isFailed(res int) bool {
	if res >= 200 && res < 300 {
		return false
	}
	return true
}
