package main

import (
	"crypto/rand"
	"os"
	"testing"
)

func BenchmarkCopyWithIoCopy(b *testing.B) {
	src := createTestFile(b, 1024*1024) // 1MB тестовый файл
	defer os.Remove(src)
	dst := src + "_copy"
	defer os.Remove(dst)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		os.Remove(dst) // Удаляем копию перед каждым запуском
		if err := CopyWithIoCopy(src, dst); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCopyWithReadAll(b *testing.B) {
	src := createTestFile(b, 1024*1024)
	defer os.Remove(src)
	dst := src + "_copy"
	defer os.Remove(dst)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		os.Remove(dst)
		if err := CopyWithReadAll(src, dst); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCopyWithBuffer(b *testing.B) {
	src := createTestFile(b, 1024*1024)
	defer os.Remove(src)
	dst := src + "_copy"
	defer os.Remove(dst)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		os.Remove(dst)
		if err := CopyWithBuffer(src, dst); err != nil {
			b.Fatal(err)
		}
	}
}

// Вспомогательная функция для создания тестового файла с случайными данными
func createTestFile(b *testing.B, size int) string {
	data := make([]byte, size)
	if _, err := rand.Read(data); err != nil {
		b.Fatal(err)
	}
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		b.Fatal(err)
	}
	return file.Name()
}
