# Artaka-Show-Sleep-Merchants

## Register Admin

curl -i -X POST -H "Content-Type: application/json" -d '{
"username":"Guntur",
"email":"gunturkurniawan@gmail.com",
"secret_password":"admin12",
"phone":"081290858463"
}' https://artaka.herokuapp.com/api/admin/register

## Sigin admin with email & Password

curl -i -X POST -H "Content-Type: application/json" -d '{
"email":"gunturkurniawan@gmail.com",
"secret_password":"admin12"
}' https://artaka.herokuapp.com/api/admin/register

## Sigin admin with phone & Password

curl -i -X POST -H "Content-Type: application/json" -d '{
"phone":"081290858463",
"secret_password":"admin12"
}' https://artaka.herokuapp.com/api/admin/register

## Show Sleep Merchants

curl -i -X GET -H "Content-Type: application/json"  
https://artaka.herokuapp.com/api/admin/ShowSleep

## Get Post by id

curl -i -X GET -H "Content-Type: application/json"  
https://artaka.herokuapp.com/api/post/getpost/1

## Update Post by id

curl -i -X PUT -H "Content-Type: application/json" -d '{
"phone":"+6282264291947","contacted": "0","content":"Sudah diangkaat"

}' https://artaka.herokuapp.com/api/post/1

## Show Sleep Already

curl -i -X GET -H "Content-Type: application/json"  
https://artaka.herokuapp.com/api/admin/Already

## Show Sleep Not Respon

curl -i -X GET -H "Content-Type: application/json"  
https://artaka.herokuapp.com/api/admin/NotRespon

## Show Sleep Not Contacted

curl -i -X GET -H "Content-Type: application/json" https://artaka.herokuapp.com/api/admin/NotYetContact

# Postgres Live

API_SECRET=secret
DB_HOST=ec2-3-211-176-230.compute-1.amazonaws.com
DB_DRIVER=postgres
DB_USER=qgkhwjnkzvvmye
DB_PASSWORD=fdc4a8c8bd801dead798e1c8afc147844e0700724eb41d26185aa903eb39ccf4
DB_NAME=d72ff344g1ll62
DB_PORT=5432
PORT1=8089
