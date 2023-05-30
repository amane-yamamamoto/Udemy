-- 全商品を累計販売個数でランク分けする(A>=20, B>=10, C<10)
-- ランクが高い順に並び替える
-- 列は、商品ID,商品名、販売個数、ランク
select
    p.id,
    p.name,
    sum(od.product_qty) num,
    case
        when sum(od.product_qty) >= 20 then 'A'
        when sum(od.product_qty) >= 10 then 'B'
        else 'C'
    end as r
from
    products p
left outer join
    order_details od
on
    p.id = od.product_id
group by
    p.id
order by
    r;