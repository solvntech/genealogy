create table if not exists position_titles(
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    gender_id CHAR(36) NOT NULL
);

ALTER TABLE position_titles add foreign key (gender_id) references genders(id);

insert into position_titles values
('truongnam', 'Trưởng nam', 'male'),
('thunam', 'Thứ nam', 'male'),
('utnam', 'Út nam', 'male'),
('truongnu', 'Trưởng nữ', 'female'),
('thunu', 'Thứ nữ', 'female'),
('utnu', 'Út nữ', 'female'),
('re', 'Rễ', 'male'),
('dau', 'Dâu', 'female');