create table users(
	id text not null primary key,
	name varchar(20) not null,
	email varchar(100) not null,
	password text not null,
	created_at timestamp default current_timestamp
)

create table rooms(
	id text not null primary key,
	name varchar(20) not null,
	user_id text not null,
	created_at timestamp default current_timestamp,
	constraint fk_rooms_user_id foreign key(user_id) references users(id)
)

select * from users
select r.id, r.name, r.user_id, r.created_at, u.id, u.name, u.email, u.created_at
	from rooms as r
	join users as u on u.id = r.user_id
	
select * from rooms where id = '2a541fe3-ca2a-4cb0-a2d5-e802bc8ad7d9'
select * from rooms

create table messages (
	username varchar(50) not null,
	contents text not null,
	room_id text not null,
	user_id text not null,
	created_at timestamp default current_timestamp,
	constraint fk_messages_user_id foreign key(user_id) references users(id),
	constraint fk_messages_room_id foreign key(room_id) references rooms(id)
)

drop table messages
select * from messages


