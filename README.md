Bank System API
===============

This project is a RESTful API for a banking system built with Golang, PostgreSQL, and JWT for authentication.

Features
--------

-   User authentication and authorization using JWT
-   Account management (create, update, delete)
-   Transaction management (deposit, withdraw, transfer)
-   Account balance inquiries
-   Transaction history retrieval

Technologies Used
-----------------

-   **Golang**: The programming language used for developing the API.
-   **PostgreSQL**: The database used for storing user, account, and transaction data.
-   **JWT**: Used for secure authentication and authorization.

Getting Started
---------------

### Prerequisites

-   [Go](https://golang.org/dl/) (version 1.16 or higher)
-   [PostgreSQL](https://www.postgresql.org/download/)
-   Docker (optional, for containerization)

### Installation

1.  **Clone the repository:**

    sh

    `git clone https://github.com/sandy1206/go-bank.git
    cd go-bank`

2.  **Set up PostgreSQL:**

    Create a database and user for the project. Update the `config.yaml` file with your database credentials.

3.  **Install dependencies:**


    `go mod tidy`

4.  **Run database migrations:**


    `go run cmd/migrate/main.go`

5.  **Start the server:**


    `go run main.go`

### Running with Docker

1.  **Build the Docker image:**


    `docker build -t bank-system-api .`

2.  **Run the Docker container:**


    `docker run -p 8080:8080 bank-system-api`

### API Endpoints

#### Authentication

-   **POST /api/auth/register**: Register a new user
-   **POST /api/auth/login**: Login and receive a JWT

#### Accounts

-   **POST /api/accounts**: Create a new account
-   **GET /api/accounts**: Get all accounts for the authenticated user
-   **GET /api/accounts/{id}**: Get account details by ID
-   **PUT /api/accounts/{id}**: Update account information
-   **DELETE /api/accounts/{id}**: Delete an account

#### Transactions

-   **POST /api/transactions/deposit**: Deposit funds into an account
-   **POST /api/transactions/withdraw**: Withdraw funds from an account
-   **POST /api/transactions/transfer**: Transfer funds between accounts
-   **GET /api/transactions**: Get all transactions for the authenticated user
-   **GET /api/transactions/{id}**: Get transaction details by ID

Configuration
-------------

Configuration settings are managed in the `config.yaml` file. Update this file with your database credentials, JWT secret, and other settings as needed.

Contributing
------------

Contributions are welcome! Please fork the repository and create a pull request with your changes.

License
-------

This project is licensed under the MIT License. See the LICENSE file for more details.
