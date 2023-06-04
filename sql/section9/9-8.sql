-- 商品IDが３に紐づく商品カテゴリ名を知りたい
-- 商品ID, 商品名、カテゴリ名
select
	p.id product_id,
    p.name product_name,
    c.name category_name
from
	products p
inner join
	products_categories pc
on
	p.id = pc.product_id
inner join
	categories c
on
	pc.category_id = c.id
where
	p.id = 3;