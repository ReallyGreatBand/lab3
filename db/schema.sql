-- Select "plants" scheme
set search_path to plants;

-- Create table
create table if not exists plants
(
	id serial not null
		constraint table_name_id_key
			unique,
	soil_moisture_level real not null,
	soil_data_timestamp timestamp not null
);

-- Create update soilDataTimestamp trigger
create or replace function aft_update()
  returns trigger as
$$
begin
update plants
    set soil_data_timestamp = now()
    where id = old.id;
return null;
end;

$$
language 'plpgsql';

drop trigger if exists update_time on plants;
create trigger update_time after update of soil_moisture_level on plants
    for each row
    execute procedure aft_update();

-- Insert demo data
insert into plants (id, soil_moisture_level, soil_data_timestamp)
values (default, 0.15, '2020-11-29 14:05:26.000000');

insert into plants (id, soil_moisture_level, soil_data_timestamp)
values (default, 0.23, '2020-11-29 13:02:39.000000');

insert into plants (id, soil_moisture_level, soil_data_timestamp)
values (default, 0.56, '2020-11-29 05:24:41.000000');
