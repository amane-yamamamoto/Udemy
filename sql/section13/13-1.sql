-- 全期間における平均客単価（１回あたりで客が支払う金額）
-- 単価は少数第一位で四捨五入する
select
    round(avg(amount), 0) as '平均客単価'
from
    orders;