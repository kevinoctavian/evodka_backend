
CREATE TYPE school_type AS ENUM ('Sekolah', 'Kampus');

CREATE TABLE IF NOT EXISTS schools (
  id VARCHAR(35) PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  type school_type NOT NULL,
  npsn VARCHAR(50) NOT NULL,
  address VARCHAR(255) NOT NULL,
  logo TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TYPE user_role AS ENUM ('Admin', 'User');

CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(35) PRIMARY KEY,
  school_id VARCHAR(35) REFERENCES schools(id) ON DELETE SET NULL,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role user_role DEFAULT 'User',
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS candidates (
  id VARCHAR(35) PRIMARY KEY,
  school_id VARCHAR(35) REFERENCES schools(id) ON DELETE SET NULL,
  election_id VARCHAR(35) REFERENCES elections(id) ON DELETE SET NULL,
  ketua_name VARCHAR(100) NOT NULL,
  wakil_name VARCHAR(100) NOT NULL,
  photo_url VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS elections (
  id VARCHAR(35) PRIMARY KEY,
  school_id VARCHAR(35) REFERENCES schools(id) ON DELETE SET NULL,
  start_at TIMESTAMP NOT NULL,
  end_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS votes (
  id VARCHAR(35) PRIMARY KEY,
  user_id VARCHAR(35) REFERENCES users(id) ON DELETE CASCADE,
  candidate_id VARCHAR(35) REFERENCES candidates(id) ON DELETE CASCADE,
  school_id VARCHAR(35) REFERENCES schools(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(user_id, school_id)
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
  id VARCHAR(35) PRIMARY KEY,
  user_id VARCHAR(35) REFERENCES users(id) ON DELETE CASCADE,
  token_hash VARCHAR(64),
  device_name VARCHAR(100),
  ip_address VARCHAR(100),
  user_agent TEXT,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  revoked BOOLEAN DEFAULT FALSE
);