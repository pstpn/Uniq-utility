package in_out

import (
	"bufio"
	"github.com/fatih/color"
	"io"
)

// PrintIncorrectMsg - Функция вывода информационного сообщения при возникновении ошибки
func PrintIncorrectMsg() {

	color.Red("\n\nIncorrect args! \n\nUsage example:\n" +
		"uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]\n\n")

	color.Red("-c - count the number of occurrences of the string in the input. Output this number before the string separated by a space.\n\n" +
		"-d - output only those lines that are repeated in the input.\n\n" + "-u Output only those lines that are not repeated in the input.\n\n" +
		"-f num_fields - ignore the first num_fields of fields in a line.\n A field in a string is a non-empty set of characters separated by " +
		"a space.\n\n" + "-s num_chars - ignore the first num_chars characters in the string. \nWhen used with the -f option, the first characters " +
		"after num_fields fields are counted\n(ignoring the delimiter space after the last field).\n\n" + "-i - do not take into account the case of letters.\n\n")
}

// GetText - Функция, реализующая чтение входных данных (строк)
func GetText(r io.Reader) []string {

	scanner := bufio.NewScanner(r)

	var inStrings []string

	for scanner.Scan() {
		inStrings = append(inStrings, scanner.Text())
	}

	return inStrings
}

// WriteText - Функция, реализующая запись выходных данных (строк) в io.Writer
func WriteText(w io.Writer, data []string) error {
	
	var err error

	for _, curString := range data {
		_, err = w.Write([]byte(curString + "\n"))
		if err != nil {
			break
		}
	}

	return err
}
