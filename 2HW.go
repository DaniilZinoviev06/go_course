package main

import (
  "fmt"
  "flag" // для флага --file
  "os"
  "io"
  "encoding/json"
)

// ищем путь к файлу по приоритету
func getFilePathSource(flag_path string) string {
    if flag_path != "" {
        return flag_path
    }

    env_path := os.Getenv("ENV_FILE")
    if env_path != "" {
        return env_path
    }

    stdin_info, err := os.Stdin.Stat()
    if err == nil && (stdin_info.Mode()&os.ModeCharDevice == 0) {
        tmp_file, err := os.CreateTemp("", "HW2-*.json")
        if err != nil {
            fmt.Println(err)
            return ""
        }
        defer tmp_file.Close()
    
        stdin_data, err := io.ReadAll(os.Stdin)
        if err != nil {
            fmt.Println(err)
            return ""
        }
    
        _, err = tmp_file.Write(stdin_data)
        if err != nil {
            fmt.Println(err)
            return ""
        }
    
        return tmp_file.Name()
    }
    
    return ""
}

func main() {
    file_flag := flag.String("file", "", "Flag for JSON file")
    flag.Parse()
    
    file_path := getFilePathSource(*file_flag)
    
    if file_path == "" {
        fmt.Println("Не указан файл.")
        return
    }
    
    file, err := os.Open(file_path)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    
    file_data, err := io.ReadAll(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    var input_data []InputData
    if err := json.Unmarshal(file_data, &input_data); err != nil {
        fmt.Printf("Ошибка при парсинге JSON: %v\n", err)
        return
    }
    
    fmt.Println("Имя файла: ", file_path)
    
    results := handlerDataFunc(input_data)

    output_file, err := os.Create("out.json")
    if err != nil {
        fmt.Println("Не удалось создать выходной файл:", err)
        return
    }
    defer output_file.Close()

    encoder := json.NewEncoder(output_file)
    encoder.SetIndent("", "\t")
    if err := encoder.Encode(results); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Обработка успешно завершена!")
    //fmt.Println(string(file_data))
}
