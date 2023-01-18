package operations

import (
	"myProject/pkgs/my_packages/project_types"
	"strconv"
	"strings"
)

func ParseOptions(opt *project_types.Options, params []string) bool {
	paramsCount := len(params)

	for i, curParam := range params {
		if curParam == "-c" ||
			curParam == "-d" ||
			curParam == "-u" {
			if opt.UniqParam != ' ' {
				return true
			}

			opt.UniqParam = curParam[1]
		} else if curParam == "-i" {
			if opt.IParam {
				return true
			}

			opt.IParam = true
		} else if curParam == "-f" {
			if i+1 >= paramsCount ||
				opt.FParamNumFields != 0 {
				return true
			}

			num, err := strconv.Atoi(params[i+1])
			if err != nil || num < 0 {
				return true
			}

			opt.FParamNumFields = num
		} else if curParam == "-s" {
			if i+1 >= paramsCount ||
				opt.SParamNumChars != 0 {
				return true
			}

			num, err := strconv.Atoi(params[i+1])
			if err != nil || num < 0 {
				return true
			}

			opt.SParamNumChars = num
		} else {
			_, err := strconv.Atoi(curParam)

			if err != nil && opt.Input == "" {
				opt.Input = curParam
			} else if err != nil && opt.Output == "" {
				opt.Output = curParam
			} else if err != nil {
				return true
			}

			if err == nil && (i == 0 || (params[i-1] != "-f" &&
				params[i-1] != "-s")) && opt.Input == "" {
				opt.Input = curParam
			} else if err == nil && (i == 0 || (params[i-1] != "-f" &&
				params[i-1] != "-s")) && opt.Output == "" {
				opt.Output = curParam
			} else if err == nil && (i == 0 || (params[i-1] != "-f" &&
				params[i-1] != "-s")) {
				return true
			}
		}
	}

	return false
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

func UniqText(opt project_types.Options, inData []string) []string {
	outData := make(map[string]int)

	var outStrings []string

	for index, rootStr := range inData {
		curCount := 0

		if opt.IParam {
			rootStr = strings.ToLower(rootStr)
		}
		if opt.FParamNumFields != 0 || opt.SParamNumChars != 0 {
			rootStr = rootStr[getNewStartIndex(opt.FParamNumFields, opt.SParamNumChars, rootStr):]
		}

		for _, tmpStr := range inData {
			if opt.IParam {
				tmpStr = strings.ToLower(tmpStr)
			}

			if opt.FParamNumFields != 0 || opt.SParamNumChars != 0 {
				tmpStr = tmpStr[getNewStartIndex(opt.FParamNumFields, opt.SParamNumChars, tmpStr):]
			}

			if tmpStr == rootStr {
				curCount++
			}
		}

		if opt.UniqParam == 'c' &&
			opt.IParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, strconv.Itoa(curCount)+" "+inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.UniqParam == 'c' &&
			!opt.IParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, strconv.Itoa(curCount)+" "+inData[index])
			outData[inData[index]] = curCount
		}

		if opt.UniqParam == 'd' &&
			curCount > 1 &&
			opt.IParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.UniqParam == 'd' &&
			curCount > 1 &&
			!opt.IParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[inData[index]] = curCount
		}

		if opt.UniqParam == 'u' &&
			curCount == 1 &&
			opt.IParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.UniqParam == 'u' &&
			curCount == 1 &&
			!opt.IParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[inData[index]] = curCount
		}

		if opt.UniqParam == ' ' &&
			opt.IParam &&
			outData[strings.ToLower(inData[index])] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[strings.ToLower(inData[index])] = curCount
		} else if opt.UniqParam == ' ' &&
			!opt.IParam &&
			outData[inData[index]] == 0 {
			outStrings = append(outStrings, inData[index])
			outData[inData[index]] = curCount
		}
	}

	return outStrings
}
