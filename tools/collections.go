package tools

func FilterStringSlice(slice []string, f func(string) bool) []string {
	fm := make([]string, 0)
	for _, v := range slice {
		if f(v) {
			fm = append(fm, v)
		}
	}
	return fm
}
