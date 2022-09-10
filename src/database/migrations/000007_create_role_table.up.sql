-- create table
CREATE TABLE IF NOT EXISTS roles(
    id CHAR(36) PRIMARY KEY NOT NULL,
    name CHAR(100) UNIQUE NOT NULL
);

-- insert data
insert into roles values
('admin', 'Admin'),
('user', 'User')