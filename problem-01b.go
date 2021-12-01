package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem1B struct {

}

func (this *Problem1B) Solve() {
	Log.Info("Problem 1B solver beginning!")


	file, err := os.Open("source-data/input-day-01b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	hasParsedFirstVal := false;
	prevVal := int64(0);
	incrementingCount := 0;
	totalVals := 0;
	window := make([]int64, 0);
	windowSize := 3;

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			val, err := strconv.ParseInt(line, 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			window = append(window, val);
			if(len(window) >= windowSize){
				if (len(window) > windowSize){
					window = window[1:];
				}
				sum := int64(0);
				for i := 0; i < windowSize; i++ {
					sum += window[i];
				}
				if(!hasParsedFirstVal){
					hasParsedFirstVal = true;
				} else if(sum > prevVal){
					incrementingCount++;
				}
				prevVal = sum;
			}
			totalVals++;
		}
	}

	Log.Info("Finished parsing file using window size %d - %d of %d lines were increasing", windowSize, incrementingCount, totalVals);
}
