# Code-Competence
---
This Folder contains an application that allows admin to create, read, update, and delete items. The application is built using Go using the Echo framework and MySQL as the database. The application is using MVC (Model, View, Controller) as the design pattern. The application is also using JWT (JSON Web Token) for authentication and authorization. The application is also using Go Modules for dependency management.

## How to run the application
1. Clone this repository
    Go into the terminal and type the following command:
    ```
    git clone https://github.com/rifkeh/go-Rifkhi-Akbar.git

    ```
2. Go into the folder
    ```
    cd go-Rifkhi-Akbar/Code-Competence
    ```
3. Run the application
    ```
    go run main.go
    ```
4. Open the application in the browser (for best practice use Postman)
    ```
    http://localhost:8080
    ```
5. Type the endpoint you want to access
    ```
    http://localhost:8080/items
    ```
6. You can see the result in the browser or in Postman

## What is the application doing?
1. The application is using JWT for authentication and authorization
    ![Before using JWT](go-Rifkhi-Akbar/Code-Competence/images/WithoutJWT.png)