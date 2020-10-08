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
673(IP), 676(IP), 698(IP), 1105(IP), 2629(IP), 2679(IP), 2687(IP), 2696(IP), 2697(IP), 2698(IP), 4076(IP), 4680(IP), 4693(IP), 4696(IP), 4697(IP), 5698(IP), 5961(IP), 2621(IP), 5387(IP), 5417(IP), 1105(UMTS), 4076(UMTS), 1408(IP), 2698(LTE), 1659(IP), 1659(LTE), 4076(LTE), 5698(LTE), 673(LTE), 698(LTE), 2683(LTE), 2681(LTE), 1408(LTE), 675(LTE), 2675(LTE), 5675(LTE), 4351(LTE), 4076(FTTX), 4076(FTTX) *