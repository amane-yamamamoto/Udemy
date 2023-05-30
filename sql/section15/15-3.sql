use book_store;
show tables;
show columns from books;

-- booksテーブルに価格を管理する列を追加
alter table books add price int after id;
show columns from books;

-- price を unit_price に変更
alter table books change price unit_price int;
show columns from books;

-- unit_price を削除
alter table books drop unit_price;
show columns from books;