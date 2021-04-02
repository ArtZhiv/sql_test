CREATE TABLE eNodeB(
    id INT NOT NULL AUTO_INCREMENT,
    number INT NOT NULL,
    dismantling VARCHAR(200),
    area VARCHAR(50) NOT NULL,
    district VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    address VARCHAR(400) NOT NULL,
    vendor VARCHAR(10) NOT NULL,
    location VARCHAR(10),
    commercial_mts BOOLEAN NOT NULL,
    commercial_life BOOLEAN NOT NULL,
    commercial_a1 BOOLEAN NOT NULL,
    PRIMARY KEY(id)
); 

CREATE TABLE sector(
    id INT NOT NULL AUTO_INCREMENT,
    basestantion INT NOT NULL,
    cell_number VARCHAR(20) NOT NULL,
    bandwidth INT NOT NULL,
    mts BOOLEAN NOT NULL,
    life BOOLEAN NOT NULL,
    a1 BOOLEAN NOT NULL,
    beCloud BOOLEAN NOT NULL,
    PRIMARY KEY(id)
);