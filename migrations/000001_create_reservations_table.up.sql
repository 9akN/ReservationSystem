CREATE TABLE reservations (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    phone       VARCHAR(50)  NOT NULL,
    date        DATE         NOT NULL,
    time        TIME         NOT NULL,
    guests      INT          NOT NULL,
    notes       TEXT,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);
