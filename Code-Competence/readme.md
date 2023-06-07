# Code-Competence

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

## Another way to run application
1. Pull from docker  
    Go into terminal and type the following command:
    ```
    sudo docker pull rifkeh/code-competence:latest
    ```
2. Run the docker
    Run the docker by this following command:
    ```
    sudo docker run -d -p 8080:8080 code-competence:latest
    ```
 **OR**
1. Using the IPV4 that already launched
    Go into postman and use the following host and port
    ```
    http://ec2-3-26-71-156.ap-southeast-2.compute.amazonaws.com:8080/YOUR_ENDPOINT
    ```

## What is the application doing?
1. The application is using JWT for authentication and authorization  
    #### This is the application without JWT
    ![WithoutJWT](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/54ffb09b-ac0b-4f7f-a248-0dd6633c468e)
    #### This is the application with JWT
    ![AfterJWT](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/a3885aea-0136-48c9-b055-c0c508a3de53)  
2. The application is having some features:
    * **Register and Login Admin**  
        Admin is the one who has the main role in this application. Before you can access another feature, you have to register or login if u already have registered account. After login, you will get a JWT that will grant access to another feature.  
        #### Register Admin
        ![Register](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/3d52741f-4151-480f-8a2a-8e9920a731b1)
        #### Login Admin
        ![Login](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/6a63e38b-58c0-46b6-b03e-2aaf3af1941f)
    * **Getting items**
        There are four ways to get items from the database, getting all items, getting items by its id, getting items from category and getting items by its keyword
        #### Get All Items
        ![Screenshot from 2023-06-07 02-50-48](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/95d6b4f4-7890-433e-9f9d-744e5431d00a)
        #### Get Items by Id
        ![GetItemById](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/f82985a4-245b-44e6-9540-2d88c5155d68)

        #### Get Items By Category
        ![GetItemByCategoty](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/cc23098d-678c-48cc-ba61-eba48df7fab2)

        #### Get Items By Keyword
        ![GetItemByKeyword](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/5380548d-6256-4d3e-b869-8a1b7cff1b12)   

    * **Creating, Updating, and Deleting Items**
        There are feature to create, delete, and update item in the database.  
        #### Create Item
        ![CreateItem](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/3a8534fe-4a3c-4698-8804-cb8764408634)
        #### Delete Item
        ![DeleteItem](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/015ed9f4-ebb5-421f-929d-bff1c825c45d)
        #### Update Item
        ![UpdateItem](https://github.com/rifkeh/go-Rifkhi-Akbar/assets/91230120/505f8d65-f4c7-4020-876f-c894e5f8a271)
