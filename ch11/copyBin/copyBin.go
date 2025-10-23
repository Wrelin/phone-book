package main

import (
	"io"
	"os"
)

// CopyWithIoCopy : Использует io.Copy — стандартный и эффективный способ.
// Копирует данные с внутренним буфером (обычно 32KB), минимизируя системные вызовы.
func CopyWithIoCopy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// CopyWithReadAll : Читает весь файл в память, затем записывает.
// Подходит для небольших файлов, но может быть медленнее и требовать больше памяти для больших.
func CopyWithReadAll(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

// CopyWithBuffer : Ручное буферизованное копирование с фиксированным буфером (4KB).
// Позволяет контролировать размер буфера, но требует больше кода и может быть менее оптимизированным, чем io.Copy.
func CopyWithBuffer(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 32768) // 32KB буфер
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := dstFile.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
