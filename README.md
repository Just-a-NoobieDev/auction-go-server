# Online Auction System

## Overview

This project is the backend server of an online auction system built with Golang, PostgreSQL, and Gin - Gorm.

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

### Features Checklist

- [x]  Authentication
- [x]  User Management
  - User Registration: Users can create accounts.
  - User Profile: View and edit user profiles.
  - Admin Roles: Admin users with elevated permissions.
- [x]  Auction Management
  - Create Auction: Users (typically sellers) can create new auction listings.
  - Join and Leave Auction
  - Auction Details: View detailed information about an auction.
  - Edit/Delete Auction: Sellers can edit or delete their auction listings.
  - Search and Filter: Search and filter auctions by various criteria (e.g., category, price range).
- [x]  Bidding System
  - Place Bids: Users can place bids on active auctions.
  - [ ] Auto-Bidding: Implement automatic bidding up to a user-defined maximum. (Will do it after the auction lifecycle)
  - Bid History: View the history of bids on an auction.
- [x]  Real-Time Updates
  - WebSockets: Implement real-time updates for new bids and auction status.
  - [ ] Notifications: Notify users of outbid status, auction start/end, etc. (not yet implemented)
- [ ]  Auction Lifecycle
  - Scheduled Start/End: Auctions start and end at scheduled times.
  - Extend Auction Time: Extend auction time if bids are placed near the end.
  - Auction Status: Display status (upcoming, active, ended).
- [ ]  Reviews and Ratings
  - User Ratings: Allow users to rate each other.
  - Auction Reviews: Allow users to leave reviews on auctions.
- [ ]  Payment Processing
  - Payment Integration: Integrate with payment gateways for handling transactions.
  - Payment Verification: Verify and record successful payments.
  - Refunds: Handle refund processes for unsuccessful auctions or disputes.
