package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

// функция принимает срез с интами и вариативный параметр для управления обработкой среза
func sliceProcessingFunc(numbers []int, settings ...string) []int {
	// Параметры для обработки слайса по умолчанию
	minV := 0
	sort_numbers := false
	del_duplicates := true
	// массив для проверки, что пользователь правильно написал название параметра
	const_settings := [3]string{"sort_numbers", "minV", "del_duplicates"}
	for i := 0; i < len(settings); i++ {
		// разделяем строку на название параметра и значение через "="
		key_value := strings.Split(settings[i], "=")
		//fmt.Println(settings[i])
		//fmt.Println(key_value[0], key_value[1])
		if len(key_value) == 2 {
			count := 0
			for j := 0; j < len(const_settings); j++ { 
				if key_value[0] != const_settings[j] {
					count++
				}
				if count == len(const_settings) {
					fmt.Println("Неправильное название параметра - " + key_value[0])
				}
			}
			// Задаем пользовательские настройки
			switch key_value[0] {
           		case "sort_numbers":
             		if key_value[1] == "true" {
				sort_numbers = true
             		} 
           		case "del_duplicates":
             			if key_value[1] == "false" {
              				del_duplicates = false
             			}
           		case "minV":
             			v, err := strconv.Atoi(key_value[1])
             			if err == nil {
                			minV = v
             			}
             		}           
     		} else {
     			fmt.Println("Некорректный параметр.")
     		}
     	}
     	
     	// цикл для удаление дубликатов
	if del_duplicates {
		for i := 0; i < len(numbers); i++ {
			for j := i+1; j < len(numbers); j++ {
				if numbers[i] == numbers[j] {
					numbers = append(numbers[:j], numbers[j+1:]...)
					j--
				}
			}
		}
	}
	
	// фильтрация по минимальному значению
	result_numbers := []int{}
	for i := 0; i < len(numbers); i++ {
        	if numbers[i] > minV {
        		result_numbers = append(result_numbers, numbers[i])
        	}
        }
        
        // сортировка
    	if sort_numbers {
    		sort.Ints(result_numbers)
    	}
    
    	fmt.Println(result_numbers)
	return result_numbers
}

func main(){
	numbers := []int{12, 4, 123, 3, 2, 34, 1, 0, 3, 234}
	sliceProcessingFunc(numbers, "sort_numbers=true", "minV=2", "del_duplicates=true")
}
