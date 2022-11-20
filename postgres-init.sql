CREATE TABLE IF NOT EXISTS Deposit(
    owner_id BIGINT PRIMARY KEY,
    balance BIGINT,

    CONSTRAINT chk_balance_not_negative
    CHECK(balance >= 0)
);

CREATE TABLE IF NOT EXISTS Transaction(
    id bigserial PRIMARY KEY,
    sender_id BIGINT NULL,
    recipient_id BIGINT NULL,
    amount BIGINT NOT NULL,
    description VARCHAR(100) NULL,
    transaction_date TIMESTAMP NOT NULL,

    CONSTRAINT chk_amount_not_negative
    CHECK(amount > 0)
);