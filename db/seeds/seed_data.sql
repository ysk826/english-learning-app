-- Insert test user
-- Note: password_hash is for 'password123' - this is just for development!
INSERT INTO users (username, email, password_hash) 
VALUES ('testuser', 'test@example.com', '$2a$10$1JE4dM3pKr6VbMO.OGBjxemfuaFpeDiD3KxAxjTQ4irsNkbQF8TN6')
ON CONFLICT (email) DO NOTHING;

-- Insert sample words
INSERT INTO words (user_id, word, meaning, example)
VALUES 
    (1, 'ubiquitous', 'present, appearing, or found everywhere', 'Mobile phones are now ubiquitous in modern society.'),
    (1, 'ephemeral', 'lasting for a very short time', 'The beauty of cherry blossoms is ephemeral.'),
    (1, 'pragmatic', 'dealing with things sensibly and realistically', 'We need a pragmatic approach to solving this problem.'),
    (1, 'eloquent', 'fluent or persuasive in speaking or writing', 'She gave an eloquent speech that moved the audience.')
ON CONFLICT DO NOTHING;

-- Insert sample sentences
INSERT INTO sentences (user_id, content)
VALUES 
    (1, 'The ubiquitous nature of technology has made our lives more convenient but also more complex.'),
    (1, 'His eloquent explanation of the pragmatic solution impressed everyone in the meeting.')
ON CONFLICT DO NOTHING;