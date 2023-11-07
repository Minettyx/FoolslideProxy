package templates

func Read(data []string) string {
	res := "<script>var pages = ["

	if len(data) > 0 {
		for _, v := range data {
			res += `{"url":"` + v + `"},`
		}

		res = res[:len(res)-1] // remove last comma
	}

	res += "];</script>"
	return res
}
