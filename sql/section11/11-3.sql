-- 全商品の平均単価よりも単価が高い商品一覧 
-- 商品単価の高い順、同じ時は登録順に
select
	*
from
	products
where 
	price > (
		select
			avg(price)
		from
			products
    )
order by
	price
desc,
	id;