package selector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelector(t *testing.T) {
	t.Run("zero counts", func(t *testing.T) {
		totalShowsCount := 0
		data := make([]BannerData, 10)
		for i := 0; i < 10; i++ {
			data[i].Id = i + 1
		}

		for i := 0; i < len(data); i++ {
			selected := SelectNext(data, totalShowsCount)
			require.Equal(t, data[i].Id, selected.Id)
			totalShowsCount++
			data[i].ShowsCount++
		}
	})

	t.Run("main formula", func(t *testing.T) {
		data := [][]BannerData{
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

		selected := SelectNext(data[0], 11)
		require.Equal(t, 4, selected.Id)

		selected = SelectNext(data[1], 11)
		require.Equal(t, 1, selected.Id)

		selected = SelectNext(data[2], 140)
		require.Equal(t, 5, selected.Id)
	})
}
