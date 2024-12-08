# Golang Serverless Application

This project is a serverless REST API built using AWS Lambda, DynamoDB, and the Go programming language. It provides functionality for managing user records in a DynamoDB table and is structured to facilitate scalability and maintainability.

---

## **Project Structure**

The project follows a modular file structure:

```
cmd
│   main.go
pkg
├── handlers
│   ├── handlers.go
│   ├── api_response.go
├── user
│   ├── user.go
├── validators
│   ├── is_email_valid.go
```

---

### **File Descriptions**

#### **`cmd/main.go`**
- The entry point of the application.
- Initializes AWS session and DynamoDB client.
- Implements AWS Lambda's request routing using `APIGatewayProxyRequest`.
- Routes HTTP methods (`GET`, `POST`, `PUT`, `DELETE`) to their respective handlers in the `pkg/handlers` package.

#### **`pkg/handlers/handlers.go`**
- Contains functions for handling user-related operations.
  - **`GetUser`**: Fetches user(s) based on query parameters.
  - **`CreateUser`**: Adds a new user to the DynamoDB table.
  - **`UpdateUser`**: Updates an existing user's data.
  - **`DeleteUser`**: Removes a user from the DynamoDB table.
  - **`UnhandledMethod`**: Handles unsupported HTTP methods.

#### **`pkg/handlers/api_response.go`**
- Provides the `apiResponse` function, which standardizes API responses for the application.
- Formats responses with a status code, headers, and body (JSON).

#### **`pkg/user/user.go`**
- Implements core logic for interacting with the DynamoDB table.
  - **`FetchUser`**: Fetches a single user by email.
  - **`FetchUsers`**: Retrieves all users.
  - **`CreateUser`**: Validates and adds a new user.
  - **`UpdateUser`**: Validates and updates user details.
  - **`DeleteUser`**: Deletes a user from the table.
- Uses AWS SDK to interact with DynamoDB.

#### **`pkg/validators/is_email_valid.go`**
- Provides the `IsEmailValid` function to validate email addresses using a regex pattern.
- Ensures email strings are within valid length limits and follow a proper format.

---

## **Setting Up the Project**

### **Prerequisites**
1. Install Go (version 1.16 or later).
2. Set up AWS CLI and configure your credentials.
3. Create a DynamoDB table with a primary key named `email`.
4. Set environment variables:
   - `AWS_REGION`: The AWS region for your DynamoDB table.
   - `TABLE_NAME`: The name of your DynamoDB table.

### **Installation**
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

---

## **Running Locally**
1. Use the [AWS SAM CLI](https://aws.amazon.com/serverless/sam/) or a similar tool to run Lambda functions locally:
   ```bash
   sam local start-api
   ```
2. Use a tool like `Postman` or `curl` to test the API endpoints.

---

## **Endpoints**

| Method | Endpoint        | Description              |
|--------|-----------------|--------------------------|
| GET    | `/users`        | Fetch all users.         |
| GET    | `/users?email=` | Fetch a user by email.   |
| POST   | `/users`        | Create a new user.       |
| PUT    | `/users`        | Update an existing user. |
| DELETE | `/users?email=` | Delete a user by email.  |

---

## **Testing**
- Use tools like `Postman` or `Insomnia` to test the API.
- Write unit tests for individual functions in the `pkg/` directory.

---

## **Contributing**
1. Fork the repository.
2. Create a new feature branch.
3. Commit your changes and create a pull request.

---

## **License**
This project is licensed under the MIT License.
