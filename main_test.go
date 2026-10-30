package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		got := generateRandomElements(0)
		if len(got) != 0 {
			t.Fatalf("generateRandomElements(0) returned length %d, want 0", len(got))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		got := generateRandomElements(-10)
		if len(got) != 0 {
			t.Fatalf("generateRandomElements(-10) returned length %d, want 0", len(got))
		}
	})

	t.Run("positive size", func(t *testing.T) {
		const size = 100
		got := generateRandomElements(size)

		if len(got) != size {
			t.Fatalf("generateRandomElements(%d) returned length %d, want %d", size, len(got), size)
		}

		for i, v := range got {
			if v <= 0 {
				t.Fatalf("element %d is not positive: %d", i, v)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "multiple elements",
			data: []int{1, 5, 3, 9, 2},
			want: 9,
		},
		{
			name: "maximum at the end",
			data: []int{1, 2, 3, 4, 100},
			want: 100,
		},
		{
			name: "maximum at the beginning",
			data: []int{100, 4, 3, 2, 1},
			want: 100,
		},
		{
			name: "negative elements",
			data: []int{-5, -2, -10},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
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
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "less than chunks",
			data: []int{5, 1, 9, 3},
			want: 9,
		},
		{
			name: "multiple elements",
			data: []int{1, 5, 3, 9, 2, 8, 7, 6, 10},
			want: 10,
		},
		{
			name: "negative elements",
			data: []int{-5, -2, -10, -3, -4, -6, -7, -8, -9},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
