CREATE TABLE fizzbuzz (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    int1        INT NOT NULL,
    int2        INT NOT NULL,
    str1        TEXT NOT NULL,
    str2        TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);