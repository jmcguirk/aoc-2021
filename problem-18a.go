package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem18A struct {

}

type SnailStatement struct{
	RawStatement string;
	CurrentStatement string;

	LeftLiteralValue int;
	RightLiteralValue int;
	LeftStatement *SnailStatement;
	RightStatement *SnailStatement;
}

type SnailStatementParser struct{
	Line string;
	Position int;
}

func (this *SnailStatement) Magnitude() int {
	leftValue := this.LeftLiteralValue;
	if(this.LeftStatement != nil){
		leftValue = this.LeftStatement.Magnitude();
	}
	rightValue := this.RightLiteralValue;
	if(this.RightStatement != nil){
		rightValue = this.RightStatement.Magnitude();
	}
	return (3*leftValue) + (2*rightValue);
}


func (this *SnailStatementParser) Parse() *SnailStatement {
	curr := this.Line[this.Position];
	if(curr != '['){
		Log.Fatal("Unexpected token when parsing snail statement");
	}
	res := &SnailStatement{};
	this.Position++;
	next := this.Line[this.Position];
	if(next == '['){
		res.LeftStatement = this.Parse();
	} else{
		literalBuffer := "";
		for{
			next = this.Line[this.Position];
			if(next == ','){
				break;
			}
			literalBuffer += fmt.Sprintf("%c", this.Line[this.Position]);
			this.Position++;
		}
		var err error;
		//Log.Info("Parse left " + literalBuffer);
		res.LeftLiteralValue, err = strconv.Atoi(literalBuffer);
		if(err != nil){
			Log.FatalError(err);
		}
	}
	next = this.Line[this.Position];
	if(next == ','){
		this.Position++;
	}
	next = this.Line[this.Position];
	if(next == '['){
		res.RightStatement = this.Parse();
		next = this.Line[this.Position];
		if(next == ']'){
			this.Position++;
		}
	} else{
		//Log.Info("Begin parse right literal %c", next);
		literalBuffer := "";
		for{
			next = this.Line[this.Position];
			if(next == ']'){
				this.Position++;
				break;
			}
			literalBuffer += fmt.Sprintf("%c", this.Line[this.Position]);
			this.Position++;
		}
		var err error;
		//Log.Info("Parse right " + literalBuffer);
		res.RightLiteralValue, err = strconv.Atoi(literalBuffer);
		if(err != nil){
			Log.Info("L2");
			Log.FatalError(err);
		}
	}

	return res;
}

func (this *SnailStatement) Reduce() {
	for{
		if(this.Explode()){
			continue;
		}
		if(this.Split()){
			continue;
		}
		break;
	}
}


func (this *SnailStatement) Add(that *SnailStatement) *SnailStatement {
	res := &SnailStatement{};
	res.RawStatement = that.CurrentStatement;
	if(this.CurrentStatement != ""){
		res.RawStatement = fmt.Sprintf("[%s,%s]", this.CurrentStatement, that.CurrentStatement);
	}
	res.CurrentStatement = res.RawStatement;
	return res;
}

func (this *SnailStatement) Split() bool {
	splitStart := -1;
	splitBuff := "";
	splitEnd := -1;
	splitValue := -1;
	for i, v := range this.CurrentStatement{
		isTerminal := false;
		if(v == '['){
			isTerminal = true;
		} else if(v == ']'){
			isTerminal = true;
		} else if(v == ','){
			isTerminal = true;
		}

		if(isTerminal){
			if(splitBuff != ""){
				parsed, err := strconv.Atoi(splitBuff);
				if(err != nil){
					Log.Info("X2");
					Log.FatalError(err);
				}
				if(parsed >= 10){
					splitValue = parsed;
					break;
				}
			}
			splitBuff = "";
			continue;
		}

		if(splitBuff == ""){
			splitStart = i;
			splitEnd = i;
		} else{
			splitEnd = i;
		}
		splitBuff += fmt.Sprintf("%c", v);
	}
	if(splitValue > 0){
		//Log.Info("Found split value %d at index %d:%d", splitValue, splitStart, splitEnd)
		splitLeft := splitValue/2;
		splitRight := splitValue/2 + splitValue % 2;
		newSliced := this.CurrentStatement[0:splitStart];
		newSliced += fmt.Sprintf("[%d,%d]", splitLeft, splitRight);
		newSliced += this.CurrentStatement[splitEnd+1:];
		this.CurrentStatement = newSliced;
		return true;
	}
	return false;
}


