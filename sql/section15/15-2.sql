-- 列id, title をもつ　booksテーブルを作成
-- 使用するDB指定
use book_store;
--  テーブル一覧
show tables;
-- テーブル作成
create table books (id int not null auto_increment primary key, title varchar(255) not null);

show tables;
show columns from books;