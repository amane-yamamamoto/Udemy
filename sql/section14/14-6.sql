-- 累計販売個数が10を超えている商品に関しては、価格を5%アップして
update
    products
set
    price = price * 1.05
where
    id in (
        select
            product_id
        from
            order_details
        group by
            product_id
        having
            sum(product_qty) >= 10

    );

-- select * from products