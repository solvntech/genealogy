CREATE TABLE IF NOT EXISTS person_statuses(
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

insert into person_statuses values
('1', 'Còn sống'),
('2', 'Đã qua đời'),
('3', 'Chưa xác định');
