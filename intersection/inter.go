package intersection

func GetIntersection(xs, ys []string) []string {
	if len(xs) == 0 {
		return ys
	}
	m := make(map[string]bool)
	iTags := make([]string, 0, len(xs))
	for _, ft := range xs {
		m[ft] = true
	}
	for _, ot := range ys {
		if _, ok := m[ot]; ok {
			iTags = append(iTags, ot)
		}
	}
	return iTags
}
