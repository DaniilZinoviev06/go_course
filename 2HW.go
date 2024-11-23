package main

import (
  "fmt"
  "flag" // для флага --flag
  "os"
  "io"
)

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
    
    fmt.Println("File name: ", file_path)
    fmt.Println(string(file_data))
}
