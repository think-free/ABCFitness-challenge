#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d '{ "studio" : "Studio 1", "class_name" : "Yoga", "start_date" : "2023-10-01T00:00:00Z", "end_date" : "2023-10-15T00:00:00Z", "capacity" : 10 }' http://localhost:8080/classes