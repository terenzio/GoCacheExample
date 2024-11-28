# Employee Caching with Redis and GoCache

## Description

This project demonstrates how to cache and retrieve complex objects (like a struct) using Redis with the GoCache library in Go. The example caches an Employee struct and retrieves it using GoCache’s marshaler.

## Features

    •	Redis Integration: Uses Redis as the caching backend.
    •	GoCache: Simplifies caching with support for multiple stores.
    •	Marshaling: Serializes and deserializes complex objects for caching.
    •	Efficient Data Retrieval: Retrieves cached objects by keys.

## Requirements

    •	Go 1.20+
    •	A running Redis server (127.0.0.1:6379)
    •	The following Go packages:
    •	github.com/eko/gocache
    •	github.com/redis/go-redis

## Installation

    1.	Clone the repository:
    > git clone https://github.com/terenzio/GoCacheExample.git
    > cd employee-caching

    2.	Install dependencies:
    > go mod tidy

    3.	Ensure Redis is running:
    > docker run --name redis -d -p 6379:6379 redis

## Usage

Execute the application:

    > go run main.go

Expected Output

    1.	The Employee struct is serialized and stored in the Redis cache.
    2.	The struct is retrieved, deserialized, and displayed in the logs.

Example log output:

    2024/11/27 10:00:00 Initialized Redis client and store.
    2024/11/27 10:00:00 Initialized cache manager with Redis store.
    2024/11/27 10:00:00 Initialized marshaler for caching complex objects.
    2024/11/27 10:00:00 Prepared Employee object for caching: {ID:101 Name:Alice}
    2024/11/27 10:00:00 Storing Employee object in cache with key: 101
    2024/11/27 10:00:00 Successfully stored Employee object in cache.
    2024/11/27 10:00:00 Retrieving Employee object from cache with key: 101
    2024/11/27 10:00:00 Successfully retrieved and unmarshalled Employee object: {ID:101 Name:Alice}

## Code Explanation

Main Steps

    1.	Initialize Redis Client:
    •	Connect to the Redis server using go-redis.
    •	Create a Redis store for caching operations.
    2.	Set Up Cache Manager:
    •	Initialize the cache manager with the Redis store.
    3.	Use Marshaler:
    •	Wrap the cache manager with a marshaler to handle serialization and deserialization of complex objects like the Employee struct.
    4.	Cache an Employee:
    •	Use the marshal.Set method to store the Employee object in Redis.
    5.	Retrieve an Employee:
    •	Use the marshal.Get method to retrieve and unmarshal the cached object.

Employee Struct

    The Employee struct contains:
    • ID: A unique identifier for the employee.
    • Name: The name of the employee.

## Customization

Change Redis Address

To use a different Redis server, update the address in the redis.Options:

    redisClient := redis.NewClient(&redis.Options{Addr: "YOUR_REDIS_SERVER:PORT"})

## Add More Data

You can extend the Employee struct with additional fields, such as:

    type Employee struct {
    ID string
    Name string
    Department string
    Salary float64
    }

## Dependencies

    •	GoCache: A multi-store caching library.
    •	Go-Redis: A Redis client for Go.

Install them using:

    go get github.com/eko/gocache
    go get github.com/redis/go-redis/v9

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
