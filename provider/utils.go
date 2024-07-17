package provider

func contains[E any](list []E, fn func(e E) bool) bool {
	for _, e := range list {
		if fn(e) {
			return true
		}
	}
	return false
}
