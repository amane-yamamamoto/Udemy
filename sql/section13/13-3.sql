-- 都道府県別平均客単価
-- 列、都道府県ID, 都道府県名、平均客単価（小数点以下第一位で四捨五入）
-- 都道府県ID昇順
select
    p.id,
    p.name,
    round(avg(o.amount), 0) as '平均客単価'
from
    orders o
inner join
    users u
on 
    o.user_id = u.id    
inner join
    prefectures p
on
    p.id = u.prefecture_id
group by
    p.id
order by
    p.id;
