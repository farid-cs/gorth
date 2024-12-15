package main

import (
	"fmt"
	"strconv"
	"regexp"
	"os"
) 

func readfile(path string) []string {

	var raw []string;
	reg := regexp.MustCompile(" |\n")

	dat, err := os.ReadFile(path);
	if (err != nil) {
		panic("failed to read file");
	}

	/* FIXME: Find a better solution */
	tmp := reg.Split(string(dat), -1);

	for i := range tmp{
		if tmp[i] != "" {
			raw = append(raw, tmp[i]);
		} 
	};

	return raw;
}

func interpret(raw[]string) {
	var err error;
	var number int;
	var stack[]int;

	for i := 0; i < len(raw); i++ {
		length := len(stack);
		number, err = strconv.Atoi(raw[i]);
		if err == nil {
			stack = append(stack, number);
			continue; 
		}

		switch raw[i] {
		case "+":
			if length < 2 {
				panic("failed to plus: stack underflow");
			}
			stack[length - 2] = stack[length - 1] + stack[length - 2]
			stack = stack[:length - 1];
		case "-":
			if length < 2 {
				panic("failed to minus: stack underflow");
			}
			stack[length - 2] = stack[length - 1] - stack[length - 2]
			stack = stack[:length - 1];
		case ".":
			if length < 1 {
				panic("failed to dump: stack underflow");
			}
			fmt.Println(stack[length - 1]);
			stack = stack[:length - 1];
		default:
			panic("invalid word");
		}
	} 
}

func main() {
	var file string;
	argv := os.Args;
	argc := len(argv);
	compile_flg := false;

	switch argc {
	case 3:
		if argv[1] != "-S" {
			panic("Invalid usage");
		}
		compile_flg = true;
		fallthrough;
	case 2:
		file = argv[argc - 1];
	default:
		panic("Invaid usage");
	}

	output := readfile(file);
	if (compile_flg) {
		panic("Not Implemented");
	} else {
		interpret(output);
	}
}