# Redis Caching Performance Test

This project demonstrates how Redis can improve the speed of database operations by caching responses. It compares the performance of database operations with and without Redis caching.

## Requirements

Before running the project, ensure the following are installed either **locally** on your computer or **set up in the cloud**:

- **Redis**: You need a Redis instance running. You can:
  - Download Redis locally from the [official Redis website](https://redis.io/download).
  - Or use a cloud-hosted Redis instance (e.g., AWS ElastiCache, Redis Cloud).

- **PostgreSQL**: You need a PostgreSQL database. You can:
  - Download PostgreSQL locally from the [official PostgreSQL website](https://www.postgresql.org/download/).
  - Or use a cloud-hosted PostgreSQL instance (e.g., AWS RDS, Google Cloud SQL).

- **Golang**: Install Go on your machine. You can download it from the [official Go website](https://golang.org/dl/).

Make sure the environment variables (DB credentials, Redis host/port, etc.) are set properly based on your setup.

## How It Works

1. The project sets up a PostgreSQL database and a Redis cache.
2. It implements CRUD operations for a `Product` model, with both cached and non-cached versions.
3. The `compare/compare.go` file contains benchmark tests that measure the performance of these operations.
4. Results are printed to the console, showing the time taken for operations with and without Redis caching.

## Project Structure

- `cmd/`: Contains the main application entry point.
- `compare/`: Houses the comparison testing code.
- `config/`: Manages configuration loading from environment variables.
- `datagen/`: Includes data generation utilities.
- `db/`: Contains database-related code:
  - `model/`: Defines data models.
  - `repository/`: Implements data access logic.
  - `service/`: Provides business logic and interfaces with the repository.
- `rclient/`: Implements Redis client functionality.
- `server/`: Contains HTTP server and route handlers.

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/redis-caching-test.git
   cd redis-caching-test
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up your environment variables in a `.env` file:
   ```
   HOST=your_db_host
   USER=your_db_user
   PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=your_db_port
   REDIS_HOST=your_redis_host
   REDIS_PORT=your_redis_port
   REDIS_PASSWORD=your_redis_password
   ```

## Project Setup

### 4. Initialize the Database and Start the Server

To load the data from `product.csv` into the `product` table and start the server, follow these steps:

1. **Uncomment the CSV loading lines** in `cmd/main.go` that handle inserting the data from `product.csv` into the database.

2. Once the lines are uncommented, run the following command to execute the CSV data insertion and start the server:

   ```bash
   go run cmd/main.go
   ```
This will load the product.csv data into the product table and simultaneously start the server.

## Performance Results

Here are the performance benchmarks for 100 requests, showing the time taken for both cached and non-cached operations:

### Fetching 1 Product 100 Times:
- **Without Redis**: 2.449971s
- **With Redis**: 58.3227ms
- **Speed Improvement**: 97.62%

### Fetching All Products 100 Times:
- **Without Redis**: 6.2698453s
- **With Redis**: 4.9381502s
- **Speed Improvement**: 21.24%

### Updating 1 Product 100 Times:
- **Without Redis**: 2.3989448s
- **With Redis**: 5.0500001s
- **Speed Improvement**: -110.51%

## Analysis

- **Read Operations**: Redis caching provides a significant boost for frequent reads, especially for single product retrieval, with an improvement of 97.62%.
  
- **Bulk Reads**: Retrieving all products showed moderate improvement (21.24%) since the overhead of fetching large datasets from Redis remains higher than smaller queries.

- **Write Operations**: Updating products with Redis introduces overhead (-110.51%) due to cache invalidation, meaning Redis needs to clear the old data from the cache after an update. This is a common trade-off with write-heavy workloads.

## Conclusion

Redis is highly effective for improving **read-heavy** operations, particularly for frequently accessed items like individual products. However, it introduces some overhead in **write-heavy** operations due to the need for cache invalidation. When implementing Redis caching, it's crucial to consider your application's **read/write ratio** and the **access patterns**.

### Recommendations:
- Use Redis caching for **read-heavy workloads** or frequently accessed items.
- Be cautious when applying Redis to **write-heavy operations**, as cache invalidation adds overhead.
- For bulk retrieval operations, consider additional optimizations such as pagination or partial caching.

## Future Improvements

- **Pagination Caching**: Implement pagination for bulk retrievals to reduce data size and improve performance.
- **Cache Invalidation Strategies**: Explore more efficient cache invalidation strategies to mitigate the performance hit for update operations.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
