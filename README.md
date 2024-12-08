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
- Entry point of the application.
- Initializes AWS session and DynamoDB client.
- Routes HTTP methods (`GET`, `POST`, `PUT`, `DELETE`) to their respective handlers.

#### **`pkg/handlers/handlers.go`**
- Implements HTTP handlers for user-related operations:
  - **`GetUser`**: Fetches user(s) based on query parameters.
  - **`CreateUser`**: Adds a new user to the DynamoDB table.
  - **`UpdateUser`**: Updates an existing user's data.
  - **`DeleteUser`**: Removes a user from the DynamoDB table.
  - **`UnhandledMethod`**: Handles unsupported HTTP methods.

#### **`pkg/handlers/api_response.go`**
- Provides the `apiResponse` function to format API responses with status codes, headers, and JSON bodies.

#### **`pkg/user/user.go`**
- Contains the core logic for interacting with DynamoDB:
  - **`FetchUser`**: Fetches a single user by email.
  - **`FetchUsers`**: Retrieves all users.
  - **`CreateUser`**: Validates and adds a new user.
  - **`UpdateUser`**: Validates and updates user details.
  - **`DeleteUser`**: Deletes a user from the table.

#### **`pkg/validators/is_email_valid.go`**
- Provides the `IsEmailValid` function to validate email addresses using regex.

---

## **Setup and Configuration**

### **Prerequisites**
1. Install Go (version 1.16 or later).
2. Configure AWS CLI with valid credentials.
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
1. Use the [AWS SAM CLI](https://aws.amazon.com/serverless/sam/) to run the application locally:
   ```bash
   sam local start-api
   ```
2. Test the endpoints with `curl` or tools like `Postman`.

---

## **API Endpoints and Example Commands**

### **1. Create a New User**
- **Endpoint**: `POST /users`
- **Command**:
  ```bash
  curl --header "Content-Type: application/json" \
       --request POST \
       --data '{"email":"chdvanshsingh@gmail.com", "firstname":"Vansh", "lastname":"Singh"}' \
       https://<api-gateway-url>/users
  ```

### **2. Get All Users**
- **Endpoint**: `GET /users`
- **Command**:
  ```bash
  curl --request GET https://<api-gateway-url>/users
  ```

### **3. Get a User by Email**
- **Endpoint**: `GET /users?email=<email>`
- **Command**:
  ```bash
  curl --request GET https://<api-gateway-url>/users?email=chdvanshsingh@gmail.com
  ```

### **4. Update a User**
- **Endpoint**: `PUT /users`
- **Command**:
  ```bash
  curl --header "Content-Type: application/json" \
       --request PUT \
       --data '{"email":"chdvanshsingh@gmail.com", "firstname":"VanshUpdated", "lastname":"SinghUpdated"}' \
       https://<api-gateway-url>/users
  ```

### **5. Delete a User**
- **Endpoint**: `DELETE /users?email=<email>`
- **Command**:
  ```bash
  curl --request DELETE https://<api-gateway-url>/users?email=chdvanshsingh@gmail.com
  ```

---

## **Testing**
- Use `curl`, `Postman`, or other tools to test the API.
- Write unit tests for individual functions in the `pkg/` directory.

---

## **Contributing**
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes and create a pull request.

---

## **License**
This project is licensed under the MIT License.
