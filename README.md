# SIoT Simulation

## Introduction
This project aims to simulate IoT devices using provided datasets. It includes functionalities to simulate IoT transactions, implement blockchain for data storage, send data to a Trust Management Server, and track performance metrics.

## Functionality Overview
- **IoT Simulation**: Simulate IoT devices using provided datasets. Code for simulating devices can be found in `internal/pkg/iot_simulation/iot_simulation.go`.
- **Blockchain Implementation**: Implement blockchain for data storage. Code for blockchain functionality can be found in `internal/blockchain/blockchain.go`.
- **Data Sending to Trust Management Server**: Send data from blockchain to the Trust Management Server. Code for data sending can be found in `internal/trust_management/data_sender.go`.
- **Performance Tracking**: Track performance metrics of the IoT model and simulations. Code for performance tracking can be found in `internal/performance_tracking/performance_tracker.go`.

# SIoT Simulation

This project simulates an Internet of Things (IoT) environment for security analysis. It includes components for generating simulated IoT device behavior, predicting transaction safety, and tracking performance metrics.

## Components

### 1. IoT Simulation Package (`/internal/pkg/iot_simulation`)

This package contains functionality for simulating IoT devices and generating transaction data.

#### Files:
- `iot_simulation.go`: Simulates IoT devices and generates transaction data based on predefined models.

### 2. Performance Tracking Package (`/internal/pkg/performance_tracking`)

This package provides functionality for tracking performance metrics of the IoT model and simulations.

#### Files:
- `performance_tracker.go`: Defines a performance tracker struct for monitoring throughput, latency, prediction accuracy, false positive rate, and false negative rate. It includes methods for starting and stopping performance monitoring, as well as tracking events and logging performance data.

### 3. Main Package (`/cmd/main`)

The main package initializes and runs the simulation, including setting up IoT devices, starting performance tracking, and handling user interactions.

#### Files:
- `main.go`: Initializes the simulation environment, starts IoT device simulation, begins performance tracking, and handles user interactions.

## Usage

To run the simulation, execute the following command:

```bash
go run cmd/main/main.go
