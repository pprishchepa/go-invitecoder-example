CREATE TABLE invite_user
(
    email       TEXT        NOT NULL PRIMARY KEY,
    invited_via TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
