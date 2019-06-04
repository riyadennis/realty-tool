-- noinspection SqlNoDataSourceInspectionForFile

ALTER TABLE property_data
ADD column price_updated DATETIME,
ADD column new_price varchar(100),
ADD column status_updated DATETIME;
