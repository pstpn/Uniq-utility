package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Options struct {
	input           string
	output          string
	uniqParam       uint8
	iParam          bool
	fParamNumFields int
	sParamNumChars  int
}

func PrintIncorrectMsg() {
	color.Red("\n\nIncorrect args! \n\nUsage example:\n" +
		"uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]\n\n")

	color.Red("-c - count the number of occurrences of the string in the input. Output this number before the string separated by a space.\n\n" +
		"-d - output only those lines that are repeated in the input.\n\n" + "-u Output only those lines that are not repeated in the input.\n\n" +
		"-f num_fields - ignore the first num_fields of fields in a line.\n A field in a string is a non-empty set of characters separated by " +
		"a space.\n\n" + "-s num_chars - ignore the first num_chars characters in the string. \nWhen used with the -f option, the first characters " +
		"after num_fields fields are counted\n(ignoring the delimiter space after the last field).\n\n" + "-i - do not take into account the case of letters.\n\n")
}

func ParseOptions(opt *Options, params []string) bool {
	paramsCount := len(params)

	for i, curParam := range params {
		if curParam == "-c" ||
			curParam == "-d" ||
			curParam == "-u" {
			if opt.uniqParam != ' ' {
				return true
			}

			opt.uniqParam = curParam[1]
		} else if curParam == "-i" {
			if opt.iParam {
				return true
			}

			opt.iParam = true
		} else if curParam == "-f" {
			if i+1 >= paramsCount ||
				opt.fParamNumFields != 0 {
				return true
			}

			num, err := strconv.Atoi(params[i+1])
			if err != nil || num < 0 {
				return true
			}

			opt.fParamNumFields = num
		} else if curParam == "-s" {
			if i+1 >= paramsCount ||
				opt.sParamNumChars != 0 {
				return true
			}

			num, err := strconv.Atoi(params[i+1])
			if err != nil || num < 0 {
				return true
			}

			opt.sParamNumChars = num
		} else {
			_, err := strconv.Atoi(curParam)

			if err != nil && opt.input == "" {
				opt.input = curParam
			} else if err != nil && opt.output == "" {
				opt.output = curParam
			} else if err != nil {
				return true
			}

			if err == nil && (i == 0 || (params[i-1] != "-f" &&
				params[i-1] != "-s")) && opt.input == "" {
				opt.input = curParam
			} else if err == nil && (i == 0 || (params[i-1] != "-f" &&
				params[i-1] != "-s")) && opt.output == "" {
				opt.output = curParam
			} else if err == nil && (i == 0 || (params[i-1] != "-f" &&
				params[i-1] != "-s")) {
				return true
			}
		}
	}

	return false
}

func GetText(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var inStrings []string

	for scanner.Scan() {
		inStrings = append(inStrings, scanner.Text())
	}

	return inStrings
}

func getNewStartIndex(numFields int, numChars int, inStr string) int {
	newStart := 0
	wordsCount := 0
	strLen := len(inStr)

	for isSepSpace := false; newStart < strLen &&
		wordsCount != numFields; newStart++ {
		if inStr[newStart] != ' ' {
			isSepSpace = true
		} else if isSepSpace {
			wordsCount++
			isSepSpace = false
		}
	}

	for curShift := 0; newStart < strLen &&
		curShift != numChars; curShift++ {
		newStart++
	}

	return newStart
}

func UniqText(opt Options, inData []string) []string {
	outData := make(map[string]int)

	var outStrings []string

	for index, rootStr := range inData {
		curCount := 0

		if opt.iParam {
			rootStr = strings.ToLower(rootStr)
		}
		if opt.fParamNumFields != 0 || opt.sParamNumChars != 0 {
			rootStr = rootStr[getNewStartIndex(opt.fParamNumFields, opt.sParamNumChars, rootStr):]
		}

		for _, tmpStr := range inData {
			if opt.iParam {
				tmpStr = strings.ToLower(tmpStr)
			}

			if opt.fParamNumFields != 0 || opt.sParamNumChars != 0 {
				tmpStr = tmpStr[getNewStartIndex(opt.fParamNumFields, opt.sParamNumChars, tmpStr):]
			}

			if tmpStr == rootStr {
				curCount++
			}
		}

		if opt.uniqParam == 'c' &&
			opt.iParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, strconv.Itoa(curCount)+" "+inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.uniqParam == 'c' &&
			!opt.iParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, strconv.Itoa(curCount)+" "+inData[index])
			outData[inData[index]] = curCount
		}

		if opt.uniqParam == 'd' &&
			curCount > 1 &&
			opt.iParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.uniqParam == 'd' &&
			curCount > 1 &&
			!opt.iParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[inData[index]] = curCount
		}

		if opt.uniqParam == 'u' &&
			curCount == 1 &&
			opt.iParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.uniqParam == 'u' &&
			curCount == 1 &&
			!opt.iParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[inData[index]] = curCount
		}

		if opt.uniqParam == ' ' &&
			opt.iParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.uniqParam == ' ' &&
			!opt.iParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[inData[index]] = curCount
		}
	}

	return outStrings
}

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

func main() {
	args := os.Args[1:]

	opt := Options{
		"",
		"",
		' ',
		false,
		0,
		0,
	}

	if ParseOptions(&opt, args) {
		PrintIncorrectMsg()
		return
	}

	var inFile *os.File
	var err error

	if opt.input == "" {
		inFile, err = os.Stdin, nil
	} else {
		inFile, err = os.Open(opt.input)
	}
	if err != nil {
		PrintIncorrectMsg()
		return
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			return
		}
	}(inFile)

	inStrings := GetText(inFile)

	outStrings := UniqText(opt, inStrings)

	var outFile *os.File

	if opt.output == "" {
		outFile, err = os.Stdout, nil
	} else {
		outFile, err = os.Open(opt.output)
	}
	if err != nil {
		PrintIncorrectMsg()
		return
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			return
		}
	}(outFile)

	err = WriteText(outFile, outStrings)
}
