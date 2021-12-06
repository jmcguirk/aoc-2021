package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem6A struct {

}



func (this *Problem6A) Solve() {
	Log.Info("Problem 6A solver beginning!")

	file, err := os.Open("source-data/input-day-06a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := make([]int, 0);

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			parts := strings.Split(line, ",");
			for _, v := range parts{
				parsed, err := strconv.Atoi(v);
				if(err != nil){
					Log.FatalError(err);
				}
				state = append(state, parsed);
			}
		}
	}

	day := 0;
	targetDay := 80;

	for{
		if(day >= targetDay){
			break;
		}
		newChildren := make([]int, 0);
		for i, v := range state{
			newV := v-1;
			if(newV < 0){
				newV = 6;
				newChildren = append(newChildren, 8);
			}
			state[i] = newV;
		}

		state = append(state, newChildren...);
		/*
		buff := "";
		for _, v := range state {
			if(buff != ""){
				buff += ",";
			}
			buff += fmt.Sprintf("%d", v);
		}*/
		day++;
		//Log.Info("After %d days: %s", day, buff);
	}

	Log.Info("After %d days, there are %d lantern fish", day, len(state));
}