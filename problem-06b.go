package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem6B struct {

}



func (this *Problem6B) Solve() {
	Log.Info("Problem 6B solver beginning!")

	file, err := os.Open("source-data/input-day-06b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := make(map[int]int);

	for i := 0; i <= 8; i++{
		state[i] = 0;
	}
	populationSize := 0;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			parts := strings.Split(line, ",");
			for _, v := range parts{
				parsed, err := strconv.Atoi(v);
				if(err != nil){
					Log.FatalError(err);
				}
				state[parsed]++;
				populationSize++;
			}
		}
	}

	/*
	Log.Info("Initial state")
	for i := 0; i <= 8; i++{
		Log.Info("%d:%d", i, state[i])
	}*/

	day := 0;
	targetDay := 256;


	for{
		if(day >= targetDay){
			break;
		}

		nextState := make(map[int]int);
		newBorns := 0;
		for k, v := range state{
			if(k == 0){
				newBorns+=v;
			} else{
				nextState[k-1] = v;
			}
		}
		nextState[8] = newBorns;
		nextState[6] += newBorns;
		populationSize += newBorns;

		day++;
		state = nextState;
		/*
		Log.Info("After %d days: %d population size. New state is", day, populationSize);
		for i := 0; i <= 8; i++{
			Log.Info("%d:%d", i, state[i])
		}*/

	}

	Log.Info("After %d days, there are %d lantern fish", day, populationSize);
}