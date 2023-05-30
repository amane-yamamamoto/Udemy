-- 一つも売れてない商品は売るのをやめにして削除する
delete from
    products
where
    id not in (
        select
            product_id
        from
            order_details
        group by
            product_id
    );