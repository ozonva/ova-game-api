-- +goose Up
INSERT INTO "users" (id, email, name, created_at) VALUES (1, 'mail@app.testfiles', 'John', NOW());

-- +goose Down
DELETE FROM "users" WHERE id = 1;