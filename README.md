# SIoT Simulation

## Introduction
This project aims to simulate IoT devices using provided datasets. It includes functionalities to simulate IoT transactions, implement blockchain for data storage, send data to a Trust Management Server, and track performance metrics.

## Directory Structure
simulation/
├── cmd/
│   └── main.go
└── internal/
    ├── blockchain/
    │   ├── block.go
    │   └── blockchain.go
    ├── performance_tracking/
    │   └── performance_tracker.go
    ├── pkg/
    │   ├── iot_simulation/
    │   │   └── iot_simulation.go
    │   └── main.go
    └── trust_management/
        └── data_sender.go

## Functionality Overview
- **IoT Simulation**: Simulate IoT devices using provided datasets. Code for simulating devices can be found in `internal/pkg/iot_simulation/iot_simulation.go`.
- **Blockchain Implementation**: Implement blockchain for data storage. Code for blockchain functionality can be found in `internal/blockchain/blockchain.go`.
- **Data Sending to Trust Management Server**: Send data from blockchain to the Trust Management Server. Code for data sending can be found in `internal/trust_management/data_sender.go`.
- **Performance Tracking**: Track performance metrics of the IoT model and simulations. Code for performance tracking can be found in `internal/performance_tracking/performance_tracker.go`.

## How to Use
1. Clone the repository:
    ```bash
    git clone git@github.com:LAKSHYAKUMAR45/SIoT_Simulation.git
    cd SIoT_Simulation
    ```
2. Run the main program:
    ```bash
    go run cmd/main.go
    ```

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or create a pull request.

## License
This project is licensed under the [MIT License](LICENSE).
