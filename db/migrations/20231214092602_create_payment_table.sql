-- migrate:up

CREATE TABLE
    payment (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        transaction UUID NOT NULL REFERENCES "order",
        request_id UUID,
        currency CHAR(20) NOT NULL,
        provider CHAR(50) NOT NULL,
        amount BIGINT NOT NULL,
        payment_dt BIGINT NOT NULL,
        bank CHAR(20),
        delivery_cost BIGINT NOT NULL,
        goods_total BIGINT NOT NULL,
        custom_fee BIGINT NOT NULL
    );

-- migrate:down

DROP TABLE payment;
