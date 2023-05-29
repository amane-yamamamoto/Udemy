-- 9-6に加えて、名字と名前を一覧に追加して表示
select
	o.id order_id,
    o.user_id user_id,
    u.last_name last_name,
    u.first_name first_name,
    o.amount amount,
    o.order_time order_time,
    p.name product_name,
    od.product_qty qty,
    p.price product_price
from
	orders o
inner join
	order_details od
on
	o.id = od.order_id
inner join
	products p
on
	od.product_id = p.id
inner join
	users u
on
	o.user_id = u.id;
	