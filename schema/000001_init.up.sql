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

CREATE TABLE usersTable (
        id         INT          NOT NULL AUTO_INCREMENT,
        Username   VARCHAR(50)  NOT NULL,
        Firstname   VARCHAR(50)  NOT NULL,
        Lastname   VARCHAR(50)  NOT NULL,
        Middlename   VARCHAR(50)  NOT NULL,
        PRIMARY KEY(id)
  );



CREATE TABLE lte(
  id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  Name  TEXT,
  Start_time TIMESTAMP NOT NULL,
  End_time TIMESTAMP,
  Time TIMESTAMP,
  Type TEXT,
  Responsibility TEXT,
  Service TEXT,
  Note TEXT,
  Alarm TEXT,
  FIO TEXT
);

CREATE TABLE egts(
  id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  Start_time TIMESTAMP,
  End_time TIMESTAMP,
  Time TIMESTAMP,
  Type TEXT,
  Responsibility TEXT,
  Service TEXT,
  Note TEXT,
  FIO TEXT
);

CREATE TABLE ppd(
  id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
  Name TEXT,
  Start_time TIMESTAMP,
  End_time TIMESTAMP,
  Time TIMESTAMP,
  Responsibility TEXT,
  Note TEXT,
  FIO TEXT
);

CREATE TABLE commercialUPD( id INT NOT NULL AUTO_INCREMENT, dateCreate TEXT NOT NULL, PRIMARY KEY(id) ) 