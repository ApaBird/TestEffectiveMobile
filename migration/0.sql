CREATE TABLE IF NOT EXISTS songs (
    id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY,
    songName VARCHAR ,
    group_id VARCHAR ,
    "text" VARCHAR ,
    link VARCHAR ,
    "date" VARCHAR ,
    created_at TIMESTAMP DEFAULT now(),
);

CREATE TABLE IF NOT EXISTS groups (
    id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" VARCHAR
);