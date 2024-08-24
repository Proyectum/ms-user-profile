CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_profiles (
   id uuid primary key DEFAULT uuid_generate_v4(),
   first_name varchar(50) not null,
   last_name varchar(50) not null,
   username VARCHAR(50) not null,
   email VARCHAR(50) not null,
   bio VARCHAR(255),
   locale VARCHAR(5) not null,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP WITH TIME ZONE,
   unique (username, email)
);

CREATE TABLE notification_types (
    id uuid primary key DEFAULT uuid_generate_v4(),
    name varchar(15) not null unique,
    description varchar(255) not null
);

CREATE TABLE notification_settings (
   id uuid primary key DEFAULT uuid_generate_v4(),
   user_id uuid not null references user_profiles(id),
   notification_type_id uuid not null references notification_types(id),
   active boolean default false
);

INSERT INTO
    notification_types (name, description)
VALUES
    ('WEB', 'Web'),
    ('PUSH', 'Push Mobile'),
    ('EMAIL', 'Send email');