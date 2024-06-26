-- +goose Up

CREATE TABLE users (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  is_organization bool NOT NULL,
  name varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  about varchar,
  avatar_url varchar,
  images varchar,
  auth_provider varchar,
  auth_hash varchar,
  registered_at timestamp NOT NULL,
  last_visited_at timestamp NOT NULL,
  is_online bool NOT NULL,
  location_lat float,
  location_long float
);

CREATE TABLE donations (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  created_by_user_id int NOT NULL,
  created_at timestamp NOT NULL,
  starts_at timestamp,
  ends_at timestamp,
  description varchar NOT NULL,
  images varchar,
  servings_total int,
  servings_left int,
  location_lat float NOT NULL,
  location_long float NOT NULL
);

CREATE TABLE donation_requests (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  donation_id int,
  requester_id int,
  comment varchar,
  images varchar,
  created_at timestamp,
  accepted_at timestamp,
  accepted_by_user_id int,
  confirmed_at timestamp,
  confirmed_by_user_id int
);

CREATE TABLE donation_events (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  donation_id int NOT NULL,
  created_by_user_id int NOT NULL,
  created_at timestamp NOT NULL,
  comment varchar NOT NULL,
  images varchar,
  reply_to_user_id int
);

CREATE TABLE ratings (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  created_at timestamp NOT NULL,
  from_user_id int NOT NULL,
  to_user_id int NOT NULL,
  rating int NOT NULL,
  comment varchar,
  images varchar,
  location_lat float NOT NULL,
  location_long float NOT NULL
);

CREATE TABLE notifications (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  created_at timestamp NOT NULL,
  donation_id int,
  donation_event_id int,
  donation_request_id int,
  rating_id int,
  user_id int,
  comment varchar
);

CREATE TABLE notification_history (
  id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  notification_id int NOT NULL,
  user_id int NOT NULL,
  sent_at timestamp NOT NULL,
  read_at timestamp
);

ALTER TABLE donations ADD FOREIGN KEY (created_by_user_id) REFERENCES users (id);

ALTER TABLE donation_requests ADD FOREIGN KEY (donation_id) REFERENCES donations (id);

ALTER TABLE donation_requests ADD FOREIGN KEY (requester_id) REFERENCES users (id);

ALTER TABLE donation_requests ADD FOREIGN KEY (accepted_by_user_id) REFERENCES users (id);

ALTER TABLE donation_requests ADD FOREIGN KEY (confirmed_by_user_id) REFERENCES users (id);

ALTER TABLE donation_events ADD FOREIGN KEY (donation_id) REFERENCES donations (id);

ALTER TABLE donation_events ADD FOREIGN KEY (created_by_user_id) REFERENCES users (id);

ALTER TABLE donation_events ADD FOREIGN KEY (reply_to_user_id) REFERENCES donation_events (id);

ALTER TABLE ratings ADD FOREIGN KEY (from_user_id) REFERENCES users (id);

ALTER TABLE ratings ADD FOREIGN KEY (to_user_id) REFERENCES users (id);

ALTER TABLE notifications ADD FOREIGN KEY (donation_id) REFERENCES donations (id);

ALTER TABLE notifications ADD FOREIGN KEY (donation_event_id) REFERENCES donation_events (id);

ALTER TABLE notifications ADD FOREIGN KEY (donation_request_id) REFERENCES donation_requests (id);

ALTER TABLE notifications ADD FOREIGN KEY (rating_id) REFERENCES ratings (id);

ALTER TABLE notifications ADD FOREIGN KEY (user_id) REFERENCES donation_requests (id);

ALTER TABLE notification_history ADD FOREIGN KEY (notification_id) REFERENCES notifications (id);

ALTER TABLE notification_history ADD FOREIGN KEY (user_id) REFERENCES users (id);

-- +goose Down
drop table notification_history;
drop table notifications;
drop table donation_requests;
drop table donation_events;
drop table donations;
drop table ratings;
drop table users;
