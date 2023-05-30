-- 全商品10%引きにする
select * from products;

-- 安全のため、更新できなくなっている仕組みをなくす
set sql_safe_updates = 0;  
update products set price = price * 0.9;

select * from products;