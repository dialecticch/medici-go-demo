CREATE TABLE tokens
(
    id       SERIAL PRIMARY KEY,
    address  VARCHAR(42) NOT NULL,
    chain_id INT         NOT NULL,
    name     TEXT        NOT NULL,
    symbol   TEXT        NOT NULL,
    decimals SMALLINT    NOT NULL DEFAULT 18,
    UNIQUE (chain_id, address)
);

CREATE TABLE strategies
(
    id       SERIAL PRIMARY KEY,
    chain_id INT  NOT NULL,
    name     TEXT NOT NULL,
    UNIQUE (chain_id, name)
);

CREATE COLLATION en_natural (LOCALE = 'en-US-u-kn-true', PROVIDER = 'icu');

CREATE TABLE versions
(
    id          SERIAL PRIMARY KEY,
    strategy_id INT                            NOT NULL,
    address     VARCHAR(42)                    NOT NULL,
    version     VARCHAR(11) COLLATE en_natural NOT NULL,
    FOREIGN KEY (strategy_id) REFERENCES strategies (id),
    UNIQUE (address, version)
);

CREATE VIEW latest_versions AS SELECT DISTINCT ON (strategy_id) * FROM versions ORDER BY strategy_id, version DESC, id DESC;

CREATE TABLE safes
(
    id       SERIAL PRIMARY KEY,
    address  VARCHAR(42) NOT NULL,
    chain_id INT         NOT NULL,
    name     TEXT        NOT NULL,
    UNIQUE (chain_id, address)
);

CREATE TABLE pools
(
    id          SERIAL PRIMARY KEY,
    pool        NUMERIC(78, 0) NOT NULL,
    name        TEXT           NOT NULL,
    strategy_id INT            NOT NULL,
    token_id    INT            NOT NULL,
    FOREIGN KEY (strategy_id) REFERENCES strategies (id),
    FOREIGN KEY (token_id) REFERENCES tokens (id),
    UNIQUE (pool, strategy_id)
);

CREATE TYPE harvest_action AS ENUM ('harvest', 'compound');
CREATE TYPE threshold_type AS ENUM ('gas_percentage', 'elapsed_blocks');

CREATE TABLE safe_pools
(
    id                           SERIAL PRIMARY KEY,
    safe_id                      INT            NOT NULL,
    pool_id                      INT            NOT NULL,
    threshold                    FLOAT          NOT NULL,
    action                       harvest_action NOT NULL DEFAULT 'harvest',
    threshold_type               threshold_type NOT NULL DEFAULT 'gas_percentage',
    last_harvested               INT            NOT NULL DEFAULT 0,
    active                       BOOLEAN                 DEFAULT TRUE,
    deposited_amount             NUMERIC(78, 0)          DEFAULT 0,
    last_deposited_amount_update INT            NOT NULL DEFAULT 0,
    FOREIGN KEY (safe_id) REFERENCES safes (id),
    FOREIGN KEY (pool_id) REFERENCES pools (id),
    UNIQUE (safe_id, pool_id)
);

CREATE TYPE transaction_type AS ENUM ('deposit', 'withdraw');

CREATE TABLE transactions
(
    timestamp        TIMESTAMPTZ      NOT NULL,
    transaction_type transaction_type NOT NULL,
    tx               VARCHAR(66)      NOT NULL,
    log_index        INT              NOT NULL,
    amount           NUMERIC(78, 0)   NOT NULL,
    pool_safe_id     INT              NOT NULL,
    FOREIGN KEY (pool_safe_id) REFERENCES safe_pools (id),
    UNIQUE (tx, log_index)
);

CREATE TABLE harvests
(
    harvested_at TIMESTAMPTZ    NOT NULL,
    tx           VARCHAR(66)    NOT NULL,
    log_index    INT            NOT NULL,
    amount       NUMERIC(78, 0) NOT NULL,
    pool_safe_id INT            NOT NULL,
    FOREIGN KEY (pool_safe_id) REFERENCES safe_pools (id),
    UNIQUE (tx, log_index)
);
