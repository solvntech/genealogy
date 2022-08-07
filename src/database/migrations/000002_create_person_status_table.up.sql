CREATE TABLE IF NOT EXISTS person_statuses(
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

insert into person_statuses values
('consong', 'Còn sống'),
('quadoi', 'Đã qua đời'),
('chuaxacdinh', 'Chưa xác định');
