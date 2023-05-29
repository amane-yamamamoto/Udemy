-- メルマガ送信用リスト（「名字＋さん」、「メアド」）で、女性のみに送信したい
select
	concat(last_name, 'さん'),
    email
from
	users
where
	gender = 2;