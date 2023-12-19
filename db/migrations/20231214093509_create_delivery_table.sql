-- migrate:up

CREATE TABLE
    delivery (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        order_uid UUID NOT NULL REFERENCES "order",
        name CHAR(50) NOT NULL,
        phone CHAR(20) NOT NULL,
        zip CHAR(10) NOT NULL,
        city CHAR(75) NOT NULL,
        address CHAR(150) NOT NULL,
        region CHAR(100) NOT NULL,
        email CHAR(100) NOT NULL
    );

-- migrate:down

DROP TABLE delivery;
