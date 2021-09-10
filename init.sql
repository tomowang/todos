-- psql -U postgres

-- create database
create database todos;
-- create user
create user todos with encrypted password 'secret';
-- change owner
alter database todos owner to todos;
-- or grant all privileges on database hweal to hweal;
