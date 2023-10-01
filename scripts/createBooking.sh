#!/bin/bash

# Check that there is two parameters
if [ $# -ne 2 ]; then
    echo "Usage: $0 <user_id> <class_id>"
    exit 1
fi

# Create the class
curl -X POST -H "Content-Type: application/json" -d '{ "user" : "'"$1"'", "class" : "'"$2"'", "date" : "2023-10-10T00:00:00Z" }' http://localhost:8080/bookings