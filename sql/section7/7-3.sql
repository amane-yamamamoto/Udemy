-- ユーザ一覧を生年月日順にならべ、生年月日が同じ場合は都道府県ID順に並べる
select
    *
from
    users
order by
	birthday,
    prefecture_id;
