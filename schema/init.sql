CREATE TABLE IF NOT EXISTS invoices
(
    id           serial       not null unique,
    title        varchar(255) not null,
    description  varchar(255) not null,
    company_name varchar(255) not null,
    date         timestamp    not null default now(),
    total_cost   decimal(10, 2)
);