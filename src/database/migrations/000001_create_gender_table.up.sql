create table if not exists genders(
   id CHAR(36) PRIMARY KEY NOT NULL,
   name VARCHAR(255) NOT NULL
);

insert into genders values
('1', 'Nam'),
('2', 'Nữ');