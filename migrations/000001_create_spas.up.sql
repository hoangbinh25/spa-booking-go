CREATE TABLE spas (
    id         UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    name       VARCHAR (150) NOT NULL,
    phone      VARCHAR (20) ,
    email      VARCHAR (150),
    status     VARCHAR (20)  DEFAULT 'ACTIVE' NOT NULL,
    created_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL
);