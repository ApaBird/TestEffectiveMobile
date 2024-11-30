CREATE TABLE IF NOT EXISTS songs (
    id VARCHAR DEFAULT gen_random_uuid () PRIMARY KEY,
    songName VARCHAR ,
    group_id VARCHAR ,
    "text" VARCHAR ,
    link VARCHAR ,
    "date" VARCHAR ,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS groups (
    id VARCHAR DEFAULT gen_random_uuid () PRIMARY KEY,
    "name" VARCHAR
);