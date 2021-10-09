
[![Issues][issues-shield]](https://github.com/Shiv10/Appointy-instagram-clone/issues)

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a>
    <img src="https://user-images.githubusercontent.com/17690376/136663828-9a59dcdb-003b-4a36-85df-67c8e2d4c258.png" alt="Logo" width="200">
  </a>

  <h3 align="center">Instagram Backend Clone</h3>

  <p align="center">
    Task for Appointy Internship 2021
    <br />
    <a href="https://github.com/Shiv10/Appointy-instagram-clone"><strong>Explore the docs »</strong></a>
    <br />
    <!-- <a href="https://www.youtube.com/watch?v=ifayZiXxAWo">View Demo</a>
    · -->
    <a href="https://github.com/Shiv10/Appointy-instagram-clone/issues">Report Bug</a>
    ·
    <a href="https://github.com/Shiv10/Appointy-instagram-clone/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Routes](#routes)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)



<!-- ABOUT THE PROJECT -->
## About The Project

This project is the task given for the Appointy interhsip for the year 2021. The project emulates the backend of Instagram in a very basic manner. There are a total of 5 routes which are listed below:

### Routes
Sr. No                     |  Description
---------------------------|--------------------------------------------------------------------------------------------------------
1. Create a User           |  POST request on route '/users' to add a user to the Database
2. Get user by ID          | GET request on route '/users/:id' to retrieve a user from the Db by using ID
3. Create a Post           |  POST request on route '/posts' to add a post to the Database
4. Get post by ID          |  GET request on route '/posts/:id' to retrieve a post from the Db by using ID
5. List all posts of a user|  GET request on route '/posts/users/:id' to retrive all posts of a user. (Pagination applied to make backend efficient)

| S. No.           | Description                                        |   |   |   |
|------------------|----------------------------------------------------|---|---|---|
| 1. Create a user | POST Request on route '/users' to a user to the DB |   |   |   |
|                  |                                                    |   |   |   |
|                  |                                                    |   |   |   |



### Built With

* [Go](https://golang.org/)
* [MongoDB](https://www.mongodb.com/)
* [Docker](https://www.docker.com/)



## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* go
* MongoDB installed on system or a Mongo Cluster
* Docker installed to run Docker image


### Installation
 
1. Clone the repo
```sh
git clone https://github.com/Shiv10/Appointy-instagram-clone
```
2. Change directory to Appointy-instagram-clone
```sh
cd Appointy-instagram-clone
```
3. Install mongo-go-driver
```sh
go mod download
```


## Usage
To start the backend server type the following in the command line:
```sh
go run main.go
```
Examples of routes using [Postman](https://www.postman.com/)
1. Create a user by sending POST request. Server responds with userID if created successfuly.
![s1](https://user-images.githubusercontent.com/17690376/136665288-0971b932-f3bb-49a2-ba26-eb44e2d7adc9.png)
<br/>
2. Sending request again with the same email. Server responds with eroor message.
![s2](https://user-images.githubusercontent.com/17690376/136665382-dd6e39d3-adaa-4ecb-a251-0d9869ab445b.png)

<br/>
3. Get user by sending userID

![s3](https://user-images.githubusercontent.com/17690376/136665478-861d98e9-9899-4935-94ab-b233a215fd95.png)

<br/>
4. Sending wrong id to server. Server responds with error message:

![s4](https://user-images.githubusercontent.com/17690376/136665585-fe2f5e49-fe44-4de5-9d06-35432e61d9da.png)

<br/>

5. Add post by sending POST request to the server. Server responds with postID
![s5](https://user-images.githubusercontent.com/17690376/136665639-f32ef91f-0197-4dee-9164-15ffd8f2b686.png)

<br/>

6. If we send wrong UserID for creating post, server responds with appropriate error message
![s6](https://user-images.githubusercontent.com/17690376/136665962-3366ba1a-12a2-47de-9d40-997a9f20aa9a.png)

<br/>

7. Get post by sending postID to server
![s7](https://user-images.githubusercontent.com/17690376/136666361-86814a3a-2f1c-4833-9bcc-9a6c517fa185.png)

<br/>

8. Get all User Posts by sending UserID
![s8](https://user-images.githubusercontent.com/17690376/136666502-18f544dc-c4f7-46cd-aa8f-06e9c260ad46.png)


## Roadmap

See the [open issues]((https://github.com/Shiv10/Appointy-instagram-clone/issues) for a list of proposed features (and known issues).


## Contributing

**Flow of contributions**

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'feat: Add some AmazingFeature'`)
4. Push to the Branch (`git push -u origin feature/AmazingFeature`)
5. Open a Pull Request

Happy Coding :smile:

## License

Distributed under the MIT License. See [`LICENSE`](./LICENSE) for more information.

