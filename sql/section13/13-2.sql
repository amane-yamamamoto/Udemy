-- 月別の平均客単価
-- 少数第一位で四捨五入
select
    date_format(order_time, '%Y/%m') as '年/月',
    round(avg(amount), 0) as '月別平均客単価'
from
    orders
group by
    date_format(order_time, '%Y/%m')
order by
    '年/月';
-- order by も念のため指定