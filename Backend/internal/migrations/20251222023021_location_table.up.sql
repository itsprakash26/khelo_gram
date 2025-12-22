CREATE TABLE locations (

id serial primary key,
user_id integer not null,
pin varchar(6) not null,
state text  not null default "west-bengal",
city text not null default "kolkata",



);