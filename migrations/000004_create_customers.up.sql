CREATE TABLE customers (
    id            UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id       UUID         ,
    full_name     VARCHAR (150) NOT NULL,
    phone         VARCHAR (20) ,
    email         VARCHAR (150),
    gender        VARCHAR (20) ,
    date_of_birth DATE         ,
    note          TEXT         ,
    created_at    TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_customers_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT uq_customers_user UNIQUE (user_id)
);