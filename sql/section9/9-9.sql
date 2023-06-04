-- ユーザとadminユーザを足し合わせた一覧
-- 列は、email, 姓、名、性別（数値）
select
	email,
    last_name,
    first_name,
    gender
from 
	users
union
select
	email,
    last_name,
    first_name,
    gender
from 
	admin_users;