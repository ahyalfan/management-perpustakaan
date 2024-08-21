-- Active: 1723868453146@@127.0.0.1@5432@management_perpustakaan
CREATE DATABASE management_perpustakaan;

CREATE TABLE public.customers (
    id character varying(36) DEFAULT gen_random_uuid () NOT NULL PRIMARY KEY,
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone
);

SELECT * FROM customers

CREATE TABLE public.users (
    id character varying(36) DEFAULT gen_random_uuid () NOT NULL PRIMARY KEY,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);

INSERT INTO
    public.users (id, email, password)
VALUES (
        '4246fb58-ff45-4d2a-8946-93e541fc39fd',
        'admin@yogi.id',
        '$2a$12$Rvslxj25D4OU7w3Ercz/IucMiDkEp1dOCSwq902oWpy0mqcUx2GAq'
    );

CREATE TABLE public.book_stocks (
    id SERIAL PRIMARY key,
    book_id character varying(36) NOT NULL,
    code character varying(50) NOT NULL,
    status character varying(50) NOT NULL,
    borrower_id character varying(36),
    borrowed_at timestamp(6) without time zone
);

CREATE TABLE public.books (
    id character varying(36) DEFAULT gen_random_uuid () NOT NULL PRIMARY KEY,
    title character varying(255) NOT NULL,
    description text,
    isbn character varying(100) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    -- cover_id character varying(36)
);

CREATE TABLE public.journals (
    id character varying(36) DEFAULT gen_random_uuid () NOT NULL,
    book_id character varying(36) NOT NULL,
    stock_code character varying(255) NOT NULL,
    customer_id character varying(36) NOT NULL,
    status character varying(50) NOT NULL,
    borrowed_at timestamp(6) without time zone NOT NULL,
    returned_at timestamp(6) without time zone,
    -- due_at timestamp(6) without time zone
);

CREATE TABLE public.media (
    id character varying(36) DEFAULT gen_random_uuid () NOT NULL,
    path text,
    created_at timestamp(6) without time zone NOT NULL
);

ALTER TABLE books ADD COLUMN cover_id CHARACTER varying(100);

ALTER TABLE journals
ADD COLUMN due_at timestamp(6) without time zone;

CREATE TABLE public.charges (
    id character varying(36) NOT NULL,
    journal_id character varying(36) NOT NULL,
    days_late integer DEFAULT 1 NOT NULL,
    daily_late_fee integer NOT NULL,
    total integer NOT NULL,
    user_id character varying(36) NOT NULL,
    created_at timestamp(6) without time zone
);