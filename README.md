# Money Transfer
Money Transfer is a basic service designed to facilitate sending and receiving money between users as 
well as a simple user management.

## Table of contents
* [Run Project](#Run)
* [Use the project](#usage)

## Run
To run this project, open git bash or cmd, then get into this project's directory using
```
$ cd moneytransfer
```
Then to run:
```
$ go run main.go
```
the server is now ready and up on your local then 
with 3 initial users are created and logged per requirements.

## usage
**Creat user**  
```
curl "localhost:8080/addUser?name=omar&balance=505.6" -X POST
```
* **name** is a mandatory field and is not unique.
* **balance** is not a mandatory field, default is 0 and can't be negative.  

The system will respond with the created user with an auto generated unique id.  
Response Sample
```
{
    "Id":"1c60a2ae-8672-4843-9164-f1b82ca068f5",
    "Name":"omar",
    "Balance":505.6
}
```

**Get user information**
```
curl localhost:8080/userInfo?id="1c60a2ae-8672-4843-9164-f1b82ca068f5" -X GET
```
Response Sample
```
{
    "Id":"1c60a2ae-8672-4843-9164-f1b82ca068f5",
    "Name":"omar",
    "Balance":505.6
}
```

**Get all users information**
```
curl localhost:8080/allUsers -X GET
```
Response Sample
```
[
    {
        "Id":"1c60a2ae-8672-4843-9164-f1b82ca068f5",
        "Name":"omar",
        "Balance":505.6
    },
    {
        "Id":"7d72c6e4-ab4a-4f9e-8b4a-ab271ac4481e",
        "Name":"Adam",
        "Balance":0
    }
]
```

**Delete user**
```
curl localhost:8080/deleteUser?id="d5d3383c-5aac-443d-bb63-07091a99ef1c" -X DELETE
```
Response Sample
```
user deleted successfully
```

**Money Transfer**
```
curl "localhost:8086/transfer?from="69974a81-4b16-4c91-81dc-4505af27a5fd"&to="7d72c6e4-ab4a-4f9e-8b4a-ab271ac4481e"&amount=50.75" -X POST
```
* **from** sender id, sender must have enough balance to do transaction with the amount specified and is a mandatory field.
* **to** receiver id and is a mandatory field.
* **amount** must be greater than 0, can have a decimal point and is mandatory field.

Response Sample
```
amount added successfully
```

