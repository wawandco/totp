-- 20240607174243 - add_users migration
CREATE TABLE IF NOT EXISTS "users" (
    id varchar(255) PRIMARY KEY,
    email varchar(255) NOT NULL,
    secret varchar(255),
    password varchar(255),
    UNIQUE (email)
);

INSERT INTO
    users (id, email, password)
VALUES
    ('ef36fb13-f1ed-48d4-b1d2-d4e7ac5353c7', 'lgraham@gopher.co', '$2a$10$acsyUA.hm76sXreMbGYh4eUZ6vaO.3L7dPXqlROT/om.XWHfN3pLa'),
    ('13bc20ed-f527-43ec-bea1-b1ef89ae9f54', 'ehowell@gopher.co', '$2a$10$acsyUA.hm76sXreMbGYh4eUZ6vaO.3L7dPXqlROT/om.XWHfN3pLa'),
    ('43fd9f68-19be-43aa-9d5c-3448baa8fc59', 'sbauchy@gopher.co', '$2a$10$acsyUA.hm76sXreMbGYh4eUZ6vaO.3L7dPXqlROT/om.XWHfN3pLa'),
    ('4a07082b-678c-40de-9bda-6295a26dc990', 'oconnor@gopher.co', '$2a$10$acsyUA.hm76sXreMbGYh4eUZ6vaO.3L7dPXqlROT/om.XWHfN3pLa'),
    ('0c773217-9a7e-4c29-bd93-46d66e704d8f', 'mbrowns@gopher.co', '$2a$10$acsyUA.hm76sXreMbGYh4eUZ6vaO.3L7dPXqlROT/om.XWHfN3pLa');
