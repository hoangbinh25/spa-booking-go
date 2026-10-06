CREATE TABLE technician_schedules (
    id            UUID        DEFAULT gen_random_uuid() PRIMARY KEY,
    technician_id UUID        NOT NULL,
    day_of_week   SMALLINT    NOT NULL,
    start_time    TIME        NOT NULL,
    end_time      TIME        NOT NULL,
    is_active     BOOLEAN     DEFAULT TRUE NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_technician_schedules_technician FOREIGN KEY (technician_id) REFERENCES technicians (id) ON DELETE CASCADE,
    CONSTRAINT chk_technician_schedules_day CHECK (day_of_week BETWEEN 1 AND 7),
    CONSTRAINT chk_technician_schedules_time CHECK (start_time < end_time)
);