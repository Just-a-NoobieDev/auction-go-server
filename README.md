# Online Auction System

## Overview

This project is an online auction system built with Golang, PostgreSQL, and Gin.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or higher
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/yourusername/auction-system.git
cd auction-system
```

### Install Dependency

```sh
go mod tidy
```

### Running PostgreSQL with Docker

Make sure you have Docker Desktop installed
Start the Docker Desktop and run this commands to the terminal

```sh
make db-up
make db-init
```

### Running the Application

```sh
make run
make dev
```

### Creating Database Migrations

Create a new SQL file in the migrations directory (e.g., migrations/20210725_create_items_table.sql).

Add your SQL commands to the file. For example:

```sh
sql
Copy code
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

Apply the migration:

```sh
make migrate
```

### Stopping the PostgreSQL Container

```sh
make db-down
```
