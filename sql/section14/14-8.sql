-- 商品ID 1001 を削除する
select * from products;

delete from
    products
where
    id = 1001;

select * from products