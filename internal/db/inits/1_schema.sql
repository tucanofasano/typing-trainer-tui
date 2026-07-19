
create table if not exists `words` (
	id integer primary key autoincrement,
	word text not null unique
);

create table if not exists `user` (
	id integer primary key autoincrement,
	name text not null unique
);

create table if not exists `stat` (
	id integer primary key autoincrement,
	id_user integer not null,
	ngram text not null,
	attempts integer not null default 0,
	errors decimal not null default 0,

	unique (id_user, ngram),
	foreign key (id_user) references user(id)
);
