create table calendars (
  id          serial primary key,
  name        varchar(100) not null,
  visibility  boolean default true,
  color       varchar(16),
  created_at  timestamp not null default current_timestamp
);
SELECT setval('calendars_id_seq', 4, false);

create table events (
  id          serial primary key,
  name        varchar(100) not null,
  started_at  timestamp not null default current_timestamp,
  ended_at    timestamp not null default current_timestamp,
  calendar_id int,
  timed       boolean default true,
  description varchar(255),
  color       varchar(16),
  created_at  timestamp not null default current_timestamp,
  foreign key(calendar_id) references calendars(id)
);
SELECT setval('events_id_seq', 5, false);
