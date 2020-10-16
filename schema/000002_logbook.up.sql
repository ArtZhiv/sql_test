CREATE TABLE incident
(
	id                  BIGINT UNSIGNED NOT NULL,
    NumInc3p            VARCHAR(50),
	Status              VARCHAR(20) NOT NULL,
	BriefDescription    TEXT NOT NULL,
	DetailedDescription MEDIUMTEXT NOT NULL,
	Classification      TINYINT UNSIGNED NOT NULL,
	Responsibility      VARCHAR(100) NOT NULL,
	IncidentStarted     DATETIME NOT NULL,
	IncidentEnded       DATETIME,
	CrashOccurred       DATETIME NOT NULL,
	DetectionMethod     VARCHAR(50) NOT NULL,
	Note                MEDIUMTEXT,
	Service             VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE works
(
	id                  BIGINT UNSIGNED NOT NULL,
	Status              VARCHAR(50) NOT NULL,
	BriefDescription    TEXT NOT NULL,
	DetailedDescription MEDIUMTEXT NOT NULL,
	Responsibility      VARCHAR(200) NOT NULL,
    Manufacturer        VARCHAR(200),
	PlannedStarted      DATETIME NOT NULL,
	PlannedEnded        DATETIME NOT NULL,
	ActualStarted       DATETIME,
    ActualEnded         DATETIME,
	Note                MEDIUMTEXT,
	Service             VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
);