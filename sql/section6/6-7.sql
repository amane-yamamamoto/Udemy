-- 都道府県別のユーザ数
select
    prefecture_id,
    count(*)
from
    users
group by
    prefecture_id;
