package main

type Operation struct {
    Type string `json:"type"`
    Value interface{} `json:"value"`
    ID interface{} `json:"id"`
    CreatedAt string `json:"created_at"`
}

type InputData struct {
    Company string `json:"company"`
    Operation *Operation `json:"operation"`
    Type string `json:"type"`
    Value interface{} `json:"value"`
    ID interface{} `json:"id"`
    CreatedAt string `json:"created_at"`
}

type OutData struct {
    Company string `json:"company"`
    ValidOperationsCount int `json:"valid_operations_count"`
    Balance int `json:"balance"`
    InvalidOperations []interface{} `json:"invalid_operations,omitempty"`
}

