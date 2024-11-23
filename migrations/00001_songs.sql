-- +goose Up
-- +goose StatementBegin
CREATE TABLE songs (
    id uuid PRIMARY KEY,
    song_name TEXT NOT NULL,
    group_name TEXT NOT NULL,
    release_date TIMESTAMPTZ NOT NULL,
    lyrics TEXT NOT NULL,
    link TEXT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE songs;

-- +goose StatementEnd