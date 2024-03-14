// This file contains code for simulating IoT devices using provided datasets.

// Prompt:
// - Define functions for reading and parsing datasets containing IoT transaction data.
// - Implement logic for simulating IoT transactions based on the parsed data.
// - Incorporate functionality for predicting the nature of transactions (malicious or benign) using a machine learning model.
// - Handle transaction outcomes (carry out or abort) based on predictions.

package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

// Define a struct to represent a transaction entry
type Transaction struct {
    Timestamp      time.Time
    UID            string
    OrigHost       string
    OrigPort       int
    RespHost       string
    RespPort       int
    Protocol       string
    Service        string
    Duration       float64
    OrigBytes      int
    RespBytes      int
    ConnectionState string
    LocalOrig      bool
    LocalResp      bool
    MissedBytes    int
    History        string
    OrigPkts       int
    OrigIpBytes    int
    RespPkts       int
    RespIpBytes    int
    TunnelParents  string
    Label          string
    DetailedLabel  string
}

// Function to read the dataset and return a slice of Transaction structs
func readDataset(filename string) ([]Transaction, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    var transactions []Transaction

    // Skip header line
    _, _ = reader.Read()

    for {
        record, err := reader.Read()
        if err != nil {
            break
        }

        timestamp, _ := strconv.ParseFloat(record[0], 64)
        ts := time.Unix(int64(timestamp), 0)

        origPort, _ := strconv.Atoi(record[3])
        respPort, _ := strconv.Atoi(record[5])
        duration, _ := strconv.ParseFloat(record[8], 64)
        origBytes, _ := strconv.Atoi(record[9])
        respBytes, _ := strconv.Atoi(record[10])
        missedBytes, _ := strconv.Atoi(record[14])
        origPkts, _ := strconv.Atoi(record[16])
        origIpBytes, _ := strconv.Atoi(record[17])
        respPkts, _ := strconv.Atoi(record[18])
        respIpBytes, _ := strconv.Atoi(record[19])

        transactions = append(transactions, Transaction{
            Timestamp:      ts,
            UID:            record[1],
            OrigHost:       record[2],
            OrigPort:       origPort,
            RespHost:       record[4],
            RespPort:       respPort,
            Protocol:       record[6],
            Service:        record[7],
            Duration:       duration,
            OrigBytes:      origBytes,
            RespBytes:      respBytes,
            ConnectionState: record[11],
            LocalOrig:      record[12] == "T",
            LocalResp:      record[13] == "T",
            MissedBytes:    missedBytes,
            History:        record[15],
            OrigPkts:       origPkts,
            OrigIpBytes:    origIpBytes,
            RespPkts:       respPkts,
            RespIpBytes:    respIpBytes,
            TunnelParents:  record[20],
            Label:          record[21],
            DetailedLabel:  record[22],
        })
    }

    return transactions, nil
}

// Function to simulate IoT devices
func simulateDevices(transactions []Transaction) {
    for _, transaction := range transactions {
        // Simulate device behavior using transaction data
        // For demonstration purposes, we're printing the details of each transaction
        fmt.Println("Simulating transaction:", transaction)
        
        // Example: Make a request to a prediction model API
        predictTransaction(transaction)
        
        // Simulated delay between transactions
        time.Sleep(1 * time.Second)
    }
}

// Function to make a request to a prediction model API
func predictTransaction(transaction Transaction) {
    // Example: Make a request to a prediction model API
    // Replace this with actual code to interact with your prediction model API
    // For demonstration purposes, we're just printing the transaction details
    fmt.Println("Making prediction for transaction:", transaction)
    
    // Example: HTTP POST request to prediction model API
    // resp, err := http.Post("http://prediction-model-api/predict", "application/json", bytes.NewBuffer(jsonData))
    // Handle response and error accordingly
}

func main() {
    // Path to the dataset file
    datasetFile := "converted_dataset_1.csv"

    // Read the dataset
    transactions, err := readDataset(datasetFile)
    if err != nil {
        fmt.Println("Error reading dataset:", err)
        return
    }

    // Simulate IoT devices using the dataset
    simulateDevices(transactions)
}
