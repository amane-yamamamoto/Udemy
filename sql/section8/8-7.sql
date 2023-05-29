-- 日付　
select current_date();
select current_timestamp();
select current_date() + 3;
select current_date() - 3;
select current_timestamp() + interval 6 hour;
select current_timestamp() - interval 6 hour;

-- order_timeカラムが2017/1のレコードを取得
select * from orders where extract(year_month from order_time) = 201701;
-- order_timeカラムが2017年のレコードを取得
select * from orders where extract(year from order_time) = 2017;
-- order_timeカラムが1月のレコードを取得
select * from orders where extract(month from order_time) = 1;