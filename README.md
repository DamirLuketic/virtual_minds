# Virtual Minds
Dockerized Go API Client with MySQL DB.

# Summary:

## Run
- Client can be run with "docker-compose up"
- For demo purpose can be accessed through mock Authorization:
  - Basic authorization: 
    - Username: vm
    - Password: vm
- List of [mock data](./db/db_mock_data.go)

## Usage
- Expose port: 8080
- Credentials
  - Basic Authorization
- Endpoints:
    - Handle Request
      - Path: /api/new_request [post]
      - Params:
        - body [Request](./api/handler/type.go)
      - Client validation:
        - Authorization
        - Malformed JSON
        - Missing one or more fields
        - Customer UUID not in the DB or customer disabled
        - Remote IP in the blacklist
        - User Agent in the blacklist
    - Fetch Hourly Services:
      - Path: /api/customer_statistic [get]
      - Params:
        - query: "clientUUID" string
        - query: "date" string (format YYYY:MM:DD)
      - Client validation:
        - Authorization
        - Missing one or more fields
        - Customer UUID not in the DB or customer disabled
        - Remote IP in the blacklist
        - User Agent in the blacklist
- Further details:
  - [Handlers](./api/handler/handler.go)    