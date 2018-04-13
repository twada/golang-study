package basename1

func basename(s string) string {
	// 最後の '/' とその前のすべてを破棄する
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 最後の '.' より前の全てを保持する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
