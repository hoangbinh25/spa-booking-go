CREATE TABLE technician_time_offs (
    id            UUID        DEFAULT gen_random_uuid() PRIMARY KEY,
    technician_id UUID        NOT NULL,
    start_at      TIMESTAMPTZ NOT NULL,
    end_at        TIMESTAMPTZ NOT NULL,
    reason        TEXT       ,
    created_at    TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_technician_time_offs_technician FOREIGN KEY (technician_id) REFERENCES technicians (id) ON DELETE CASCADE,
    CONSTRAINT chk_technician_time_offs_range CHECK (start_at < end_at)
);