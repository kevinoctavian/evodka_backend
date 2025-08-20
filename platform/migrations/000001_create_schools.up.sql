CREATE TABLE IF NOT EXISTS schools (
  id SERIAL PRIMARY KEY,
  public_id UUID NOT NULL DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  type VARCHAR(10)  NOT NULL, -- sekolah / kampus
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL,
  UNIQUE (public_id, name),
  CHECK (type IN ('sekolah', 'kampus'))
);
