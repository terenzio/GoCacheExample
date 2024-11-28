package main

import (
	"context"
	"log"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/marshaler"
	redis_store "github.com/eko/gocache/store/redis/v4"
	"github.com/redis/go-redis/v9"
)

// Employee represents the structure for employee data to be cached
type Employee struct {
	ID   string // Unique identifier for the employee
	Name string // Name of the employee
}

func main() {
	// Context for cache operations
	ctx := context.Background()

	// Step 1: Initialize Redis client and store
	redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	redisStore := redis_store.NewRedis(redisClient)
	log.Println("Initialized Redis client and store.")

	// Step 2: Initialize the cache manager with Redis store
	cacheManager := cache.New[any](redisStore)
	log.Println("Initialized cache manager with Redis store.")

	// Step 3: Initialize marshaler for handling complex object serialization
	marshal := marshaler.New(cacheManager)
	log.Println("Initialized marshaler for caching complex objects.")

	// Example Employee object to be cached
	employee := Employee{
		ID:   "101",   // Employee ID
		Name: "Alice", // Employee Name
	}
	log.Printf("Prepared Employee object for caching: %+v\n", employee)

	// Cache key and value
	key := employee.ID // Use Employee ID as the cache key

	// Step 4: Store the Employee object in the cache
	log.Printf("Storing Employee object in cache with key: %s\n", key)
	err := marshal.Set(ctx, key, employee)
	if err != nil {
		log.Fatalf("Failed to store Employee object in cache: %v\n", err)
	}
	log.Println("Successfully stored Employee object in cache.")

	// Step 5: Retrieve the Employee object from the cache
	log.Printf("Retrieving Employee object from cache with key: %s\n", key)
	var result Employee
	_, err = marshal.Get(ctx, key, &result)
	if err != nil {
		log.Fatalf("Failed to retrieve Employee object from cache: %v\n", err)
	}
	log.Printf("Successfully retrieved and unmarshalled Employee object: %+v\n", result)
}
