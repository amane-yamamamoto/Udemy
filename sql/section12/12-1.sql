-- ユーザのアクティビティ度合いによって、施策を変えたい
-- ユーザの累計注文回数によって場合分け（A：５回以上、B:2回以上、C:1回、０回のユーザは出力不要）
-- ユーザID, 累計注文回数、ユーザランク
select
    u.id user_id,
    count(*) num,
    case
        when count(*) >= 5 then 'A'
        when count(*) >= 2 then 'B'
        else 'C'
    end as user_rank
from
    users u
inner join
    orders o
on
    u.id = o.user_id
group by
    u.id
order by
    user_rank;