CREATE TABLE users
(
    id            serial       primary key,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(5) check (role in ('admin', 'user')) default 'user'
);

CREATE TABLE org_unit
(
    id                serial       primary key,
    name              varchar(255) not null unique,
    description       text
);

CREATE TABLE staff
(
    id                serial       primary key,
    fio               varchar(255) not null,
    birth_year        int          not null,
    employment_date   timestamp ,
    snils             varchar(14)  not null unique
);

CREATE TABLE post_rate
(
    id                serial       primary key,
    staffer_id        int,
    org_unit_id       int          references org_unit on update cascade on delete cascade,
    post              varchar(255) not null,
    rate              real         not null
);

CREATE TABLE product
(
    id          serial       primary key,
    name        varchar(255) not null,
    cost        int,
    count       int,
    description text
);

CREATE TABLE order_table
(
    id      serial       primary key,
    name    varchar(255) not null,
    user_id int          references users on update cascade on delete cascade,
    status  varchar(12)  check (status in ('IS_PAID',
                                          'IS_NOT_PAID')) default 'IS_NOT_PAID'
);

CREATE TABLE order_products_table
(
    order_id       int    references order_table on update cascade on delete cascade,
    product_id     int    references product on update cascade on delete cascade,
    product_count  int
);