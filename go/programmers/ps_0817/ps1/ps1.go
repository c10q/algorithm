package ps1

const (
	RJelly = 100
	SJelly = 50
	YJelly = 5
	NJelly = 1
)

func solution(score int) int {
	if score < 1 {
		return 0
	}

	var result int
	for true {
		if score >= RJelly {
			RJellyCount := score / RJelly
			result += RJellyCount
			score -= RJellyCount * RJelly
		}
		if score >= SJelly {
			SJellyCount := score / SJelly
			result += SJellyCount
			score -= SJellyCount * SJelly
		}
		if score >= YJelly {
			YJellyCount := score / YJelly
			result += YJellyCount
			score -= YJellyCount * YJelly
		}
		if score >= NJelly {
			NJellyCount := score / NJelly
			result += NJellyCount
			score -= NJellyCount * NJelly
		}
		if score == 0 {
			break
		}
	}

	return result
}
