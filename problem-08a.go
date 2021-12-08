package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem8A struct {

}



func (this *Problem8A) Solve() {
	Log.Info("Problem 8A solver beginning!")

	file, err := os.Open("source-data/input-day-08a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	system := &SevenSegmentSystem{};
	system.Init(4);

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			//Log.Info(line);
			system.AddEntry(line);
		}
	}
	//system.Solve();
	Log.Info("Parsed seven segment system, found %d log entries, searching with %d displays, %d unique values", len(system.Entries), len(system.Displays), system.CountUniques());
}