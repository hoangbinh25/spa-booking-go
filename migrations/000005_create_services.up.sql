CREATE TABLE services (
    id                       UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    code                     VARCHAR (50)  NOT NULL,
    name                     VARCHAR (150) NOT NULL,
    description              TEXT         ,
    default_duration_minutes INT           NOT NULL,
    status                   VARCHAR (20)  DEFAULT 'ACTIVE' NOT NULL,
    created_at               TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at               TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT uq_services_code UNIQUE (code),
    CONSTRAINT chk_services_duration CHECK (default_duration_minutes > 0)
);

CREATE TABLE branch_services (
    id               UUID            DEFAULT gen_random_uuid() PRIMARY KEY,
    branch_id        UUID            NOT NULL,
    service_id       UUID            NOT NULL,
    price            NUMERIC (14, 2) NOT NULL,
    duration_minutes INT             NOT NULL,
    is_active        BOOLEAN         DEFAULT TRUE NOT NULL,
    created_at       TIMESTAMPTZ     DEFAULT NOW() NOT NULL,
    updated_at       TIMESTAMPTZ     DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_branch_services_branch FOREIGN KEY (branch_id) REFERENCES branches (id),
    CONSTRAINT fk_branch_services_service FOREIGN KEY (service_id) REFERENCES services (id),
    CONSTRAINT uq_branch_services UNIQUE (branch_id, service_id),
    CONSTRAINT chk_branch_services_price CHECK (price >= 0),
    CONSTRAINT chk_branch_services_duration CHECK (duration_minutes > 0)
);