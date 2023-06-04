-- 商品ID=3の商品名を「SQL入門」に変える

select * from products;

update products set name='SQL入門' where id = 3;

select * from products where id = 3;