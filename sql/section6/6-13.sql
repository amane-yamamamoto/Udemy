-- 2017年１月〜2017年６月までの月ごとのリクエスト数のうち、アクセス数が1000以上の月だけ抜き出して　
select
    request_month,
    count(*)
from
    access_logs
where
    request_month >= '2017-01-01'
    and request_month < '2017-07-01'
group by
    request_month
having
    count(*) >= 1000;