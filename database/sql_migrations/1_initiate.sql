-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id BIGINT NOT NULL,
    name VARCHAR(256) DEFAULT 'nameFromEmail',
    email VARCHAR(256) NOT NULL,
    username VARCHAR(256) DEFAULT 'usernameRandom',
    password VARCHAR(256) NOT NULL
);

CREATE TABLE articles (
    id BIGINT NOT NULL,
    id_user BIGINT NOT NULL,
    title VARCHAR(256),
    image VARCHAR(256),
    content VARCHAR(256)
);

CREATE TABLE likes (
    id BIGINT NOT NULL,
    id_user BIGINT NOT NULL,
    id_article BIGINT NOT NULL,
    respon VARCHAR(256)
);

CREATE TABLE comments (
    id BIGINT NOT NULL,
    id_user BIGINT NOT NULL,
    id_article BIGINT NOT NULL,
    id_reply_comment BIGINT NOT NULL,
    comment VARCHAR(256),
    image VARCHAR(256)
);

CREATE TABLE comments_likes (
    id BIGINT NOT NULL,
    id_user BIGINT NOT NULL,
    id_comment BIGINT NOT NULL,
    respon VARCHAR(256)
);

-- +migrate StatementEnd