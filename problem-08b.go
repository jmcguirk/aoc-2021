package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem8B struct {

}



func (this *Problem8B) Solve() {
	Log.Info("Problem 8B solver beginning!")

	file, err := os.Open("source-data/input-day-08b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	system := &SevenSegmentSystem{};
	system.Init(1);

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			//Log.Info(line);
			system.AddEntry(line);
		}
	}
	solution := system.Solve();
	Log.Info("Parsed seven segment system, found %d log entries, searching with %d displays, %d unique values, solution was: %d", len(system.Entries), len(system.Displays), system.CountUniques(), solution);
}