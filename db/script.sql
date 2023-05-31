create database extended_todo;

\c extended_todo;

create table users (
                       id serial primary key,
                       name varchar(255),
                       email varchar(255),
                       password varchar(255)
);

create table token (
                       user_id integer references users(id) on DELETE CASCADE,
                       refresh_token varchar(255)
);


create table card (
                      id serial primary key,
                      name varchar(255),
                      admin_id integer REFERENCES users(id) on DELETE CASCADE
);

create table task (
                      id serial primary key,
                      card_id integer REFERENCES card(id) on DELETE CASCADE ,
                      title varchar(255),
                      description varchar(255),
                      start timestamptz null,
                      percent smallint,
                      deadline timestamptz null,
                      completed boolean
);



-- Заполнение таблицы users
INSERT INTO users (name, email, password)
VALUES
    ('John Doe', 'john@example.com', 'password123'),
    ('Jane Smith', 'jane@example.com', 'password456'),
    ('Michael Johnson', 'michael@example.com', 'password789'),
    ('Emily Davis', 'emily@example.com', 'password123'),
    ('Daniel Wilson', 'daniel@example.com', 'password456'),
    ('Olivia Brown', 'olivia@example.com', 'password789'),
    ('William Taylor', 'william@example.com', 'password123'),
    ('Sophia Anderson', 'sophia@example.com', 'password456'),
    ('James Martinez', 'james@example.com', 'password789'),
    ('Ava Harris', 'ava@example.com', 'password123');

-- Заполнение таблицы token
INSERT INTO token (user_id, refresh_token)
VALUES
    (1, 'token123'),
    (2, 'token456'),
    (3, 'token789'),
    (4, 'token123'),
    (5, 'token456'),
    (6, 'token789'),
    (7, 'token123'),
    (8, 'token456'),
    (9, 'token789'),
    (10, 'token123');

-- Заполнение таблицы card
INSERT INTO card (name, admin_id)
VALUES
    ('Card 1', 1),
    ('Card 2', 2),
    ('Card 3', 3),
    ('Card 4', 4),
    ('Card 5', 5),
    ('Card 6', 6),
    ('Card 7', 7),
    ('Card 8', 8),
    ('Card 9', 9),
    ('Card 10', 10);

-- Заполнение таблицы task
INSERT INTO task (card_id, title, description, start, percent, deadline, completed)
VALUES
    (1, 'Task 1', 'Description 1', '2023-05-01', 50, '2023-05-10', true),
    (1, 'Task 2', 'Description 2', '2023-05-02', 75, '2023-05-15', false),
    (2, 'Task 3', 'Description 3', '2023-05-03', 25, '2023-05-12', false),
    (2, 'Task 4', 'Description 4', '2023-05-04', 0, '2023-05-20', false),
    (3, 'Task 5', 'Description 5', '2023-05-05', 100, '2023-05-08', true),
    (3, 'Task 6', 'Description 6', '2023-05-06', 50, '2023-05-18', false),
    (4, 'Task 7', 'Description 7', '2023-05-07', 75, '2023-05-14', false),
    (4, 'Task 8', 'Description 8', '2023-05-08', 25, '2023-05-25', false),
    (5, 'Task 9', 'Description 9', '2023-05-09', 0, '2023-05-10', true),
    (5, 'Task 10', 'Description 10', '2023-05-10', 100, '2023-05-30', false);