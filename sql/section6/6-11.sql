-- 2017年１月〜2017年６月までのアクセスログ
select
    *
from
    access_logs
where
    request_month >= '2017-01-01'
    and request_month < '2017-07-01';