CREATE TABLE users (
    id            UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    email         VARCHAR (150),
    phone         VARCHAR (20) ,
    password_hash VARCHAR (255),
    full_name     VARCHAR (150) NOT NULL,
    status        VARCHAR (20)  DEFAULT 'ACTIVE' NOT NULL,
    last_login_at TIMESTAMPTZ  ,
    created_at    TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT uq_users_email UNIQUE (email),
    CONSTRAINT uq_users_phone UNIQUE (phone)
);

CREATE TABLE roles (
    id         UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    code       VARCHAR (50)  NOT NULL,
    name       VARCHAR (100) NOT NULL,
    created_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT uq_roles_code UNIQUE (code)
);

CREATE TABLE user_roles (
    user_id    UUID        NOT NULL,
    role_id    UUID        NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
);