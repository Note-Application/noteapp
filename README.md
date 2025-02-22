
# NoteApp

A Powerful, Cloud based Note taking Application


![Logo](https://raw.githubusercontent.com/Note-Application/noteapp/refs/heads/main/assets/logo.png)

# Badges

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)   
[![Test Coverage](https://img.shields.io/badge/test--coverage-90%25-brightgreen.svg)](#)  
[![Kubernetes](https://img.shields.io/badge/kubernetes-deployed-blue)](#)  


# Features

## Authentication & Security  
🔑 **Google Authentication** – Secure login & user identity verification  
🔒 **Secure Data Storage** – Notes stored safely in PostgreSQL  
🛡️ **Role-Based Access Control (RBAC)** – Granular permissions for users  

## Performance & Scalability  
⚡ **gRPC-powered Note Updates** – Faster than traditional REST APIs  
🚀 **Kafka Event Streaming** – Handles high-traffic note updates  
📴 **Offline Support** – Continue working without an internet connection  

## Productivity & Collaboration  
📝 **Real-time Collaboration** – Edit notes together with instant sync  
🔄 **Autosave & Versioning** – Never lose progress, rollback changes  
📲 **Multi-Device Sync** – Seamlessly access notes across devices  
📱 **Mobile-Friendly UI** – Adaptive design for all screen sizes  

## Integrations  
🌐 **Envoy Proxy** – Efficient API gateway for gRPC & REST  
🕵️ **Istio Service Mesh** – Observability & traffic control  
🏗️ **Terraform-based Deployment** – Infrastructure as code for easy scaling  


# Tech Stack

### 🔧 **Backend:**  
- ⚡ **Go** – High-performance microservices  
- 🔗 **gRPC & REST** – Communication between services  
- 📩 **Kafka** – Realtime note updates  
- 🗄️ **PostgreSQL** – User & note data storage  
- ⚡ **Redis** – Caching & fast data retrieval  
- 🌉 **Envoy Proxy** – gRPC Web gateway  

### 🎨 **Frontend:**  
- ⚛️ **React.js** – Intuitive user interface  
- 🔐 **Google Login** – Secure authentication  
- ⚡ **gRPC-Web** – Efficient backend communication  
- 🎨 **Tailwind CSS** – Sleek UI styling  

### ☁ **Infrastructure & Deployment:**  
- 🐳 **Docker & Kubernetes** – Containerized & scalable deployment  
- 🔄 **Istio** – Service mesh for security & observability  
- ☁️ **AWS** – Cloud-based hosting  



# API Reference

### Create User

```grpc
  rpc CreateUser (CreateUserRequest) returns (UserResponse);
```

| Parameter     | Type   | Description            |
|--------------|--------|------------------------|
| `email`      | string | **Required.** User email |
| `name`       | string | **Required.** User name |
| `profile_pic`| string | Profile picture URL   |


### Get User By Email

```grpc
  rpc GetUserByEmail (GetUserByEmailRequest) returns (UserResponse);
```
| Parameter | Type   | Description                      |
|-----------|--------|----------------------------------|
| `email`   | string | **Required.** User email to fetch |

### Update User

```grpc
  rpc UpdateUser (UpdateUserRequest) returns (UserResponse);
```
| Parameter     | Type   | Description            |
|--------------|--------|------------------------|
| `email`      | string | **Required.** User email |
| `name`       | string | **Required.** User name |
| `profile_pic`| string | Profile picture URL   |

### Delete User

```grpc
  rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse);
```
| Parameter | Type   | Description                     |
|-----------|--------|---------------------------------|
| `email`   | string | **Required.** User email to delete |

### Create Note

```grpc
  rpc CreateNote (CreateNoteRequest) returns (NoteResponse);
```
| Parameter  | Type   | Description            |
|------------|--------|------------------------|
| `user_id`  | int32  | **Required.** User ID  |
| `title`    | string | **Required.** Note title |
| `content`  | string | **Required.** Note content |

### Get Note By ID

```grpc
  rpc GetNoteByID (GetNoteByIDRequest) returns (NoteResponse);
```
| Parameter | Type   | Description                   |
|-----------|--------|-------------------------------|
| `id`      | int32  | **Required.** Note ID to fetch |

### Get Notes By User ID

```grpc
  rpc GetNotesByUserID (GetNotesByUserIDRequest) returns (GetNotesByUserIDResponse);
```
| Parameter  | Type   | Description            |
|------------|--------|------------------------|
| `user_id`  | int32  | **Required.** User ID  |

### Update Note

```grpc
  rpc UpdateNote (UpdateNoteRequest) returns (NoteResponse);
```
| Parameter  | Type   | Description           |
|------------|--------|-----------------------|
| `id`       | int32  | **Required.** Note ID |
| `title`    | string | **Required.** New title |
| `content`  | string | **Required.** New content |


### Delete Note

```grpc
  rpc DeleteNote (DeleteNoteRequest) returns (DeleteNoteResponse);
```
| Parameter | Type   | Description           |
|-----------|--------|-----------------------|
| `id`      | int32  | **Required.** Note ID |


# Optimizations

We have made several optimizations to enhance performance, scalability, and user experience:

- **gRPC over REST** – Faster and more efficient communication  
- **Kafka Event Streaming** – Handles high-traffic note updates asynchronously  
- **Redis Caching** – Reduces database load and speeds up retrieval  
- **Efficient Query Optimization** – Improved PostgreSQL queries for performance  
- **Frontend Code Splitting** – Faster loading times with optimized React bundles  
- **Lazy Loading & Prefetching** – Loads only necessary components to enhance speed  
- **Accessibility Enhancements** – Improved UI for better usability across devices  



# Run Locally

### 📥 Clone the project  
```
  git clone https://github.com/Note-Application/noteapp.git
 ```

### 📂 Navigate to the project directory  
```
  cd notesapp
```

### 📦 Install dependencies  
```
  go mod tidy
```

### ▶ Start the gRPC server  
```
  go run cmd/server/main.go
  ```



# Environment Variables

To run this project, add the following environment variables to your `.env` file:

`API_KEY` – Your API key  
`ANOTHER_API_KEY` – Your other API key  
`DATABASE_URL` – Your database connection string  
`GRPC_SERVER_ADDRESS` – Your gRPC server address  
`KAFKA_BROKER` – Your Kafka broker URL



# Running Tests

You can run Go tests using the following command:

```bash
  go test ./...

```

If you want to see detailed output, use:

```bash
  go test -v ./...

```

To run tests for a specific package:

```bash
  go test -v ./path/to/package

```

To run a specific test function:

```bash
  go test -run TestFunctionName ./...

```

If you need to measure test coverage:

```bash
  go test -cover ./...

```

For generating a coverage report:

```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out

```




# Documentation

[Documentation](https://linktodocumentation)




# Contributing

Thank you for considering contributing to NoteApp! We welcome all contributions, whether it's bug reports, feature requests, or code contributions.

## Getting Started

**Fork the repository**: Click the "Fork" button on GitHub and clone your fork locally.
   ```sh
   git clone https://github.com/Note-Application/noteapp.git
   cd noteapp
   ```

**Set up the environment**:
   - Ensure you have Go, gRPC, and other dependencies installed.
   - Install dependencies using:
     ```sh
     go mod tidy
     ```

**Create a feature branch**:
   ```sh
   git checkout -b feature-branch-name
   ```

## Making Changes

- Follow best coding practices.
- Write clear commit messages:
  ```sh
  git commit -m "feat: add user authentication"
  ```
- Run tests before submitting:
  ```sh
  go test ./...
  ```

## Submitting a Pull Request

**Push changes to your fork**:
   ```sh
   git push origin feature-branch-name
   ```

**Open a Pull Request**:
   - Go to the original repository.
   - Click "New Pull Request".
   - Select your branch and add a descriptive title and explanation.

**Wait for Review**:
   - The maintainers will review your PR and suggest changes if necessary.

## Reporting Issues

If you find a bug or have a feature request:
- Open a [GitHub Issue](https://github.com/Note-Application/noteapp/issues)
- Provide a clear description and reproduction steps.

## Code of Conduct

Be respectful, constructive, and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

Happy Coding! 🚀






# Authors

- [@Ganesh Manchi](https://github.com/ganeshvenkatasai)



# Related Projects

Here are some related projects:

- [NoteApp Frontend](https://github.com/Note-Application/client) – React-based client with Google Login  



# Support

For support, email ganeshmanchi123@gmail.com




# Feedback

If you have any feedback, please reach out to us at ganeshmanchi123@gmail.com




## 🚀 About Me
I am a Software Engineer 👦 not because of having a CS Degree 🎓 or working in Software Company 💻 I make Software Products 📱 I write Quality Code ✍🏼 I love to Contribute ❤️




## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ganeshvenkatasai.github.io/Portfolio/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ganesh-manchi-a214501a0/)
[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://x.com/GanaManchi123)



# Skills

💻 **Frontend:** React, Next.js, HTML, CSS, Tailwind, JavaScript, TypeScript  
⚙ **Backend:** Golang, Node.js, Django, gRPC, REST APIs  
🗄 **Databases:** PostgreSQL, MySQL, MongoDB, Redis  
☁ **DevOps:** Docker, Kubernetes, Terraform, OpenTofu, Istio, Jenkins  
📡 **Messaging & Streaming:** Kafka, RabbitMQ  
🔍 **Testing:** Jest, Cypress, Go Testing  
🚀 **Cloud & Hosting:** AWS, Firebase, Vercel, Netlify  
