CREATE TABLE technicians (
    id            UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id       UUID          NOT NULL,
    branch_id     UUID          NOT NULL,
    employee_code VARCHAR (50)  NOT NULL,
    display_name  VARCHAR (150) NOT NULL,
    bio           TEXT         ,
    status        VARCHAR (20)  DEFAULT 'ACTIVE' NOT NULL,
    created_at    TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_technicians_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_technicians_branch FOREIGN KEY (branch_id) REFERENCES branches (id),
    CONSTRAINT uq_technicians_user UNIQUE (user_id),
    CONSTRAINT uq_technicians_branch_employee_code UNIQUE (branch_id, employee_code)
);

CREATE TABLE technician_services (
    technician_id     UUID         NOT NULL,
    branch_service_id UUID         NOT NULL,
    skill_level       VARCHAR (30),
    created_at        TIMESTAMPTZ  DEFAULT NOW() NOT NULL,
    PRIMARY KEY (technician_id, branch_service_id),
    CONSTRAINT fk_technician_services_technician FOREIGN KEY (technician_id) REFERENCES technicians (id) ON DELETE CASCADE,
    CONSTRAINT fk_technician_services_branch_service FOREIGN KEY (branch_service_id) REFERENCES branch_services (id) ON DELETE CASCADE
);