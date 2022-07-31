create table if not exists position_titles(
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    gender_id CHAR(36) NOT NULL
);

ALTER TABLE position_titles add foreign key (gender_id) references genders(id);

insert into position_titles values
('1', 'Trưởng nam', '1'),
('2', 'Thứ nam', '1'),
('3', 'Út nam', '1'),
('4', 'Trưởng nữ', '2'),
('5', 'Thứ nữ', '2'),
('6', 'Út nữ', '2'),
('7', 'Rễ', '1'),
('8', 'Dâu', '2');