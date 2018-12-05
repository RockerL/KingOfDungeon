package shared

func Int32Abs(v int32) int32 {
	if v < 0 {
		return -v
	}
	return v
}
