INSERT INTO server (summary, host, port, status, created_at, deleted_at) VALUES ('local', '127.0.0.1', 9000, 1, extract(epoch from now()), null);