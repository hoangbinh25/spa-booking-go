CREATE TABLE branches (
    id         UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    spa_id     UUID          NOT NULL,
    code       VARCHAR (50)  NOT NULL,
    name       VARCHAR (150) NOT NULL,
    phone      VARCHAR (20) ,
    email      VARCHAR (150),
    address    TEXT         ,
    timezone   VARCHAR (50)  DEFAULT 'Asia/Ho_Chi_Minh' NOT NULL,
    status     VARCHAR (20)  DEFAULT 'ACTIVE' NOT NULL,
    created_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_branches_spa FOREIGN KEY (spa_id) REFERENCES spas (id),
    CONSTRAINT uq_branches_spa_code UNIQUE (spa_id, code)
);