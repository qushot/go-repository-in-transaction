CREATE TABLE users (
    id int NOT NULL,
    name varchar(10) NOT NULL,
    address varchar(10) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (id, name, address)
VALUES (1, 'Yamada', 'Tokyo'),
       (2, 'Tanaka', 'Kanagawa'),
       (3, 'Sato', 'Chiba')
;
