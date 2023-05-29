-- 2017年1月の東京に住むユーザの、注文情報を出す
-- 列は、注文id, 注文日時、注文金額合計、ユーザid, 名字、名前
select
	o.id,
    o.order_time,
    o.amount,
    u.id,
    u.last_name,
    u.first_name
from
	users u
inner join
	orders o
on
	u.id = o.user_id
where
	u.prefecture_id = 13
and 
    o.order_time >= '2017-01-01' and o.order_time < '2017-02-01'
order by 
	o.id;
-- order by  以下は不要 
    