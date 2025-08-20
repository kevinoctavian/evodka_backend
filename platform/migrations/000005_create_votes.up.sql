CREATE TABLE IF NOT EXISTS votes (
  id SERIAL PRIMARY KEY,
  public_id UUID NOT NULL DEFAULT gen_random_uuid(),
  voting_id INT NOT NULL REFERENCES votings(id) ON DELETE CASCADE,
  candidate_id INT NOT NULL REFERENCES candidates(id) ON DELETE CASCADE,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL,
  UNIQUE(voting_id, user_id, public_id) -- 1 user hanya bisa voting sekali per voting
);
