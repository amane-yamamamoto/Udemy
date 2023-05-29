-- 2017年12月に商品を購入したユーザ一覧
-- ユーザID, 名字、e-mail

select 
	id,
    last_name,
    email
from
	users
where id in (
	select
		user_id
	from 
		orders
	where
		order_time >= '2017-12-01'
		and
		order_time < '2018-01-01'
);
