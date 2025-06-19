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

func generateScoreFormat(info ScoreData) string {
	//do some magic here lol
	return "hello!"
}

func getMangaId(data *MangaData, name string) (int, string) {
	if theManga := getManga(data, name); len(theManga) > 0 {
		i := getRightMangaNum(theManga)
		first := theManga[i]
		return first.Id, first.Images.Jpg.Image_URL
	} else {
		return -1, "no image"
	}
}

func getRightMangaNum(mangas []Manga) int {
	curr := 0
	for i := 1; i < len(mangas); i++ {
		if mangas[curr].Members <= mangas[i].Members {
			curr = i
		}
	}
	return curr
}
