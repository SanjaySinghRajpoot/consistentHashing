# Consistent Hashing Implementation

## Purpose
This project aims to provide a basic implementation of a Hash Ring data structure in Go. The primary motivation behind creating this project is to understand the inner workings of a Hash Ring and gain insight into the implementation details.

## What is a Hash Ring?
A Hash Ring is a data structure used in distributed systems to efficiently distribute data across multiple nodes. It allows for consistent hashing, which ensures that similar data is mapped to the same node, aiding in load balancing and fault tolerance.

## Implementation Overview
The project consists of a single `main.go` file containing the implementation of the Hash Ring data structure. Here's a brief overview of the key components:
- **HashRing**: Struct representing the Hash Ring, containing an array of HashNodes and the number of virtual nodes per physical node.
- **HashNode**: Struct representing a node in the Hash Ring, storing its hash value and server name.
- **CreateHashRing**: Function to create a new Hash Ring with a specified number of virtual nodes per physical node.
- **AddServer**: Method to add a new server to the Hash Ring.
- **RemoveServer**: Method to remove a server from the Hash Ring.
- **GetServer**: Method to retrieve the server responsible for a given key based on consistent hashing.

## Usage
To use the Hash Ring implementation:
1. Instantiate a Hash Ring using the `CreateHashRing` function, specifying the number of virtual nodes per physical node.
2. Add servers to the Hash Ring using the `AddServer` method.
3. Use the `GetServer` method to retrieve the server responsible for a given key based on consistent hashing.

## Example
```go
virtualNodes := 10
hashRing := CreateHashRing(virtualNodes)

hashRing.AddServer("server1")
hashRing.AddServer("server2")
hashRing.AddServer("server3")

// Retrieve the server responsible for the key
server := hashRing.GetServer("test")

fmt.Println("Server for key:", server)
```

## Future Improvements
- Enhance error handling and edge case scenarios.

