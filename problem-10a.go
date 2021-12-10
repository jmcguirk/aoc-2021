package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem10A struct {

}


type SyntaxLine struct {
	Runes []rune;
	LastStackContents []rune;
}


func (this *Problem10A) Solve() {
	Log.Info("Problem 10A solver beginning!")

	file, err := os.Open("source-data/input-day-10a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	syntaxLines := make([]*SyntaxLine, 0);

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			syntaxLine := &SyntaxLine{};
			syntaxLine.Runes = make([]rune, 0);
			for _, v := range line {
				syntaxLine.Runes = append(syntaxLine.Runes, v);
			}
			syntaxLines = append(syntaxLines, syntaxLine);
		}
	}

	Log.Info("Parsed %d syntax lines", len(syntaxLines));
	lineNum := 0;
	corruptionPoints := 0;
	for _, line := range syntaxLines {
		stack := make([]rune, 0);
		corruptionFound := false;
		for _, v := range line.Runes {
			isClosing := false;
			isMatch := false;
			points := 0;
			if(v == ')'){
				isClosing = true;
				if(stack[len(stack) - 1] == '('){
					isMatch = true;
				} else{
					isMatch = false;
					points = 3;
				}
			}
			if(v == ']'){
				isClosing = true;
				if(stack[len(stack) - 1] == '['){
					isMatch = true;
				} else{
					isMatch = false;
					points = 57;
				}
			}
			if(v == '}'){
				isClosing = true;
				if(stack[len(stack) - 1] == '{'){
					isMatch = true;
				} else{
					isMatch = false;
					points = 1197;
				}
			}
			if(v == '>'){
				isClosing = true;
				if(stack[len(stack) - 1] == '<'){
					isMatch = true;
				} else{
					isMatch = false;
					points = 25137;
				}
			}
			if(isClosing){
				if(isMatch){
					stack = stack[:len(stack)-1]
				} else{
					Log.Info("Found corruption at line %d, points %d", lineNum, points);
					corruptionPoints += points;
					corruptionFound = true;
					break;
				}
			} else{
				stack = append(stack, v);
			}
		}
		if(len(stack) > 0 && !corruptionFound){
			//Log.Info("Incomplete line at %d", lineNum)
		}
		lineNum++;
	}
	Log.Info("Completed analysis of %d lines, total corruption points %d", len(syntaxLines), corruptionPoints);
}
