# Postgres Live

API_SECRET=secret
DB_HOST=127.0.0.1
DB_DRIVER=postgres
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=forum_db_test
DB_PORT=5432
PORT1=8087

# Artaka-Show-Sleep-Merchants

curl -i -X POST -H "Content-Type: application/json" -d '{

    "username":"Guntur Budi Kurniawan",
    "email":"gunturkurniawan238@gmail.com",
    "secret_password":"admin123",
    "phone":"081290858462",

"idcard_image": ["file", "file"]

}' https://artaka.herokuapp.com/api/admin/register

## Create User

curl -i -X POST -H "Content-Type: application/json" -d '{
"email":"guntur21@gmail.com",
"password":"admin123"
}' http://localhost:8086/api/admin/login

curl -i "https://artaka.herokuapp.com/api/admin/ShowSleep"

## Login Admin

curl -i -X POST -H "Content-Type: application/json" -d '{
"email":"guntur21@gmail.com",
"password":"admin123"
}' http://localhost:8086/api/admin/login

curl -i "https://artaka.herokuapp.com/api/admin/ShowSleep"

## Update Admin with token

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJpZCI6MjJ9.uftt2oMYscnFHqIbKw7O58ZKiEC3e44t6wGYgh_u20M
curl -i -X PUT -H "Content-Type: application/json" -d '{
"username":"gunturaedbhwdwdasedaewa",
"email":"guntuqdweewdusbjhxwf21@gmail.com",
"password":"admin123",
"phone":"081290858472"
}' http://localhost:8086/api/admin/update/22

## Register Merchants

curl -i -X POST -H "Content-Type: application/json" -d '{
"user_id":"+62811196196",
"owner_name":"Iim Rusyamsi",
"fcm_token":"teststing",
"idcard_name":"291947",
"idcard_number":"291947",
"bank_holder_name":"291947",
"bank_name":"291947",
"bank_account":"291947",
"referral_code":"291947",
"email":"guntur@gmail.com",
"secret_password":"291947jdj",
"idcard_image":"291947"
}' http://localhost:8086/api/merchant/register

## Login Merchants

curl -i -X POST -H "Content-Type: application/json" -d '{
"email":"guntur@gmail.com",
"secret_password":"291947jdj"
}' http://localhost:8086/api/merchant/login

## Update Merchant with token

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJpZCI6MjJ9.uftt2oMYscnFHqIbKw7O58ZKiEC3e44t6wGYgh_u20M
curl -i -X PUT -H "Content-Type: application/json" -d '{
"user_id":"+62811196199",
"owner_name":"Iim Rusyamsa",
"fcm_token":"teststi",
"idcard_name":"2919",
"idcard_number":"2919",
"bank_holder_name":"21947",
"bank_name":"29194",
"bank_account":"29947",
"referral_code":"21947",
"email":"gunturaa@gmail.com",
"secret_password":"29147jdj",
}' http://localhost:8086/api/merchant/update/1

## Saved Order

curl -i -X POST -H "Content-Type: application/json" -d '{
"outlet_id":"ok-oce",
"saved_orders_id":"tetqtqt",
"name":"Guntur",
"phone":"081290858472",
"table_id":"25"
}' https://artaka1.herokuapp.com/api/transaction/savedorder

## Get ALl Merchant

curl -i -X GET -H "Content-Type: application/json"  
http://localhost:8086/api/admin/getall

## GET Merchant by id

curl -i -X GET -H "Content-Type: application/json"
http://localhost:8086/api/admin/getall/1

## Create Feed Back with token

curl -i -X POST -H "Content-Type: application/json" -d '{
"phone":"81290191",
"content":"hahaha",
"author_id":28
}' http://localhost:8086/api/post/create

SELECT user_id, owner_name, email, Z.created_at as last_trx FROM (
SELECT user_id,owner_name, email, (SELECT created_at FROM sales1 WHERE created_at > current_date-7 AND user_id = b.user_id ORDER BY id DESC LIMIT 1) FROM subscribers1 b
UNION SELECT user_id, owner_name, email, (SELECT created_at FROM onlinesales1 WHERE created_at > current_date-7 AND user_id = b.user_id ORDER BY id DESC LIMIT 1) FROM subscribers1 b
UNION SELECT user_id, owner_name, email, (SELECT created_at FROM saved_orders so WHERE created_at > current_date-7 AND user_id = b.user_id ORDER BY id DESC LIMIT 1) FROM subscribers1 b) AS Z

# Postgres Live

heroku pg:psql --app https://git.heroku.com/artaka.git < artakatable_web_monitor1.sql

/usr/bin/pg_dump --file "/var/lib/pgadmin/storage/gunturkurniawan238_gmail.com/artaka.sql" --host "localhost" --port "5432" --username "postgres" --no-password --verbose --format=t --blobs "forum_db_test"

API_SECRET=secret
DB_HOST=ec2-3-211-176-230.compute-1.amazonaws.com
DB_DRIVER=postgres
DB_USER=qgkhwjnkzvvmye
DB_PASSWORD=fdc4a8c8bd801dead798e1c8afc147844e0700724eb41d26185aa903eb39ccf4
DB_NAME=d72ff344g1ll62
DB_PORT=5432
PORT1=8087
