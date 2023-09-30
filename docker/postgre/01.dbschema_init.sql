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

CREATE TABLE shows_clicks_count (
    slot_id integer,
    banner_id integer,
    dem_group_id integer,
    shows_count integer,
    click_count integer
);