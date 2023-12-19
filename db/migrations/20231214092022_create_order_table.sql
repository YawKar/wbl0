-- migrate:up

CREATE TABLE
    "order" (
        order_uid UUID PRIMARY KEY,
        track_number CHAR(20) NOT NULL,
        entry CHAR(10) NOT NULL,
        locale CHAR(10) NOT NULL,
        internal_signature CHAR(20),
        customer_id CHAR(20) NOT NULL,
        delivery_service CHAR(40) NOT NULL,
        shardkey CHAR(40) NOT NULL,
        sm_id BIGINT NOT NULL,
        date_created TIMESTAMP NOT NULL,
        oof_shard CHAR(40) NOT NULL
    );

-- migrate:down

DROP TABLE "order";
