-- create table
CREATE TABLE IF NOT EXISTS users(
    id CHAR(36) PRIMARY KEY NOT NULL,
    email CHAR(100) UNIQUE NOT NULL,
    password CHAR(100) NOT NULL,
    role_id CHAR(36) NOT NULL,
    created_at BIGINT,
    updated_at BIGINT
);

-- create relationship
ALTER TABLE users add constraint fk_role_id foreign key (role_id) references roles(id);

-- insert data
insert into users(id, email, password, role_id) values
('267249fe-3005-435c-bd4e-f4b1b5b61dc6', 'superuser', '$2a$14$r3tDnIK00VzRzWHZNzaiz.Qgw9w.NgmFxK8NZGBDnVNgjPgikjZHS', 'admin')