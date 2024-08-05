
# RESTful API IMPLEMENTED IN GO USING THE BEEGO FRAMEWORK 🚀

## UP-COMING 🔜
This API is under continuous development and receives updates regularly to enhance its functionalities and features. The goal is to evolve into a comprehensive API for forum management. Here are some of the exciting features and improvements that are on the way:

- **Thread Management**: Create, read, update, and delete discussion threads.
- **Commenting System**: Allow users to post comments on threads, including nested replies.
- **User Roles and Permissions**: Advanced role-based access control to manage permissions for different types of users (admins, moderators, members).
- **Search Functionality**: Implement search capabilities to find threads, posts, and users.
- **Notifications**: Real-time notifications for user interactions such as replies, mentions, and likes.
- **Profile Management**: Allow users to update their profiles with avatars, bio, and other personal information.
- **Activity Feed**: Provide a feed of user activity including new threads, comments, and likes.
- **Tagging System**: Enable tagging of threads and comments for better organization and searchability.
- **Moderation Tools**: Advanced tools for moderators to manage threads and comments, including reporting and blocking.
- **Statistics and Analytics**: Insights into user activity, popular threads, and other relevant statistics.

## Overview

A RESTful API written in Go, using the Beego framework. It provides user management functionalities, including user registration, login, and CRUD operations on user data. The API supports JWT-based authentication to secure the endpoints. 🔐

## Features

- User Registration and Authentication 📝
- JWT-Based Authentication 🔑
- Secure User Data Management 🛡️
- Role-Based Authorization 🧑‍⚖️

## Endpoints

### Public Endpoints 🌐

- **GET /**: Entry point for the application 🏠
- **POST /v1/api/auth/login**: Logs in a user and returns a JWT 🔐
- **POST /v1/api/auth/register**: Registers a new user 📝

### Protected Endpoints 🔒

- **GET /v1/api/users/**: Fetches a list of users (Requires authorization header with JWT) 👥
- **GET /v1/api/users/:id**: Fetches user's data by ID (Requires authorization header with JWT) 🆔
- **DELETE /v1/api/users/:id**: Deletes a user by ID (Requires authorization header with JWT) ❌

## Getting Started 🏁

### Prerequisites 📋

- Go 1.22 or higher 🚀
- Beego framework 🐝
- MySQL 🗄️

### Installation 💻

1. **Clone the repository:**
   ```sh
   git clone https://github.com/Gobakos/golang-forum-api.git
   cd beego-test-api
2. **Install dependencies**
   ```sh
   go mod tidy
3. **Create a .env file with:**
   ```
   SECRET_KEY=your_secret_key
   PASSWORD_KEY=your_password_key
4. **Change the default configurations in app.conf**
5. **Run the application**
   ```sh
   bee run
