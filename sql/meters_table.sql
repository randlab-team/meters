CREATE TABLE meters
(
    id            SERIAL PRIMARY KEY,
    sn            INTEGER,
    correct       BOOL,
    param_name    VARCHAR(16),
    "index"       INTEGER,
    date_register TIMESTAMP,
    "value"       INTEGER,
    log_interval  INTEGER,
    status        INTEGER
);

CREATE UNIQUE INDEX meters_unique_sn_date_reg
    ON meters (sn, date_register);