CREATE TABLE eNodeB 
(
    id         INT          NOT NULL AUTO_INCREMENT,
    number     INT          NOT NULL,
    address    VARCHAR(400) NOT NULL,
	vendor     VARCHAR(10)  NOT NULL,
	region     VARCHAR(50)  NOT NULL,
	district   VARCHAR(50)  NOT NULL,
	demolition VARCHAR(200),
    mts        BOOLEAN      NOT NULL,
    life       BOOLEAN      NOT NULL,
    a1         BOOLEAN      NOT NULL,
	place      VARCHAR(10),
    PRIMARY KEY(id)
);

CREATE TABLE sector 
(
    id         INT          NOT NULL AUTO_INCREMENT,
    number     INT          NOT NULL,
    sector     VARCHAR(20)  NOT NULL,
	band       INT          NOT NULL,
	mts        BOOLEAN      NOT NULL,
    life       BOOLEAN      NOT NULL,
    a1         BOOLEAN      NOT NULL,
    beCloud    BOOLEAN      NOT NULL,
    PRIMARY KEY(id)
);

