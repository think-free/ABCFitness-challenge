#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d '{"name":"Elon","surname":"Musk", "email" : "elon.musk@example.com", "phone" : "+34123456789" }' http://localhost:8080/users
