CREATE TABLE IF NOT EXISTS candidates (
  id SERIAL PRIMARY KEY,
  public_id UUID NOT NULL DEFAULT gen_random_uuid(),
  voting_id INT NOT NULL REFERENCES votings(id) ON DELETE CASCADE,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  vision TEXT,
  mission TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL,
  CONSTRAINT unique_public_id UNIQUE (public_id)
);
