CREATE TABLE banners (
    id serial primary key,
    description text
);

CREATE TABLE slots (
    id serial primary key,
    description text
);

CREATE TABLE dem_groups (
    id serial primary key,
    description text
);

CREATE TABLE slot_banners (
    slot_id integer,
    banner_id integer,
)

CREATE TABLE shows_clicks_count (
    slot_id integer,
    banner_id integer,
    dem_group_id integer default -1,
    shows_count integer default 0,
    click_count integer default 0
);