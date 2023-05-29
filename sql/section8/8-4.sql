-- 商品一覧を税込価格で少数第一位で四捨五入して表示
select 
	id, name, round(price*1.10, 0)
from
	products;