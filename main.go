package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

// Character представляет структуру информации о персонаже
type Character struct {
	Name      string         // Имя персонажа
	Age       string         // Возраст персонажа
	Childhood string         // Детство персонажа
	Adulthood string         // Зрелость персонажа
	Flaws     string         // Минусы персонажа
	Traits    string         // Особенности персонажа
	Skills    map[string]int // Навыки персонажа
	Health    string         // Здоровье персонажа
	Relations string         // Отношения персонажа
}

func main() {
	client := gosseract.NewClient() // Создание нового клиента для OCR
	defer client.Close()            // Закрытие клиента по завершении работы

	client.SetImage("/net/photo_2024-05-31_01-36-06.jpg") // Установка изображения для OCR
	text, err := client.Text()                            // Извлечение текста из изображения
	if err != nil {
		log.Fatal(err) // Обработка ошибки, если OCR не удалось
	}

	character := parseTextToCharacter(text) // Парсинг текста в структуру Character
	fmt.Printf("%+v\n", character)          // Вывод информации о персонаже
}

// parseTextToCharacter парсит извлеченный текст в структуру Character
func parseTextToCharacter(text string) Character {
	lines := strings.Split(text, "\n")                   // Разделение текста на строки
	character := Character{Skills: make(map[string]int)} // Инициализация структуры Character с картой навыков

	for _, line := range lines {
		if strings.Contains(line, "Предыстория") {
			character.Childhood = extractField(line, "Детство:")  // Извлечение детства
			character.Adulthood = extractField(line, "Зрелость:") // Извлечение зрелости
		} else if strings.Contains(line, "Минусы") {
			character.Flaws = extractField(line, "Минусы") // Извлечение минусов
		} else if strings.Contains(line, "Особенности") {
			character.Traits = extractField(line, "Особенности") // Извлечение особенностей
		} else if strings.Contains(line, "Здоровье") {
			character.Health = extractField(line, "Здоровье") // Извлечение информации о здоровье
		} else if strings.Contains(line, "Отношения") {
			character.Relations = extractField(line, "Отношения") // Извлечение отношений
		} else if strings.Contains(line, "навыки") {
			skillName := extractSkillName(line)      // Извлечение имени навыка
			skillValue := extractSkillValue(line)    // Извлечение значения навыка
			character.Skills[skillName] = skillValue // Добавление навыка в карту навыков
		} else if strings.Contains(line, "Он") {
			character.Name = extractName(line) // Извлечение имени персонажа
			character.Age = extractAge(line)   // Извлечение возраста персонажа
		}
	}

	return character
}

// extractField извлекает значение поля из строки
func extractField(line, fieldName string) string {
	parts := strings.Split(line, fieldName) // Разделение строки по имени поля
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1]) // Возврат значения поля
	}
	return ""
}

// extractSkillName извлекает имя навыка из строки
func extractSkillName(line string) string {
	parts := strings.Split(line, " ")
	return parts[0] // Возврат имени навыка
}

// extractSkillValue извлекает значение навыка из строки
func extractSkillValue(line string) int {
	parts := strings.Split(line, " ")
	value := 0
	if len(parts) > 1 {
		fmt.Sscanf(parts[len(parts)-1], "%d", &value) // Парсинг значения навыка
	}
	return value
}

// extractName извлекает имя персонажа из строки
func extractName(line string) string {
	parts := strings.Split(line, " ")
	if len(parts) > 1 {
		return parts[1] // Возврат имени персонажа
	}
	return ""
}

// extractAge извлекает возраст персонажа из строки
func extractAge(line string) string {
	parts := strings.Split(line, "возраст")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1]) // Возврат возраста персонажа
	}
	return ""
}
