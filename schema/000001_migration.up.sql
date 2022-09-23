CREATE TABLE genres
(
    id   smallint               NOT NULL,
    name character varying(255) NOT NULL,
    CONSTRAINT genres_pkey PRIMARY KEY (id)
);


CREATE TABLE authors
(
    id          smallserial       NOT NULL,
    name        character varying NOT NULL,
    surname     character varying NOT NULL,
    description text,
    CONSTRAINT authors_pkey PRIMARY KEY (id)
);


CREATE TABLE books
(
    id          smallserial            NOT NULL,
    name        character varying(255) NOT NULL,
    description text,
    genre_id    smallint               NOT NULL,
    author_id   smallint               NOT NULL,
    CONSTRAINT books_pkey PRIMARY KEY (id),
    CONSTRAINT author_id_fk FOREIGN KEY (author_id) REFERENCES authors (id) MATCH SIMPLE,
    CONSTRAINT genre_id_fk FOREIGN KEY (genre_id) REFERENCES genres (id) MATCH SIMPLE
);


CREATE TABLE users
(
    id       smallserial            NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_username_key UNIQUE (username)
);


CREATE TABLE lists
(
    id      smallserial            NOT NULL,
    title   character varying(255) NOT NULL,
    user_id integer                NOT NULL,
    CONSTRAINT lists_pkey PRIMARY KEY (id),
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) MATCH SIMPLE ON DELETE CASCADE
);


CREATE TABLE list_items
(
    list_id integer                NOT NULL,
    book_id integer                NOT NULL,
    status  character varying(255) NOT NULL,
    CONSTRAINT list_items_pkey PRIMARY KEY (list_id, book_id),
    CONSTRAINT book_id_fk FOREIGN KEY (book_id) REFERENCES books (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT list_id_fk FOREIGN KEY (list_id) REFERENCES lists (id) MATCH SIMPLE ON DELETE CASCADE,
    CONSTRAINT status_constraint CHECK (status::text = 'done'::text OR
                                        status::text = 'in process'::text OR
                                        status::text = 'in plans'::text)
);
