 CREATE TABLE users (
    id TEXT PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    last_seen timestamp without time zone
);

CREATE TABLE hobbies (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE users_hobbies (
  user_id TEXT  REFERENCES users(id) ON DELETE CASCADE,
  hobby_id uuid REFERENCES hobbies(id) ON DELETE CASCADE ,
  PRIMARY KEY (user_id, hobby_id)
);

CREATE TABLE chats (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id_a TEXT NOT NULL,
  user_id_b TEXT NOT NULL
);

CREATE TABLE messages (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  chat_id uuid REFERENCES chats(id) ON DELETE CASCADE,
  sender_id TEXT NOT NULL,
  content TEXT NOT NULL,
  sent_at timestamp without time zone
);



CREATE INDEX idx_user_id ON users_hobbies (user_id);
CREATE INDEX idx_hobby_id ON users_hobbies (hobby_id);