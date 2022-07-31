create table if not exists marital_statuses(
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

insert into marital_statuses values
('1', 'Đã cưới'),
('2', 'Đã ly hôn');