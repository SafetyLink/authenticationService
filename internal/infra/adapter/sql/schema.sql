CREATE TABLE users
(
    id         BIGSERIAL,
    username   VARCHAR(255) UNIQUE NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    avatar_id  BIGINT                   default 110308276497481728,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)

);

CREATE TABLE privacy
(
    user_id       BIGINT NOT NULL,
    is_private    BOOLEAN DEFAULT FALSE,
    is_searchable BOOLEAN DEFAULT TRUE,

    FOREIGN KEY (user_id) REFERENCES Users (id)

);

CREATE TABLE user_security
(
    user_id    BIGINT                          NOT NULL,
    password   VARCHAR(255)                    NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    device_id  BIGSERIAL REFERENCES Users (id) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES Users (id)

);

CREATE TABLE session
(
    id         BIGSERIAL,
    user_id    BIGINT       NOT NULL,
    device_id  VARCHAR(255) NOT NULL,
    device_os  VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);



CREATE TABLE friend_list
(
    id         BIGSERIAL,
    user_id    BIGINT NOT NULL,
    friend_id  BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (friend_id) REFERENCES Users (id),
    UNIQUE (user_id, friend_id)
);
