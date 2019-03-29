CREATE TABLE Person (
    Personid int AUTO_INCREMENT,
    LastName varchar(255) NOT NULL,
    FirstName varchar(255),
    Age int,
    UNIQUE KEY(LastName),
    PRIMARY KEY (Personid),
);

INSERT INTO Persons VALUES (1, 'Jane', 'Doe', '23');

ALTER TABLE orders DROP COLUMN( total );

CREATE TABLE referrals (
    referred_by VARCHAR(255),
    who VARCHAR(255),
);

INSERT INTO Referrals (refered_by, who) VALUES
    ('Doe', 'Paul');
    ('Doe', 'Melissa');
    ('Doe', 'Jose');
    ('Doe', 'Doe');

CREATE TABLE Orders (
    id_order INT AUTO_INCREMENT,
    date TIMESTAMP,
    total TINYINT,

)

CREATE TABLE user_passwords (
    username VARCHAR(255),
    password VARCHAR(255),
)





DROP TABLE garbage;


CREATE USER 'admin'@'*' IDENTIFIED BY 'secret';     
CREATE USER 'test'@'*' IDENTIFIED  BY 'test';

GRANT ALL PRIVILEGES ON *.* TO 'test'@'localhost' IDENTIFIED BY 'password';

