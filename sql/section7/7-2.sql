-- 商品一覧を価格が高い順にして、価格が同じ時は登録順に
select
    *
from
    products
order by
    price
desc,
    id;
