CREATE TABLE fizzbuzz (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    int1          INT NOT NULL,
    int2          INT NOT NULL,
    str1          TEXT NOT NULL,
    str2          TEXT NOT NULL,
    request_count INT NOT NULL DEFAULT 1,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT fizzbuzz_unique_params UNIQUE (int1, int2, str1, str2)
);