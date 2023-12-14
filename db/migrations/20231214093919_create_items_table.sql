-- migrate:up

CREATE TABLE
    items (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        order_uid UUID NOT NULL REFERENCES "order",
        chrt_id BIGINT NOT NULL,
        track_number CHAR(20) NOT NULL,
        price BIGINT NOT NULL,
        rid UUID NOT NULL,
        name CHAR(50) NOT NULL,
        sale BIGINT NOT NULL,
        size CHAR(50) NOT NULL,
        total_price BIGINT NOT NULL,
        nm_id BIGINT NOT NULL,
        brand CHAR(50) NOT NULL,
        status INT NOT NULL
    );

-- migrate:down

DROP TABLE items;
