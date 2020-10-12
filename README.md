# sql_test

```sql
CREATE TABLE test 
(
    id         INT          NOT NULL AUTO_INCREMENT,
    address    VARCHAR(400) NOT NULL,
    fTime      DATETIME     NOT NULL,
    sTime      DATETIME,
    dTime      TIME,
    incident   VARCHAR(10),
    guilty     VARCHAR(10),
    mts        BOOLEAN,
    life       BOOLEAN,
    a1         BOOLEAN,
    beCloud    BOOLEAN,
    comment    VARCHAR(400) NOT NULL,
    crash      VARCHAR(400) NOT NULL,
    PRIMARY KEY(id)
);

DROP TABLE test;
```
```
```
go get github.com/go-sql-driver/mysql
go get github.com/360EntSecGroup-Skylar/excelize

C:\Users\artem.zhivushko\Downloads\migrate.windows-amd64.exe -path ./schema -database 'mysql://Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database' down
C:\Users\artem.zhivushko\Downloads\migrate.windows-amd64.exe -path ./schema -database 'mysql://Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database' up
```