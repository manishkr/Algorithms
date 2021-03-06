package main

func longestCommonSubsequenceBottomUp(text1 string, text2 string) int {
	lcsMap := make([][]int, 2)
	for i := 0; i < 2; i++ {
		lcsMap[i] = make([]int, len(text2)+1)
	}
	bi := 0
	for i := 0; i <= len(text1); i++ {
		bi = i & 1
		for j := 0; j <= len(text2); j++ {
			if i != 0 && j != 0 {
				if text1[i-1] == text2[j-1] {
					lcsMap[bi][j] = 1 + lcsMap[1-bi][j-1]
				} else {
					lcsMap[bi][j] = max(lcsMap[bi][j-1], lcsMap[1-bi][j])
				}
			}
		}
	}
	return lcsMap[bi][len(text2)]
}
