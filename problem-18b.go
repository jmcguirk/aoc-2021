package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem18B struct {

}


func (this *Problem18B) Solve() {
	Log.Info("Problem 18B solver beginning!")

	file, err := os.Open("source-data/input-day-18b.txt");
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

	peakMagnitude := 0;

	for i, s1 := range allStatements{
		for j, s2 := range allStatements{
			if(i == j){
				continue;
			}
			sum := s1.Add(s2);
			sum.Reduce();

			parser := &SnailStatementParser{};
			parser.Line = sum.CurrentStatement;

			parsed := parser.Parse();
			m := parsed.Magnitude();
			if(m> peakMagnitude){
				peakMagnitude = m;
			}
		}
	}


	Log.Info("Completed parsing - peak magnitude is is %d", peakMagnitude);
}


