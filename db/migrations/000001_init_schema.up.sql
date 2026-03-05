--  Create type data ENUM
CREATE TYPE user_role AS ENUM ('admin', 'relawan');
CREATE TYPE request_status AS ENUM ('pending', 'processing', 'fulfilled');
CREATE TYPE urgency_level AS ENUM ('rendah', 'sedang', 'tinggi', 'kritis');

-- Table Master User
CREATE TABLE users (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    role user_role DEFAULT 'relawan',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table Master Location (Posko)
CREATE TABLE posko (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    address TEXT,
    latitude DECIMAL,
    longitude DECIMAL,
    coordinator_id VARCHAR REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table Master Item (Logistic)
CREATE TABLE items (
    id VARCHAR PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    unit VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table Inventory Stock per Posko
CREATE TABLE posko_inventory (
    id VARCHAR PRIMARY KEY,
    posko_id VARCHAR REFERENCES posko(id),
    item_id VARCHAR REFERENCES items(id),
    quantity INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table Header Report/Request (Input AI)
CREATE TABLE logistics_requests (
    id VARCHAR PRIMARY KEY,
    posko_id VARCHAR REFERENCES posko(id),
    requested_by VARCHAR REFERENCES users(id),
    original_prompt TEXT,
    status request_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table Detail Report/Request
CREATE TABLE request_items (
    id VARCHAR PRIMARY KEY,
    request_id VARCHAR REFERENCES logistics_requests(id),
    item_id VARCHAR REFERENCES items(id),
    quantity INT NOT NULL,
    urgency urgency_level DEFAULT 'sedang',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);