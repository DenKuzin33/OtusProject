package selector

import "math"

type BannerData struct {
	Id          int
	ShowsCount  int
	ClicksCount int
}

func SelectNext(banners []BannerData, commonCount int) BannerData {
	var maxValue float64
	var maxValueBanner BannerData
	for _, banner := range banners {
		if banner.ShowsCount == 0 {
			return banner
		}

		bannerValue := (float64(banner.ClicksCount) / float64(banner.ShowsCount)) +
			math.Sqrt((2*math.Log(float64(commonCount)))/float64(banner.ShowsCount))
		if bannerValue > maxValue {
			maxValue = bannerValue
			maxValueBanner = banner
		}
	}

	return maxValueBanner
}
