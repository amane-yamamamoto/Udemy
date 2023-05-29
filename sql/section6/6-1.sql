-- 2017年１月の合計売上金額
select
    sum(amount)
from
    orders
where
    order_time >= '2017-01-01'
    and order_time < '2017-02-01';
    