package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Problem10B struct {

}




func (this *Problem10B) Solve() {
	Log.Info("Problem 10B solver beginning!")

	file, err := os.Open("source-data/input-day-10b.txt");
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
	inCompleteLines := make([]*SyntaxLine, 0);
	for _, line := range syntaxLines {
		stack := make([]rune, 0);
		corruptionFound := false;
		for _, v := range line.Runes {
			isClosing := false;
			isMatch := false;
			if(v == ')'){
				isClosing = true;
				if(stack[len(stack) - 1] == '('){
					isMatch = true;
				} else{
					isMatch = false;
				}
			}
			if(v == ']'){
				isClosing = true;
				if(stack[len(stack) - 1] == '['){
					isMatch = true;
				} else{
					isMatch = false;
				}
			}
			if(v == '}'){
				isClosing = true;
				if(stack[len(stack) - 1] == '{'){
					isMatch = true;
				} else{
					isMatch = false;
				}
			}
			if(v == '>'){
				isClosing = true;
				if(stack[len(stack) - 1] == '<'){
					isMatch = true;
				} else{
					isMatch = false;
				}
			}
			if(isClosing){
				if(isMatch){
					stack = stack[:len(stack)-1]
				} else{
					//Log.Info("Found corruption at line %d, points %d", lineNum, points);
					//corruptionPoints += points;
					corruptionFound = true;
					break;
				}
			} else{
				stack = append(stack, v);
			}
		}
		if(len(stack) > 0 && !corruptionFound){
			//Log.Info("Incomplete line at %d", lineNum)
			line.LastStackContents = stack;
			inCompleteLines = append(inCompleteLines, line);

		}
		lineNum++;
	}
	Log.Info("Completed analysis of %d lines, incomplete lines %d", len(syntaxLines), len(inCompleteLines));
	scores := make([]int, 0);
	for _, v := range inCompleteLines{
		score := 0;
		for {
			if(len(v.LastStackContents) <= 0){
				break;
			}
			score = score * 5;
			last := v.LastStackContents[len(v.LastStackContents) - 1]
			if(last == '('){
				score += 1;
			}
			if(last == '['){
				score += 2;
			}
			if(last == '{'){
				score += 3;
			}
			if(last == '<'){
				score += 4;
			}

			v.LastStackContents = v.LastStackContents[:len(v.LastStackContents)-1]
		}
		//Log.Info("%d", score);
		scores = append(scores, score);
	}

	sort.Ints(scores);
	luckyScore := scores[(len(scores)/2)];
	Log.Info("Completed analysis. %d incomplete lines, lucky score %d", len(inCompleteLines), luckyScore)
}
