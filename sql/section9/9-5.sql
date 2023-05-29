-- すべての商品についての販売個数一覧
select
	p.id,
    p.name,
    sum(od.product_qty) num
from
	products p
left outer join
	order_details od
on 
	p.id = od.product_id
group by
	p.id;
