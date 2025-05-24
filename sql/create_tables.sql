-- UUID extension'ı aktif et
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ROLES tablosu
CREATE TABLE IF NOT EXISTS roles (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_name VARCHAR(45) NOT NULL
);

-- USERS tablosu
CREATE TABLE IF NOT EXISTS users (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(45) NOT NULL,
    surname VARCHAR(45) NOT NULL,
    username VARCHAR(45) NOT NULL,
    email VARCHAR(45) NOT NULL,
    roleID UUID REFERENCES roles(ID)
);