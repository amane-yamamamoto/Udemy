-- ユーザ一覧を都道府県名を明示して表示してほしい
-- 列はユーザID, 名字、名前、都道府県名
select
	u.id, 
    u.last_name,
    u.first_name, 
    p.name
from 
	users u
inner join
	prefectures p
on u.prefecture_id = p.id;