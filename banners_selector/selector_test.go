package bannersselector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelector(t *testing.T) {
	t.Run("zero counts", func(t *testing.T) {
		totalShowsCount := 0
		data := make([]BannerInfo, 10)
		for i := 0; i < 10; i++ {
			data[i].Id = i + 1
		}

		for i := 0; i < len(data); i++ {
			selectedId := SelectNext(data, totalShowsCount)
			require.Equal(t, data[i].Id, selectedId)
			totalShowsCount++
			data[i].ShowsCount++
		}
	})

	t.Run("main formula", func(t *testing.T) {
		data := [][]BannerInfo{
			{
				{Id: 1, ShowsCount: 5, ClicksCount: 4},
				{Id: 2, ShowsCount: 2, ClicksCount: 1},
				{Id: 3, ShowsCount: 3, ClicksCount: 1},
				{Id: 4, ShowsCount: 1, ClicksCount: 0},
			},
			{
				{Id: 1, ShowsCount: 5, ClicksCount: 4},
				{Id: 2, ShowsCount: 4, ClicksCount: 1},
				{Id: 3, ShowsCount: 3, ClicksCount: 1},
				{Id: 4, ShowsCount: 4, ClicksCount: 1},
			},
			{
				{Id: 1, ShowsCount: 50, ClicksCount: 4},
				{Id: 2, ShowsCount: 17, ClicksCount: 4},
				{Id: 3, ShowsCount: 13, ClicksCount: 7},
				{Id: 4, ShowsCount: 7, ClicksCount: 2},
				{Id: 5, ShowsCount: 1, ClicksCount: 1},
				{Id: 6, ShowsCount: 52, ClicksCount: 21},
			},
		}

		selectedId := SelectNext(data[0], 11)
		require.Equal(t, 4, selectedId)

		selectedId = SelectNext(data[1], 11)
		require.Equal(t, 1, selectedId)

		selectedId = SelectNext(data[2], 140)
		require.Equal(t, 5, selectedId)
	})
}
