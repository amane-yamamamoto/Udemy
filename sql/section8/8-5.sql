-- ユーザ一覧を「名字＋スペース＋名前さん」で出力してほしい
select
	concat(last_name, ' ', first_name, 'さん')
from 
	users;