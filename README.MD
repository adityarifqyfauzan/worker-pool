# Worker Pool Example

This project is a simple worker pool example. The service provides endpoints to check user IDs with and without a worker pool, demonstrating concurrency in Go using goroutines and channels.



## Table of Contents
- [Worker Pool Example](#worker-pool-example)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Endpoints](#endpoints)
    - [Example Request](#example-request)
  - [Contributing](#contributing)

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/adityarifqyfauzan/worker-pool.git
   cd worker-pool
2. **Install dependencies:**
    ```bash
    go mod tidy
3. **Run the application:**
    ```bash
    go run main.go
## Usage

The service provides two endpoints to check user IDs:

### Endpoints

1. **`POST` /checkUserIDs**
   Checks user IDs using a worker pool for concurrent processing.
2. **`POST` /checkUserIDsWithoutWorkerPool**
   Checks user IDs sequentially without using a worker pool.

### Example Request

1. **`POST` /checkUserIDs**
- Request
    ```json
    {
        "user_ids": [1, 2, 3, 40, 70]
    } 
- Response
    ```json
    {
        "found_user_ids": [1, 2, 3, 40, 70]
    }
2. **`POST` /checkUserIDsWithoutWorkerPool**
- Request
    ```json
    {
        "user_ids": [1, 2, 3, 40, 70]
    } 
- Response
    ```json
    {
        "found_user_ids": [1, 2, 3, 40, 70]
    }
## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
1. Fork the repository
2. Create your feature branch `git checkout -b feature/my-new-feature`
3. Commit your changes `git commit -am 'Add some feature'`
4. Push to the branch `git push origin feature/my-new-feature`
5. Create a new Pull Request