package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DenKuzin33/OtusProject/banners_selector/selector"
	_ "github.com/jackc/pgx/stdlib"
)

var db = initConnection()

func initConnection() *sql.DB {
	dsn := "host=localhost port=5432 dbname=banners_rotator user=user password=user"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("failed to load driver: %v", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("failed to connect to db: %v", err)
	}
	return db
}

func GetSlot(id int) (Entity, error) {
	query := "select * from slots where id=$1"
	return getEntity(query, id)
}

func GetBanner(id int) (Entity, error) {
	query := "select * from banners where id=$1"
	return getEntity(query, id)
}

func GetDemGroup(id int) (Entity, error) {
	query := "select * from dem_groups where id=$1"
	return getEntity(query, id)
}

func AddBannerToSlot(ctx context.Context, bannerId int, slotId int) error {
	if _, err := GetBanner(bannerId); err != nil {
		return fmt.Errorf("Banner with id %d does not exist", bannerId)
	}

	if _, err := GetSlot(slotId); err != nil {
		return fmt.Errorf("Slot with id %d does not exist", slotId)
	}

	query := "INSERT INTO shows_clicks_count VALUES ($1, $2)"
	_, err := db.ExecContext(ctx, query, slotId, bannerId)
	return err
}

func RemoveBannerFromSlot(ctx context.Context, bannerId int, slotId int) error {
	query := `DELETE FROM slot_banners WHERE banner_id=$1 and slot_id=$2;
	DELETE FROM shows_clicks_count WHERE banner_id=$1 and slot_id=$2`
	_, err := db.ExecContext(ctx, query, bannerId, slotId)
	return err
}

func GetSlotBanners(ctx context.Context, slotId int64, demGroupId int64) ([]selector.BannerData, int, error) {
	banners := []selector.BannerData{}
	totalShowsCount := 0
	query := "SELECT slot_banners.banner_id, dem_group_id, shows_count, clicks_count FROM slot_banners left join shows_clicks_count on slot_banners.slot_id=shows_clicks_count.slot_id and slot_banners.banner_id=shows_clicks_count.banner_id and dem_group_id=$1 where slot_banners.slot_id=$2"
	//query := "SELECT * FROM shows_clicks_count WHERE slot_id=$1 AND dem_group_id IN (-1, $2)"
	rows, err := db.QueryContext(ctx, query, demGroupId, slotId)
	if err != nil {
		return banners, totalShowsCount, err
	}
	defer rows.Close()

	var banner selector.BannerData
	for rows.Next() {
		if err := rows.Scan(&banner.Id, &banner.ShowsCount, &banner.ClicksCount); err != nil {
			return banners, totalShowsCount, err
		}
		banners = append(banners, banner)
		totalShowsCount += banner.ShowsCount
	}
	return banners, totalShowsCount, nil
}

func UpdateBannerData(ctx context.Context, banner selector.BannerData, slotId int, demGroupId int) error {
	query := "UPDATE shows_clicks_count SET shows_count=$1, click_count=$2 WHERE slot_id=$3 and banner_id=$4 and dem_group_id=$5"
	_, err := db.ExecContext(ctx, query, banner.ShowsCount, banner.ClicksCount, slotId, banner.Id, demGroupId)
	return err
}

func IncrementClicksCount(ctx context.Context, slotId int, bannerId int, demGroupId int) error {
	query := "UPDATE shows_clicks_count SET click_count=click_count+1 WHERE slot_id=$1 and banner_id=$2 and dem_group_id=$3"
	_, err := db.ExecContext(ctx, query, slotId, bannerId, demGroupId)
	return err
}

func getEntity(query string, id int) (Entity, error) {
	result := Entity{}
	rows, err := db.Query(query, id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&result.id, &result.description); err != nil {
			return result, err
		}
	}
	return result, nil
}
