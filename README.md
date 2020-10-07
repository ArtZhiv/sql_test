# sql_test

``` sql
CREATE TABLE eNodeB (
    E_id       INT AUTO_INCREMENT,
    number     INT NOT NULL,
    address    VARCHAR(400) NOT NULL,
	vendor     VARCHAR(10) NOT NULL,
	region     VARCHAR(50) NOT NULL,
	province   VARCHAR(50) NOT NULL,
	demolition VARCHAR(200),
	erspd      BOOLEAN,
	place      VARCHAR(10),
    PRIMARY KEY(E_id)
);
CREATE TABLE sector (
    S_id       INT AUTO_INCREMENT,
    number     INT NOT NULL,
    sector     VARCHAR(20) NOT NULL,
	band       INT NOT NULL,
	mts        BOOLEAN,
    life       BOOLEAN,
    a1         BOOLEAN,
    beCloud    BOOLEAN,
    PRIMARY KEY(S_id)
);
```
```
go get github.com/go-sql-driver/mysql
go get github.com/360EntSecGroup-Skylar/excelize
```
