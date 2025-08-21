package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var ccb = color.New(color.FgBlue).SprintFunc()
	println(ccb("Здравствуйте это пробная версия которая демонстрирует какой язык программирования или машинный код можно написать новичку"))
	var s_1 string
	var s_2 string
	// vari для цвета
	var ccr = color.New(color.FgRed).SprintFunc()
	var ccby = color.New(color.FgBlue, color.Bold).SprintFunc()
	var ccg = color.New(color.FgGreen).SprintFunc()
	// vari для программы
	var j string = "код успешно выполнен"
	var k string
	// приветствие
	println(ccb("можно сделать принт что бы это сделать напишите и иметь не более 2 логических строк"))
	println(ccb("программа может иметь до 2 переменных обьявление сначала vare потом новая строка название потом значение"))
	println(ccb("print"))
	println(ccb("или напишите"))
	println(ccb("no"))
	// задаем стринги для переменных
	var a11 string
	var a12 string

	// выбираем тип строки
	// основная логика общения
	fmt.Scan(&s_1)
	if s_1 == "print" {
		println(ccb("тогда напишите ? что хотите запринтить(если переменную то пиши ; название переменной"))
		var s_1_1 string
		var s_1_2 string
		fmt.Scan(&s_1_1, &s_1_2)
		if s_1_1 == "?" {
			println(ccb("хотите скомпилировать ваш код yes или no"))
			var k string
			fmt.Scan(&k)
			if k == "yes" {
				println(ccg(s_1_2)) // Использовал s_1_2, так как это второе введённое значение
				println(ccby(j))
			} else if k == "no" {
				if s_2 == "vare" {
					var s_2_2 string
					fmt.Scan(&s_2_2)
					fmt.Scan(&a11, &a12)
					fmt.Scan(&k)
					if k == "yes" {
						println(ccg(s_1_2))
						println(ccby(j))
					} else if k == "no" {
						var s_3 string
						fmt.Scan(&s_3)
						if s_3 == "print" {
							var s_3_1 string
							var s_3_2 string
							fmt.Scan(&s_3_1, &s_3_2)
							if s_3_1 == "?" {
								println(ccg(s_1_1))
								println(ccg(s_3_2))
								println(ccby(j))
							} else {
								println(ccr("ошибка компиляции"))
							}
						} else if s_3 == "vare" {
							println(ccg(s_1_1))
							println(ccby(j))
						} else {
							println(ccr("ошибка компиляции"))
						}
					}
				}
			}
		} else {
			println(ccr("ошибка компиляции"))
		}
	} else if s_1 == "vare" {
		fmt.Scan(&a11, &a12)
		fmt.Scan(&k)
		if k == "no" {
			fmt.Scan(&s_2)
			if s_2 == "print" {
				var s_2_1 string
				var s_2_2 string
				fmt.Scan(&s_2_1, &s_2_2)
				if s_2_1 == "?" {
					println(ccg(s_2_2))
					println(ccby(j))
				} else if s_2_1 == ";" {
					if s_2_2 == a11 {
						println(ccg(a12))
						println(ccby(j))
					} else {
						println(ccr("ошибка компиляции"))
					}
				}
			} else {
				println(ccr("ошибка компиляции"))
			}
		} else if k == "yes" {
			println(ccby(j))
		} else {
			println(ccr("ошибка компиляции"))
		}
	} else {
		println(ccr("ошибка компиляции"))
	}
}
