## Introduction
This repository is written to test candidate's skills when it comes to writing clean and maintainable code base. In the real world, we rarely get the chance to develop our code base from scratch and most often we have to deal with the code which is either legacy or not ours. In order to write this repository, we wrote a simple tax calculator which was used to calculate the taxes you had to pay on your income. Then we remove the parts we didn't need and wrote a simple shopping cart manager. 


* Clone the repository
* Make small changes, commit
* Repeat
* The finished work should include all the git commits whether in a repository or zip file
## 
Changes Made
```
Hexagonal Architecture: The project structure has been modified to adhere to the Hexagonal Architecture pattern, promoting better organization and maintainability.

Repository Layer: A dedicated repository layer has been added to manage data access, ensuring a clean separation between business logic and data storage.

Enhanced Error Handling: Additional error handling has been implemented to improve the robustness of the application. While not overly advanced, it provides better feedback to users and logs errors for easier debugging.

SQL for Testing: The project now uses SQL for testing to eliminate the need for an extra dependency on MySQL, simplifying the testing environment.

Embed Templates: Templates are now embedded within the project to parse them, eliminating the need for direct file paths and improving portability.

Secret Key for Cookie Hashing: A secret key has been introduced to hash cookies, enhancing security and preventing unauthorized access.

SQL-Injection Prevention: Query modifications have been made to mitigate the risk of SQL injection, ensuring safer interactions with the database.

Integration Tests: Integration tests have been added for both the SQL layer and the service layer, ensuring a more comprehensive and reliable testing suite.
```


## How to run the test?
First you need docker to build the dependencies. You can skip this part if you already have mysql on your system. All you need to do it changing MySQL credentials in `pkg/db/get_db.go`
```
make test-integration
make test-e2e
```

## How to run the application?
First you need docker to build the dependencies. You can skip this part if you already have mysql on your system. All you need to do it changing MySQL credentials in `pkg/db/get_db.go`
```
cd docker
docker compose up -d --build
```

Once the containers are up and ready, you can run the application:
```
cd cmd
cd web-api
go run main.go
```
