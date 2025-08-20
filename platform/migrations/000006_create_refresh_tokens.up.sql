CREATE TABLE IF NOT EXISTS refresh_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_public_id UUID NOT NULL, -- untuk referensi publik
    token TEXT NOT NULL, -- refresh token yang di-hash (bukan raw token)
    device_name VARCHAR(100), -- nama device misalnya: "Chrome on Windows"
    ip_address VARCHAR(45),   -- IPv4/IPv6 address
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    FOREIGN KEY (user_public_id) REFERENCES users(public_id) ON DELETE CASCADE
);

CREATE INDEX idx_refresh_token_user ON refresh_tokens(user_public_id);
CREATE INDEX idx_refresh_token_token ON refresh_tokens(token);