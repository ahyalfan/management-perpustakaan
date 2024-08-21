# Golang Base Project

This is a base project for implementasi simple fast identity online using Golang with the Fiber framework, Gorm for database operations, and PostgreSQL as the database. Configuration is managed via an env file.

## Features

- **Fiber Framework**: Fast and minimalistic web framework for Go.
- **GORM**: SQL builder and query library for Go.
- **PostgreSQL**: Relational database for storing application data.
- **Env File**: Simple configuration management using environment variables.
- **Validator**: simple validator for your application
- **Jwt**: JWT authentication token for your application

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Golang](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/shellrean/golang-base-project-clean-directory.git
   cd golang-base-project-clean-directory
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Create and configure `.env` file:**

   Create a `.env` file in the root directory and add your configuration variables.

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASS=yourpassword
   DB_NAME=yourdbname

   SERVER_HOST=localhost
   SERVER_PORT=8700

   JWT_KEY=secret
   JWT_EXP=100
   ```

4. **Set up PostgreSQL database:**

   Make sure your PostgreSQL server is running and create a database matching your `.env` configuration.

   ```sh
   psql -U yourusername -c "CREATE DATABASE yourdbname;"
   ```

### Running the Application

Start the application with the following command:

```sh
go run main.go
```
