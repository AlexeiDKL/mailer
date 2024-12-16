-- SQLite

-- INSERT INTO companys (name) VALUES ('EEC');

-- DELETE FROM companys WHERE name = 'EEC';

-- CREATE TABLE IF NOT EXISTS companys (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	name TEXT UNIQUE
-- );

-- CREATE TABLE IF NOT EXISTS domens (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	domen TEXT UNIQUE,
-- 	company INTEGER NOT NULL,
-- FOREIGN KEY(company) REFERENCES company(id)
-- );

-- CREATE TABLE IF NOT EXISTS Users (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	mail TEXT NOT NULL,
-- 	domen INTEGER NOT NULL,
-- FOREIGN KEY(domen) REFERENCES domens(id)
-- );

-- CREATE TABLE IF NOT EXISTS Pins (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	User INTEGER NOT NULL,
-- FOREIGN KEY(User) REFERENCES Users(id)
-- );

-- CREATE TABLE IF NOT EXISTS mails (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	user INTEGER NOT NULL,
-- 	body TEXT NOT NULL,
-- 	sending REAL NOT NULL DEFAULT '0',
-- FOREIGN KEY(user) REFERENCES Users(id)
-- );

-- CREATE TABLE IF NOT EXISTS links (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	company TEXT NOT NULL,
-- 	link TEXT NOT NULL,
-- 	link_type INTEGER NOT NULL,
-- FOREIGN KEY(company) REFERENCES company(id),
-- FOREIGN KEY(link_type) REFERENCES Link_type(id)
-- );

-- CREATE TABLE IF NOT EXISTS Link_type (
-- 	id integer primary key NOT NULL UNIQUE,
-- 	prod INTEGER NOT NULL UNIQUE,
-- 	test INTEGER NOT NULL,
-- 	dev INTEGER NOT NULL UNIQUE
-- );