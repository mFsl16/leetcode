package longestcommonprefix

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return prefix[0:i]
			}
		}
	}

	return prefix
}
