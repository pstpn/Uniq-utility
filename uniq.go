package main

import (
	"myProject/pkgs/my_packages/in_out"
	"myProject/pkgs/my_packages/operations"
	"myProject/pkgs/my_packages/project_types"
	"os"
)

func main() {
	args := os.Args[1:]

	opt := project_types.Options{
		UniqParam: ' ',
	}

	if operations.ParseOptions(&opt, args) {
		in_out.PrintIncorrectMsg()
		return
	}

	var inFile *os.File
	var err error

	if opt.Input == "" {
		inFile, err = os.Stdin, nil
	} else {
		inFile, err = os.Open(opt.Input)
	}
	if err != nil {
		in_out.PrintIncorrectMsg()
		return
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			return
		}
	}(inFile)

	inStrings := in_out.GetText(inFile)

	outStrings := operations.UniqText(opt, inStrings)

	var outFile *os.File

	if opt.Output == "" {
		outFile, err = os.Stdout, nil
	} else {
		outFile, err = os.Open(opt.Output)
	}
	if err != nil {
		in_out.PrintIncorrectMsg()
		return
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			return
		}
	}(outFile)

	err = in_out.WriteText(outFile, outStrings)
}
