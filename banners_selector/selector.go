package bannersselector

import "math"

type BannerInfo struct {
	Id          int
	ShowsCount  int
	ClicksCount int
}

func SelectNext(banners []BannerInfo, commonCount int) int {
	var maxValue float64
	var maxValueBannerId int
	for _, banner := range banners {
		if banner.ShowsCount == 0 {
			return banner.Id
		}

		bannerValue := (float64(banner.ClicksCount) / float64(banner.ShowsCount)) +
			math.Sqrt((2*math.Log(float64(commonCount)))/float64(banner.ShowsCount))
		if bannerValue > maxValue {
			maxValue = bannerValue
			maxValueBannerId = banner.Id
		}
	}

	return maxValueBannerId
}
