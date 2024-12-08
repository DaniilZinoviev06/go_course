package main

import (
    "fmt"
    "sort"
)

// функция проверки валидного "type"
func isValidType(t string) bool {
    return t == "Income" || t == "Outcome" || t == "+" || t == "-"
}

// Функця для приведения типов
func castToTypeFunc(value interface{}) (int, bool) {
    switch v := value.(type) {
    case int:
        return v, true
    case float64:
        if float64(int(v)) == v {
            return int(v), true
        }
    case string:
        var intValue int
        if _, err := fmt.Sscanf(v, "%d", &intValue); err == nil {
            return intValue, true
        }
    }
    return 0, false
}

// функция для обработки json файла
func handlerDataFunc(data []InputData) []OutData {
    company_res := groupAndProcessData(data)
    return prepareResult(company_res)
}

// функция группировки
func groupAndProcessData(data []InputData) map[string]*OutData {
    company_res := make(map[string]*OutData)
    
    for _, item := range data {
        createdAt := getCreationDate(item)
        if item.Company == "" || createdAt == "" {
            continue
        }

        if _, exists := company_res[item.Company]; !exists {
            company_res[item.Company] = &OutData{
                Company: item.Company,
            }
        }

        result := company_res[item.Company]
        processOperationItem(item, result)
    }
    
    return company_res
}

// функия получения даты
func getCreationDate(item InputData) string {
    if item.Operation != nil && item.Operation.CreatedAt != "" {
        return item.Operation.CreatedAt
    }
    return item.CreatedAt
}

// обработка операций
func processOperationItem(item InputData, result *OutData) {
    var operation_Type string
    var operation_Value int
    var operation_ID interface{}
    var valid_value bool

    if item.Operation != nil {
        operation_Type = item.Operation.Type
        operation_ID = item.Operation.ID
        operation_Value, valid_value = castToTypeFunc(item.Operation.Value)
    } else {
        operation_Type = item.Type
        operation_ID = item.ID
        operation_Value, valid_value = castToTypeFunc(item.Value)
    }

    if !isValidType(operation_Type) || !valid_value || operation_ID == nil {
        result.InvalidOperations = append(result.InvalidOperations, operation_ID)
        return
    }

    calculateBalance(result, operation_Type, operation_Value)
}

// посдсчет баланса, функция-счетчик
func calculateBalance(result *OutData, operation_Type string, operation_Value int) {
    if operation_Type == "Income" || operation_Type == "+" {
        result.Balance += operation_Value
    } else if operation_Type == "Outcome" || operation_Type == "-" {
        result.Balance -= operation_Value
    }
    result.ValidOperationsCount++
}

func prepareResult(company_res map[string]*OutData) []OutData {
    var result_data []OutData

    for _, result := range company_res {
        sort.Slice(result.InvalidOperations, func(i, j int) bool {
            return fmt.Sprintf("%v", result.InvalidOperations[i]) < fmt.Sprintf("%v", result.InvalidOperations[j])
        })
        result_data = append(result_data, *result)
    }

    sort.Slice(result_data, func(i, j int) bool {
        return result_data[i].Company < result_data[j].Company
    })

    return result_data
}

