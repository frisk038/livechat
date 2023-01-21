CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE hobbies (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE users_hobbies (
  user_id uuid  REFERENCES users(id) ON DELETE CASCADE,
  hobby_id uuid REFERENCES hobbies(id) ON DELETE CASCADE ,
  PRIMARY KEY (user_id, hobby_id)
);


CREATE INDEX idx_user_id ON users_hobbies (user_id);
CREATE INDEX idx_hobby_id ON users_hobbies (hobby_id);