CREATE TABLE users (
    id UUID PRIMARY KEY,

    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,

    role VARCHAR(255) NOT NULL DEFAULT 'user',
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT user_role_chk CHECK (role IN ('admin', 'moderator', 'user'))
);

CREATE TYPE patent_type_enum AS ENUM (
    'invention',
    'utility_model',
    'industrail_design'
);

CREATE TABLE ip_objects (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,

    jurisdiction VARCHAR(10) NOT NULL,
    patent_type patent_type_enum NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);