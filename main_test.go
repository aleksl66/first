package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"размер 1", 1, 1},
		{"размер 100", 100, 100},
		{"размер 1000", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("generateRandomElements(%d) вернул слайс длиной %d, ожидалось %d",
					tt.size, len(got), tt.want)
			}
			for _, v := range got {
				if v <= 0 {
					t.Errorf("generateRandomElements(%d) содержит неположительное число: %d", tt.size, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"два элемента", []int{5, 9}, 9},
		{"два элемента по убыванию", []int{9, 5}, 9},
		{"все одинаковые", []int{7, 7, 7}, 7},
		{"максимум в середине", []int{1, 100, 3}, 100},
		{"максимум в конце", []int{1, 2, 3, 99}, 99},
		{"большой слайс", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"меньше 8 элементов", []int{1, 9, 3}, 9},
		{"ровно 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 100}, 100},
		{"больше 8, не делится нацело", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"больше 8, максимум в остатке", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 777}, 777},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

// Пишите тесты в этом файле
