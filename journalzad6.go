package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range s.Grades {
		sum += g
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s Student) String() string {
	return fmt.Sprintf("%s: %v (ср. %.2f)", s.Name, s.Grades, s.Average())
}

type Journal struct {
	students map[string]Student
}

func (j Journal) AddStudent(name string, grades []int) {
	j.students[name] = Student{Name: name, Grades: grades}
}

func (j Journal) ShowAll() {
	if len(j.students) == 0 {
		fmt.Println("Нет студентов")
		return
	}
	for _, s := range j.students {
		fmt.Println(s)
	}
}

func (j Journal) FilterByAverage(threshold float64) {
	found := false
	for _, s := range j.students {
		if s.Average() < threshold {
			fmt.Println(s)
			found = true
		}
	}
	if !found {
		fmt.Println("Студентов с таким средним баллом нет")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	journal := Journal{students: make(map[string]Student)}

	for {
		fmt.Println("\n1. Добавить студента")
		fmt.Println("2. Показать всех")
		fmt.Println("3. Фильтр по среднему баллу")
		fmt.Println("4. Выход")
		fmt.Print("Выберите: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("ФИО: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Оценки через пробел: ")
			scanner.Scan()
			gradesInput := scanner.Text()

			var grades []int
			parts := strings.Fields(gradesInput)
			for _, p := range parts {
				if g, err := strconv.Atoi(p); err == nil && g >= 1 && g <= 5 {
					grades = append(grades, g)
				}
			}

			journal.AddStudent(name, grades)
			fmt.Printf("Студент %s добавлен\n", name)

		case "2":
			journal.ShowAll()

		case "3":
			fmt.Print("Порог среднего балла: ")
			scanner.Scan()
			thresholdStr := scanner.Text()

			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Ошибка: введите число!")
				continue
			}

			journal.FilterByAverage(threshold)

		case "4":
			return

		default:
			fmt.Println("Неверный выбор!")
		}
	}
}
