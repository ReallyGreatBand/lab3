set search_path to "Plants"; -- select "Plants" scheme

-- Create tables.
drop table if exists "Plants";
create table "Plants"
(
    id serial unique not null,
    "soilMoistureLevel" real not null,
    "soilDataTimestamp" timestamp not null
);

-- Create update soilDataTimestamp trigger
create or replace function aft_update()
  returns trigger as
$$
begin
update "Plants"
    set "soilDataTimestamp" = now()
    where id = old.id;
return null;
end;

$$
language 'plpgsql';

drop trigger update_time on "Plants";
create trigger update_time after update of "soilMoistureLevel" on "Plants"
    for each row
    execute procedure aft_update();

-- Insert demo data.
insert into "Plants" (id, soilMoistureLevel, soilDataTimestamp)
values (default, 0.15, '2020-11-29 14:05:26.000000');

insert into "Plants" (id, soilMoistureLevel, soilDataTimestamp)
values (default, 0.23, '2020-11-29 13:02:39.000000');

insert into "Plants" (id, soilMoistureLevel, soilDataTimestamp)
values (default, 0.56, '2020-11-29 05:24:41.000000');