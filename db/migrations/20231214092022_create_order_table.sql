-- migrate:up

CREATE TABLE
    "order" (
        order_uid UUID PRIMARY KEY,
        track_number CHAR(75) NOT NULL,
        entry CHAR(75) NOT NULL,
        locale CHAR(75) NOT NULL,
        internal_signature CHAR(75),
        customer_id CHAR(75) NOT NULL,
        delivery_service CHAR(75) NOT NULL,
        shardkey CHAR(75) NOT NULL,
        sm_id BIGINT NOT NULL,
        date_created TIMESTAMP NOT NULL,
        oof_shard CHAR(75) NOT NULL
    );

-- migrate:down

DROP TABLE "order";
