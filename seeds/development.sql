-- Development-only reference data. Every statement is idempotent so it can run
-- again after `docker compose up` without creating duplicate records.
INSERT INTO spas (id, name, phone, email)
VALUES ('00000000-0000-0000-0000-000000000001', 'Serenity Spa', '028-5555-0100', 'hello@serenity.local')
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name, phone = EXCLUDED.phone, email = EXCLUDED.email, updated_at = NOW();

INSERT INTO branches (id, spa_id, code, name, phone, email, address, timezone)
VALUES (
    '00000000-0000-0000-0000-000000000002',
    '00000000-0000-0000-0000-000000000001',
    'D1',
    'Serenity Spa District 1',
    '028-5555-0101',
    'd1@serenity.local',
    '1 Nguyen Hue, District 1, Ho Chi Minh City',
    'Asia/Ho_Chi_Minh'
)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name, phone = EXCLUDED.phone, email = EXCLUDED.email,
    address = EXCLUDED.address, timezone = EXCLUDED.timezone, updated_at = NOW();

INSERT INTO users (id, email, full_name)
VALUES ('00000000-0000-0000-0000-000000000003', 'linh.nguyen@serenity.local', 'Linh Nguyen')
ON CONFLICT (id) DO UPDATE
SET email = EXCLUDED.email, full_name = EXCLUDED.full_name, updated_at = NOW();

INSERT INTO customers (id, full_name, phone, email)
VALUES ('00000000-0000-0000-0000-000000000004', 'Minh Tran', '0900000001', 'minh.tran@example.com')
ON CONFLICT (id) DO UPDATE
SET full_name = EXCLUDED.full_name, phone = EXCLUDED.phone, email = EXCLUDED.email, updated_at = NOW();

INSERT INTO services (id, code, name, description, default_duration_minutes)
VALUES
    ('00000000-0000-0000-0000-000000000005', 'BODY-60', 'Body Massage 60 Minutes', 'Relaxing full body massage.', 60),
    ('00000000-0000-0000-0000-000000000006', 'FACIAL-45', 'Facial Care 45 Minutes', 'Basic facial treatment.', 45)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name, description = EXCLUDED.description,
    default_duration_minutes = EXCLUDED.default_duration_minutes, updated_at = NOW();

INSERT INTO branch_services (id, branch_id, service_id, price, duration_minutes)
VALUES
    ('00000000-0000-0000-0000-000000000007', '00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000005', 500000, 60),
    ('00000000-0000-0000-0000-000000000008', '00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000006', 350000, 45)
ON CONFLICT (id) DO UPDATE
SET price = EXCLUDED.price, duration_minutes = EXCLUDED.duration_minutes, updated_at = NOW();

INSERT INTO technicians (id, user_id, branch_id, employee_code, display_name, bio)
VALUES (
    '00000000-0000-0000-0000-000000000009',
    '00000000-0000-0000-0000-000000000003',
    '00000000-0000-0000-0000-000000000002',
    'TECH-001',
    'Linh Nguyen',
    'Massage and facial care specialist.'
)
ON CONFLICT (id) DO UPDATE
SET display_name = EXCLUDED.display_name, bio = EXCLUDED.bio, updated_at = NOW();

INSERT INTO technician_services (technician_id, branch_service_id, skill_level)
VALUES
    ('00000000-0000-0000-0000-000000000009', '00000000-0000-0000-0000-000000000007', 'SENIOR'),
    ('00000000-0000-0000-0000-000000000009', '00000000-0000-0000-0000-000000000008', 'SENIOR')
ON CONFLICT (technician_id, branch_service_id) DO UPDATE
SET skill_level = EXCLUDED.skill_level;

INSERT INTO technician_schedules (id, technician_id, day_of_week, start_time, end_time)
VALUES (
    '00000000-0000-0000-0000-000000000010',
    '00000000-0000-0000-0000-000000000009',
    1,
    '09:00',
    '18:00'
)
ON CONFLICT (id) DO UPDATE
SET start_time = EXCLUDED.start_time, end_time = EXCLUDED.end_time, updated_at = NOW();
