package apiFolder

// File dedicated for helping the API main file do it's job
func calculateAverageRating(info ScoreData) float64 {
	sum := 0
	total := 0
	for i := 0; i < 10; i++ {
		sum += ((i + 1) * info.Scores[i].Votes)
		total += info.Scores[i].Votes
	}
	return float64(sum) / float64(total)
}
