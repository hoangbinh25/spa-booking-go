CREATE TABLE payments (
    id                      UUID            DEFAULT gen_random_uuid() PRIMARY KEY,
    booking_id              UUID            NOT NULL,
    amount                  NUMERIC (14, 2) NOT NULL,
    payment_method          VARCHAR (30)    NOT NULL,
    status                  VARCHAR (30)    DEFAULT 'PENDING' NOT NULL,
    provider                VARCHAR (50)   ,
    provider_transaction_id VARCHAR (150)  ,
    paid_at                 TIMESTAMPTZ    ,
    created_at              TIMESTAMPTZ     DEFAULT NOW() NOT NULL,
    updated_at              TIMESTAMPTZ     DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_payments_booking FOREIGN KEY (booking_id) REFERENCES bookings (id),
    CONSTRAINT chk_payments_amount CHECK (amount > 0),
    CONSTRAINT chk_payments_method CHECK (payment_method IN ('CASH', 'BANK_TRANSFER', 'ONLINE_PAYMENT')),
    CONSTRAINT chk_payments_status CHECK (status IN ('PENDING', 'PAID', 'FAILED', 'REFUNDED', 'PARTIALLY_REFUNDED'))
);