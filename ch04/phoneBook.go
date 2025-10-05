package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Tel        string
	LastAccess string
}

type EntryWithSurname struct {
	Entry
	Surname string
}

type EntryInterface interface {
	GetName() string
	GetTel() string
	GetLastAccess() string
	SetLastAccess(string)
}

type SurnameInterface interface {
	EntryInterface
	GetSurname() string // Расширение для типов с Surname
}

func (e Entry) GetName() string                 { return e.Name }
func (e Entry) GetTel() string                  { return e.Tel }
func (e Entry) GetLastAccess() string           { return e.LastAccess }
func (e Entry) SetLastAccess(LastAccess string) { e.LastAccess = LastAccess }

func (e EntryWithSurname) GetSurname() string { return e.Surname }

type PhoneBook[T EntryInterface] []T

var data PhoneBook[EntryInterface]

// var dataWithSurname PhoneBook[EntryWithSurname]
var index map[string]int

// CsvFile resides in the home directory of the current user
var CsvFile = "/tmp/phonebook.csv"

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// CSV file read all at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		var temp EntryInterface = Entry{
			Name:       line[0],
			Tel:        line[1],
			LastAccess: line[2],
		}

		if len(line) == 4 {
			temp = EntryWithSurname{
				Entry:   temp.(Entry),
				Surname: line[3],
			}
		}

		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filepath string) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	for _, row := range data {
		temp := []string{row.GetName(), row.GetTel(), row.GetLastAccess()}
		ews, ok := row.(SurnameInterface)
		if ok {
			temp = append(temp, ews.GetSurname())
		}
		_ = csvWriter.Write(temp)
	}

	csvWriter.Flush()
	return nil
}

func createIndex() error {
	index = make(map[string]int)

	for i, k := range data {
		key := k.GetTel()
		index[key] = i
	}

	return nil
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func initEntry(N, T string) *Entry {
	// Both of them should have a value
	if T == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Tel: T, LastAccess: LastAccess}
}

func initEntryWithSurname(N, S, T string) *EntryWithSurname {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}
	return &EntryWithSurname{
		Entry:   *initEntry(N, T),
		Surname: S,
	}
}

func insert(e EntryInterface) error {
	// If it already exists, do not add it
	_, ok := index[e.GetTel()]
	if ok {
		return fmt.Errorf("%s already exists", e.GetTel())
	}
	data = append(data, e)
	// Update the index
	_ = createIndex()

	err := saveCSVFile(CsvFile)
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}

	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist anymore
	delete(index, key)

	err := saveCSVFile(CsvFile)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) EntryInterface {
	i, ok := index[key]
	if !ok {
		return nil
	}

	la := strconv.FormatInt(time.Now().Unix(), 10)
	data[i].SetLastAccess(la)
	return data[i]
}

func list(isReverse bool) {
	if isReverse {
		sort.Sort(sort.Reverse(data))
	} else {
		sort.Sort(data)
	}

	for _, v := range data {
		fmt.Println(v)
	}
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func setCsvFile() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		CsvFile = filepath
	}

	_, err := os.Stat(CsvFile)
	if err != nil {
		fmt.Println("Creating", CsvFile)
		f, err := os.Create(CsvFile)
		if err != nil {
			return err
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CsvFile)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", CsvFile)
	}
	return nil
}

func (a PhoneBook[T]) Len() int {
	return len(a)
}

func (a PhoneBook[T]) Less(i, j int) bool {
	switch c := any(a[i]).(type) {
	case EntryWithSurname:
		d := any(a[j]).(EntryWithSurname)
		if c.GetSurname() != d.GetSurname() {
			return c.GetSurname() < d.GetSurname()
		}
		return c.GetName() < d.GetName()
	case Entry:
		d := any(a[j]).(Entry)
		return c.GetName() < d.GetName()
	}
	return false
}

func (a PhoneBook[T]) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	err := setCsvFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readCSVFile(CsvFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "insert":
		if len(arguments) < 4 {
			fmt.Println("Usage: insert Name Telephone Surname")
			return
		}
		t := strings.ReplaceAll(arguments[3], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		if len(data) > 0 {
			switch data[0].(type) {
			case Entry:
				if len(arguments) != 4 {
					fmt.Println("Not valid argument count, already use 3, current", len(arguments)-1)
					return
				}
			case EntryWithSurname:
				if len(arguments) < 5 {
					fmt.Println("Not valid argument count, already use 4, current", len(arguments)-1)
					return
				}
			}
		}

		var err error
		var temp EntryInterface
		if len(arguments) == 4 {
			temp = initEntry(arguments[2], t)
		} else {
			temp = initEntryWithSurname(arguments[2], arguments[4], t)
		}

		// If it was nil, there was an error
		if temp != nil {
			err = insert(temp)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(temp)
	case "list":
		isReverse := len(arguments) >= 3 && arguments[2] == "reverse"
		list(isReverse)
	default:
		fmt.Println("Not a valid option")
	}
}
