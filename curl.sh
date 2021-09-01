#! /bin/bash 


for i in {2..9}; do
    username="joaquin$i"
    password="password$i"
    email="juanito$i@hotmail.com"
# curl http://localhost:8080/api/v1/users/ -X POST -s -d '{"username":"'${username}'","password":"'${password}'","email":"'${email}'"}' -H "Content-Type:application/json"
# curl http://localhost:8080/api/v1/users/$i -s -X PUT  -d '{"username":"cambio'$i'usuario","password":"cambio password","email":"email@email.com"}' -H "Content-Type:application/json"
#  curl http://localhost:8080/api/v1/users/$i -s -X DELETE   -i 
done
