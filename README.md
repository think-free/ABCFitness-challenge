# Code challenge ABCFitness

I'm implemented a kind of user persistance besides the required createClass and createBooking requirements
I've also added the ability to retreive the informations like the users, the classes, the bookings and to ability to retreive the complete information of a booking.
As the class overbooking, the persistance in a database where not required, it's not implemented.
Test coverage is not complete but I think that the basis is covered for the purpose of this exercice.

*******************


# Run it 

```shell
docker-compose up --build
```

It will listen on http://localhost:8080

# Test it

- Run tests :

```shell
go test ./...
```


# Sample response of the api

### Create user : 

##### Request 

```shell
curl -X POST -H "Content-Type: application/json" -d '{"name":"Elon","surname":"Musk", "email" : "elon.musk@example.com", "phone" : "+34123456789" }' http://localhost:8080/users
```

##### Response 

    {"status":"ok","data":{"id":"0d53e96d-8c85-41ec-b37b-7c39a75c35a5","name":"Elon","surname":"Musk","email":"elon.musk@example.com","phone":"+34123456789"},"metadata":{"createdAt":"2023-10-01T17:44:03Z"}}

### Create class : 

##### Request 

```shell
curl -X POST -H "Content-Type: application/json" -d '{ "studio" : "Studio 1", "class_name" : "Yoga", "start_date" : "2023-10-01T00:00:00Z", "end_date" : "2023-10-15T00:00:00Z", "capacity" : 10 }' http://localhost:8080/classes
```

    {"status":"ok","data":{"id":"d5e3e8db-8586-450e-a3d7-e3f8ff4f7ac4","studio":"Studio 1","class_name":"Yoga","start_date":"2023-10-01T00:00:00Z","end_date":"2023-10-15T00:00:00Z","capacity":10},"metadata":{"createdAt":"2023-10-01T17:44:07Z"}}

### Create booking :

##### Request 

```shell
curl -X POST -H "Content-Type: application/json" -d '{ "user" : "0d53e96d-8c85-41ec-b37b-7c39a75c35a5", "class" : "d5e3e8db-8586-450e-a3d7-e3f8ff4f7ac4", "date" : "2023-10-10T00:00:00Z" }' http://localhost:8080/bookings 
```

##### Response

    {"status":"ok","data":{"id":"52340bef-3dc6-4ee6-bf1b-01dfe5fa27fe","class":"d5e3e8db-8586-450e-a3d7-e3f8ff4f7ac4","user":"0d53e96d-8c85-41ec-b37b-7c39a75c35a5","date":"2023-10-10T00:00:00Z"},"metadata":{"createdAt":"2023-10-01T17:44:25Z"}}

### List bookings :

##### Request 

```shell
curl -X GET http://localhost:8080/bookings
```

##### Response

    {"status":"ok","data":[{"id":"52340bef-3dc6-4ee6-bf1b-01dfe5fa27fe","class":"d5e3e8db-8586-450e-a3d7-e3f8ff4f7ac4","user":"0d53e96d-8c85-41ec-b37b-7c39a75c35a5","date":"2023-10-10T00:00:00Z"}],"metadata":{"createdAt":"2023-10-01T17:44:32Z"}}

### Get one booking :

##### Request 

```shell
curl -X GET http://localhost:8080/booking?id=52340bef-3dc6-4ee6-bf1b-01dfe5fa27fe
```

##### Response

    {"status":"ok","data":{"id":"52340bef-3dc6-4ee6-bf1b-01dfe5fa27fe","date":"2023-10-10T00:00:00Z","class":{"id":"d5e3e8db-8586-450e-a3d7-e3f8ff4f7ac4","studio":"Studio 1","class_name":"Yoga","start_date":"2023-10-01T00:00:00Z","end_date":"2023-10-15T00:00:00Z","capacity":10},"user":{"id":"0d53e96d-8c85-41ec-b37b-7c39a75c35a5","name":"Elon","surname":"Musk","email":"elon.musk@example.com","phone":"+34123456789"}},"metadata":{"createdAt":"2023-10-01T17:44:53Z"}}