package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem9A struct {

}



func (this *Problem9A) Solve() {
	Log.Info("Problem 9A solver beginning!")

	file, err := os.Open("source-data/input-day-09a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := IntegerGrid2D{};
	grid.Init()

	row := 0;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		col := 0;
		if (line != "") {
			for _, v := range line{
				parsed, err := strconv.Atoi(fmt.Sprintf("%c", v));
				if(err != nil){
					Log.FatalError(err);
				}
				grid.SetValue(col, row, parsed);
				col++;
			}
			row++;
		}
	}

	fmt.Print(grid.Print());

	lowPoints := 0;
	lowPointsSum := 0;

	rowMin := grid.MinY();
	rowMax := grid.MaxY();

	colMin := grid.MinX();
	colMax := grid.MaxX();

	Log.Info("Starting scan %d:%d %d:%d", rowMin, rowMax, colMin, colMax);

	for i := rowMin; i <= rowMax; i++{
		for j := colMin; j <= colMax; j++{
			if(!grid.HasValue(j, i)){
				continue;
			}
			val := grid.GetValue(j, i);
			//Log.Info("Considering %d, %d : %d", i, j, val);

			if(grid.HasValue(j - 1, i) && grid.GetValue(j - 1, i) <= val){
				continue;
			}
			if(grid.HasValue(j + 1, i) && grid.GetValue(j + 1, i) <= val){
				continue;
			}
			if(grid.HasValue(j, i - 1) && grid.GetValue(j, i - 1) <= val){
				continue;
			}
			if(grid.HasValue(j, i + 1) && grid.GetValue(j, i + 1) <= val){
				continue;
			}
			//Log.Info("Found low point %d, at pos %d, %d", val, i, j);
			lowPoints++;
			lowPointsSum += val+1;
		}
	}

	Log.Info("Finished processing heatmap, %d low points, total sum %d", lowPoints, lowPointsSum)
}
