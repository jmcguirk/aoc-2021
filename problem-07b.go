package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem7B struct {

}



func (this *Problem7B) Solve() {
	Log.Info("Problem 7B solver beginning!")

	file, err := os.Open("source-data/input-day-07b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make([]int, 0);

	maxVal := -1000000;
	minVal := 1000000;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			parts := strings.Split(line, ",");
			for _, v := range parts{
				parsed, err := strconv.Atoi(v);
				if(err != nil){
					Log.FatalError(err);
				}
				if(parsed < minVal){
					minVal = parsed;
				}
				if(parsed > maxVal){
					maxVal = parsed;
				}
				values = append(values, parsed);
			}
		}
	}

	bestFuel := 0;
	bestFuelIndex := -1;

	for i := minVal; i <= maxVal; i++{
		total := 0;
		for _, v := range values {
			delta := v - i;
			if(delta < 0){
				delta = delta * -1;
			}
			total += this.GetFuelForDist(delta);
		}
		if(bestFuelIndex < 0 || total < bestFuel){
			bestFuel = total;
			bestFuelIndex = i;
		}
	}

	Log.Info("Parsed %d crabs. %d Min Value, %d Max Value. Best position %d with total fuel %d", len(values), minVal, maxVal, bestFuelIndex, bestFuel);
}

func (this *Problem7B) GetFuelForDist(dist int) int {
	return ((dist) * (dist+1))/2;
}