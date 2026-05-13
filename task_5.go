package main

import (
	"fmt"
	"time"
)

// Описание структуры Employee (Работник)
type Employee struct {
	Name     string    // Имя работника
	Position string    // Должность
	Salary   float64   // Зарплата
	HireDate time.Time // Дата приема на работу
}

// ==========================================
// КОНСТРУКТОР
// ==========================================
// Функция-конструктор, которая возвращает указатель на новую структуру
func NewEmployee(name, position string, salary float64, hireDate time.Time) *Employee {
	return &Employee{
		Name:     name,
		Position: position,
		Salary:   salary,
		HireDate: hireDate,
	}
}

// ==========================================
// МЕТОДЫ СТРУКТУРЫ
// ==========================================

// Метод 1: Вывод полной информации о работнике
func (e *Employee) PrintInfo() {
	// e.HireDate.Format позволяет вывести дату в привычном виде: День.Месяц.Год
	fmt.Printf("Сотрудник: %s\nДолжность: %s\nЗарплата: %.2f руб.\nПринят на работу: %s\n",
		e.Name, e.Position, e.Salary, e.HireDate.Format("02.01.2006"))
	fmt.Println("-------------------------")
}

// Метод 2: Повышение зарплаты (изменяет данные структуры)
func (e *Employee) GiveRaise(amount float64) {
	if amount > 0 {
		e.Salary += amount
		fmt.Printf("Успех! Зарплата сотрудника %s повышена на %.2f руб.\n", e.Name, amount)
		fmt.Println("-------------------------")
	}
}

// Метод 3: Расчет стажа (сколько полных лет человек работает)
func (e *Employee) GetYearsOfExperience() int {
	now := time.Now()
	years := now.Year() - e.HireDate.Year()

	// Если в текущем году "годовщина" работы еще не наступила, отнимаем один год
	if now.YearDay() < e.HireDate.YearDay() {
		years--
	}
	return years
}

// ==========================================
// ОСНОВНАЯ ФУНКЦИЯ ДЛЯ ПРОВЕРКИ
// ==========================================
func main() {
	// 1. Задаем дату приема на работу (год, месяц, день, часы, минуты, секунды, наносекунды, часовой пояс)
	dateOfHire := time.Date(2020, time.March, 15, 0, 0, 0, 0, time.UTC)

	// 2. Создаем работника через наш конструктор
	emp := NewEmployee("Алексей Смирнов", "Backend-разработчик", 150000.0, dateOfHire)

	// 3. Вызываем Метод 1 (Вывод информации)
	emp.PrintInfo()

	// 4. Вызываем Метод 3 (Расчет стажа)
	fmt.Printf("Текущий стаж работы: %d лет\n", emp.GetYearsOfExperience())
	fmt.Println("-------------------------")

	// 5. Вызываем Метод 2 (Повышение зарплаты)
	emp.GiveRaise(25000.0)

	// Снова выводим информацию, чтобы убедиться, что зарплата изменилась
	emp.PrintInfo()
}
