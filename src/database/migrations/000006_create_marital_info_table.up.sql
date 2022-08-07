create table if not exists marital_infos(
    husband_id CHAR(36) NOT NULL,
    wife_id CHAR(36) NOT NULL,
    wedding_date CHAR(36),
    divorce_date CHAR(36),
    marital_status_id CHAR(36) not null,
    descriptions text,
    primary key (husband_id, wife_id)
);

ALTER TABLE marital_infos add constraint fk_marital_status_id foreign key (marital_status_id) references marital_statuses(id);
ALTER TABLE marital_infos add constraint fk_husband_id foreign key (husband_id) references people(id);
ALTER TABLE marital_infos add constraint fk_wife_id foreign key (wife_id) references people(id);
