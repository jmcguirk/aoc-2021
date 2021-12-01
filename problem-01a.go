package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem1A struct {

}

func (this *Problem1A) Solve() {
	Log.Info("Problem 1A solver beginning!")


	file, err := os.Open("source-data/input-day-01a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	hasParsedFirstVal := false;
	prevVal := int64(0);
	incrementingCount := 0;
	totalVals := 0;

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			val, err := strconv.ParseInt(line, 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			if(!hasParsedFirstVal){
				hasParsedFirstVal = true;
			} else if(val > prevVal){
				incrementingCount++;
			}
			prevVal = val;
			totalVals++;
		}
	}

	Log.Info("Finished parsing file - %d of %d lines were increasing", incrementingCount, totalVals);
}
