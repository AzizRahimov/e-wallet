-- созданиет юзера
CREATE TABLE users
(
    id    bigserial primary key,
    fio   varchar not null,
    age   integer,
    phone varchar not null unique,
    pin   varchar not null

);





-- создание кошелька
create table wallet
(
    id            bigserial,
    user_id       int
        constraint wallet_users_id_fk
            references users,
    balance       float4,
    account       varchar not null,
    is_identified bool      default false,
    created_at    timestamp default current_timestamp,
    updated_at    timestamp default current_timestamp,
    deleted_at    timestamp default current_timestamp,
    is_removed    bool      default false
);

create unique index wallet_account_uindex
    on wallet (account);

create unique index wallet_id_uindex
    on wallet (id);

alter table wallet
    add constraint wallet_pk
        primary key (id);



SELECT id, user_id, balance
FROM wallet;


-- добавление тестовых данных
INSERT INTO wallet(user_id, balance, account)
VALUES (3, 100000, 3213);



-- создание таблицы для транзакции
CREATE TABLE transactions
(
    id         bigserial primary key,
    "from"     varchar not null,
    "to"       varchar not null,
    status     varchar not null,

    amount     float   not null,
    trn_type   varchar not null default '',
    created_at timestamp        default current_timestamp
);




