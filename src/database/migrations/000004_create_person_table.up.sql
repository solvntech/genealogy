CREATE TABLE IF NOT EXISTS persons(
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

ALTER TABLE persons add constraint fk_father_id foreign key (father_id) references persons(id);
ALTER TABLE persons add constraint fk_mother_id foreign key (mother_id) references persons(id);
ALTER TABLE persons add constraint fk_status_id foreign key (status_id) references person_statuses(id);
ALTER TABLE persons add constraint fk_position_title_id foreign key (position_title_id) references position_titles(id);