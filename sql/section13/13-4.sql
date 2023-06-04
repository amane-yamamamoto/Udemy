-- 都道府県別、月別平均客単価
-- 列、都道府県ID, 都道府県名、年月、平均客単価（少数第一位四捨五入）
-- 都道府県ID、年月昇順
select
    p.id,
    p.name,
    date_format(o.order_time, '%Y%m') as '年月',
    round(avg(o.amount), 0) as '平均客単価'
from
    orders o
inner join
    users u
on
    u.id = o.user_id
inner join
    prefectures p
on 
    u.prefecture_id = p.id
group by
    p.id,
    date_format(o.order_time, '%Y%m')
order by
    p.id,
    '年月'