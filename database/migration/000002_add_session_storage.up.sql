CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BYTEA NOT NULL,
    expiry TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX idx_sessions_expiry ON sessions(expiry);