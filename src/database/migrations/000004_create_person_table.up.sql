CREATE TABLE IF NOT EXISTS people(
    id CHAR(36) PRIMARY KEY NOT NULL,
    birth_name VARCHAR(255) NOT NULL,
    gene_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255),
    birthday bigint,
    father_id CHAR(36),
    mother_id CHAR(36),
    position_title_id CHAR(36) NOT NULL,
    email VARCHAR(225),
    phone_number VARCHAR(225),
    address VARCHAR(225),
    descriptions TEXT,
    status_id CHAR(36) NOT NULL
);

ALTER TABLE people add constraint fk_father_id foreign key (father_id) references people(id);
ALTER TABLE people add constraint fk_mother_id foreign key (mother_id) references people(id);
ALTER TABLE people add constraint fk_status_id foreign key (status_id) references person_statuses(id);
ALTER TABLE people add constraint fk_position_title_id foreign key (position_title_id) references position_titles(id);

insert into people (id, birth_name, gene_name, birthday, father_id, position_title_id, status_id) values
('131d0699-d7d9-48ba-9cbd-36343035687a', 'Nguyễn Đức Liêm', 'Nguyễn Đức Liêm', null, null, 'thunam', 'quadoi'),
('29c81192-33cd-48a8-a4f1-60a4f56fb150', 'Nguyễn Đức Thắng', 'Nguyễn Đức Thắng', null, '131d0699-d7d9-48ba-9cbd-36343035687a', 'thunam', 'quadoi'),
('8d193965-75a8-4a5d-8712-83fdb4d73499', 'Nguyễn Đức Hinh', 'Nguyễn Đức Hinh', null, '131d0699-d7d9-48ba-9cbd-36343035687a', 'thunam', 'quadoi'),
('783c7b5e-f5e0-4c89-927d-9e9e9818bd99', 'Nguyễn Đức Ngữ', 'Nguyễn Đức Ngữ', null, '8d193965-75a8-4a5d-8712-83fdb4d73499', 'truongnam', 'consong'),
('9c1d010c-45ad-4871-be01-03071d62d106', 'Nguyễn Đức Thành Luận', 'Nguyễn Đức Luận', null, '783c7b5e-f5e0-4c89-927d-9e9e9818bd99', 'truongnam', 'consong'),
('99869989-05ff-4c62-890a-07cb38e52a25', 'Nguyễn Đức Thành Phát', 'Nguyễn Đức Phát', null, '783c7b5e-f5e0-4c89-927d-9e9e9818bd99', 'utnam', 'consong'),
('4d932b25-6246-45dd-b988-2678487d7f56', 'Nguyễn Đức Hân', 'Nguyễn Đức Hân', 25545600000, '8d193965-75a8-4a5d-8712-83fdb4d73499', 'thunam', 'consong'),
('2205adcb-e7d1-4685-98d3-7a38485ab9fa', 'Nguyễn Đức Hải', 'Nguyễn Đức Hải', 901472400000, '4d932b25-6246-45dd-b988-2678487d7f56', 'truongnam', 'consong'),
('d796c4db-fad6-4d49-a924-99e5e3f44278', 'Nguyễn Mỹ Duyên', 'Nguyễn Mỹ Duyên', 1206637200000, '4d932b25-6246-45dd-b988-2678487d7f56', 'utnu', 'consong'),
('2cd070ce-e3ae-4a3c-bc63-a48d129d0664', 'Nguyễn Hồng Vân', 'Nguyễn Hồng Vân', 1050598800000, '4d932b25-6246-45dd-b988-2678487d7f56', 'thunu', 'consong')