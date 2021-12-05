package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem5A struct {

}



func (this *Problem5A) Solve() {
	Log.Info("Problem 5A solver beginning!")

	file, err := os.Open("source-data/input-day-05a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]*IntegerLineSegment2D, 0);
	straightLines := make([]*IntegerLineSegment2D, 0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			parsed := &IntegerLineSegment2D{};
			if(parsed.Parse(line)){
				lines = append(lines, parsed)
				if(parsed.IsStraight()){
					straightLines = append(straightLines, parsed)
				}
				//parsed.Log();
			}
		}
	}

	grid := &IntegerGrid2D{};
	grid.Init();



	for _, line := range straightLines {
		grid.Paint(line);
	}

	Log.Info("Parsed %d lines, %d were straight. Total with at least one intersection %d", len(lines), len(straightLines), grid.CountGreaterThan(1));
}