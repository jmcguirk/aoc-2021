package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem5B struct {

}



func (this *Problem5B) Solve() {
	Log.Info("Problem 5B solver beginning!")

	file, err := os.Open("source-data/input-day-05b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]*IntegerLineSegment2D, 0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			parsed := &IntegerLineSegment2D{};
			if(parsed.Parse(line)){
				lines = append(lines, parsed);
			}
		}
	}

	grid := &IntegerGrid2D{};
	grid.Init();



	for _, line := range lines {
		grid.Paint(line);
	}

	Log.Info("Parsed %d lines, Total with at least one intersection %d", len(lines),  grid.CountGreaterThan(1));
}