func (this *SnailStatement) Explode() bool {

	depth := 0;
	pivot := -1;
	for i, v := range this.CurrentStatement{
		if(v == '['){
			depth++;
			if(depth >= 5){
				simpleStatement := false;
				for j := i+1; j < len(this.CurrentStatement); j++{
					if(this.CurrentStatement[j] == '['){
						break;
					}
					if(this.CurrentStatement[j] == ']'){
						simpleStatement = true;
						break;
					}
				}
				if(simpleStatement){
					pivot = i+1;
					break;
				}
			}
			continue;
		} else if(v == ']'){
			depth--;
			continue;
		} else if(v == ','){
			continue;
		}
	}
	if(pivot < 0){
		return false;
	}
	//Log.Info("Discovered pivot at index %d:%c", pivot, this.CurrentStatement[pivot]);

	luckyPairString := "";
	i := pivot;
	closePivot := 0;
	for{
		if(this.CurrentStatement[i] == ']'){
			closePivot = i;
			break;
		}
		luckyPairString += fmt.Sprintf("%c", this.CurrentStatement[i]);
		i++;
	}


	pairParts := strings.Split(luckyPairString, ",");
	left, err := strconv.Atoi(pairParts[0]);
	if(err != nil){
		Log.Info("X3");
		Log.FatalError(err);
	}
	right, err := strconv.Atoi(pairParts[1]);
	if(err != nil){
		Log.Info("X4" + luckyPairString);
		Log.FatalError(err);
	}

	//Log.Info("Extracted lucky pair %d,%d in pivots %d-%d", left, right, pivot, closePivot);

	sliced := this.CurrentStatement[0:pivot-1];
	scanStartLeft := len(sliced);
	sliced += "0";
	scanRightStart := len(sliced);
	sliced += this.CurrentStatement[closePivot+1:];
	//Log.Info("First stage slice complete %s", sliced);

	i = scanRightStart;
	rightExplodePivot := -1;
	for {
		i++;
		if(i >= len(sliced)){
			break;
		}
		if(sliced[i] == '['){
			continue;
		} else if(sliced[i] == ']'){
			continue;
		} else if(sliced[i] == ','){
			continue;
		}
		rightExplodePivot = i;
		break;
	}

	if(rightExplodePivot > 0){
		//Log.Info("Found number right %c", sliced[i]);

		rightNumberString := fmt.Sprintf("%c", sliced[i]);

		i = rightExplodePivot;
		rightExplodePivotEnd := 0;
		for {
			i++;
			rightExplodePivotEnd = i;
			if (i >= len(sliced)) {
				break;
			}
			if (sliced[i] == '[') {
				break;
			} else if (sliced[i] == ']') {
				break;
			} else if (sliced[i] == ',') {
				break;
			}
			rightNumberString += fmt.Sprintf("%c", sliced[i]);
		}

		//Log.Info("Extracted right number string %s", rightNumberString);
		rightNumberParsed, err := strconv.Atoi(rightNumberString);
		if(err != nil){
			Log.FatalError(err);
		}

		rightNumberParsed += right;
		//Log.Info("Incremented value now %d", rightNumberParsed);
		newSliced := sliced[0:rightExplodePivot];
		newSliced += fmt.Sprintf("%d", rightNumberParsed);
		newSliced += sliced[rightExplodePivotEnd:];
		//Log.Info("Substituted string is %s", newSliced)
		sliced = newSliced;
	}

	scanStartLeftStart := -1;
	scanStartLeftEnd := -1;
	i = scanStartLeft;
	scanLeftFoundNumber := false;
	for{
		i--;
		if(scanLeftFoundNumber){
			scanStartLeftStart = i;
		}
		if(i <= 0){
			break;
		}
		if(sliced[i] == '['){
			if(scanLeftFoundNumber){
				break;
			}
			continue;
		} else if(sliced[i] == ']'){
			if(scanLeftFoundNumber){
				break;
			}
			continue;
		} else if(sliced[i] == ','){
			if(scanLeftFoundNumber){
				break;
			}
			continue;
		}
		if(!scanLeftFoundNumber){
			scanStartLeftEnd = i;
			scanLeftFoundNumber = true;
		}

	}
	//Log.Info("Completed scan left. %t, %d:%d", scanLeftFoundNumber, scanStartLeftStart, scanStartLeftEnd);
	if(scanLeftFoundNumber){
		leftNumberString := sliced[scanStartLeftStart+1:scanStartLeftEnd+1];
		leftNumberParsed, err := strconv.Atoi(leftNumberString);
		if(err != nil){
			Log.Info("X1");
			Log.FatalError(err);
		}
		leftNumberParsed += left;
		newSliced := sliced[0:scanStartLeftStart+1];
		newSliced += fmt.Sprintf("%d", leftNumberParsed);
		newSliced += sliced[scanStartLeftEnd+1:];
		sliced = newSliced;
	}

	this.CurrentStatement = sliced;
	return true;
}

func (this *Problem18A) Solve() {
	Log.Info("Problem 18A solver beginning!")

	file, err := os.Open("source-data/input-day-18a-test.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	statementCount := 0;
	allStatements := make([]*SnailStatement, 0);

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			statement := &SnailStatement{};
			statement.RawStatement = line;
			statement.CurrentStatement = line;
			allStatements = append(allStatements, statement);
			statementCount++;
		}
	}

	Log.Info("Finished parsing %d snail statements", statementCount)

	sum := &SnailStatement{};

	for _, s := range allStatements {
		sum = sum.Add(s);
		sum.Reduce();
	}

	sum.Reduce();

	Log.Info("Sum string is %s", sum.CurrentStatement)

	parser := &SnailStatementParser{};
	parser.Line = sum.CurrentStatement;

	parsed := parser.Parse();

	Log.Info("Completed parse - magnitude is %d", parsed.Magnitude());
}


