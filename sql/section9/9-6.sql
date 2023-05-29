-- 注文詳細情報と商品情報を入れて、注文一覧を出力
select
	o.id order_id,
    o.user_id user_id,
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
	od.product_id = p.id;