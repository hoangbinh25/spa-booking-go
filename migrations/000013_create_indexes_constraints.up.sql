CREATE INDEX idx_bookings_branch_start
    ON bookings(branch_id, start_at);

CREATE INDEX idx_bookings_customer_created
    ON bookings(customer_id, created_at DESC);

CREATE INDEX idx_booking_items_technician_time
    ON booking_items(technician_id, start_at, end_at);

CREATE INDEX idx_technician_schedules_lookup
    ON technician_schedules(technician_id, day_of_week);

CREATE INDEX idx_technician_time_offs_lookup
    ON technician_time_offs(technician_id, start_at, end_at);

CREATE INDEX idx_payments_booking
    ON payments(booking_id);

CREATE INDEX idx_notifications_user
    ON notifications(user_id);

CREATE INDEX idx_notifications_booking
    ON notifications(booking_id);