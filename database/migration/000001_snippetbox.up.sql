-- Create a `snippets` table.
CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    expires TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...
A frog jumps into the pond,
splash! Silence again.

- Matsuo Bashō',
    now() AT TIME ZONE 'UTC',
    now() AT TIME ZONE 'UTC' + INTERVAL '1 year'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry
forest, winds howl in rage
with no leaves to blow.

- Natsume Soseki',
    now() AT TIME ZONE 'UTC',
    now() AT TIME ZONE 'UTC' + INTERVAL '1 year'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning
the mirror I stare into
shows my father''s face.

- Murakami Kijo',
    now() AT TIME ZONE 'UTC',
    now() AT TIME ZONE 'UTC' + INTERVAL '1 year'
);
