# Project Stress Tester

This project is a CLI tool written in Go for performing load tests on a web service. 
The user can specify the URL of the service, the total number of requests, and the number of concurrent requests. 
The tool generates a report with specific information after the tests are executed.

## Features

- Perform HTTP requests to a specified URL.
- Distribute requests according to the defined concurrency level.
- Ensure the total number of requests is met.
- Generate a report with:
    - Total time spent on execution.
    - Total number of requests made.
    - Number of requests with HTTP status 200.
    - Distribution of other HTTP status codes (e.g., 404, 500, etc.).

## Prerequisites

- Go 1.20 or later
- Docker (optional, for running the application in a container)

## Installation

1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Build the project:
   ```sh
   go build -o project-stress
   ```

## Usage

### Running the Application

To run the application, use the following command:
```sh
./project-stress --url=http://httpbin.org/status/200,404,500 --requests=1000 --concurrency=100
```

### Running with Docker

1. Build the Docker image:
   ```sh
   docker build -t project-stress .
   ```

2. Run the Docker container:
   ```sh
   docker run project-stress --url=http://httpbin.org/status/200,404,500 --requests=1000 --concurrency=100
   ```

## Example Output

```
Tempo total gasto: 1m23.456s
Quantidade total de requests realizados: 1000
Quantidade de requests com status HTTP 200: 700
Distribuição de outros códigos de status HTTP:
Status 404: 200
Status 500: 100
```

## License

This project is licensed under the MIT License.