CREATE TABLE notifications (
    id         UUID          DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id    UUID         ,
    booking_id UUID         ,
    type       VARCHAR (50)  NOT NULL,
    channel    VARCHAR (20)  NOT NULL,
    title      VARCHAR (255) NOT NULL,
    content    TEXT          NOT NULL,
    status     VARCHAR (20)  DEFAULT 'PENDING' NOT NULL,
    sent_at    TIMESTAMPTZ  ,
    created_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ   DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_notifications_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_notifications_booking FOREIGN KEY (booking_id) REFERENCES bookings (id) ON DELETE CASCADE,
    CONSTRAINT chk_notifications_channel CHECK (channel IN ('EMAIL', 'SMS', 'PUSH', 'IN_APP')),
    CONSTRAINT chk_notifications_status CHECK (status IN ('PENDING', 'SENT', 'FAILED'))
);