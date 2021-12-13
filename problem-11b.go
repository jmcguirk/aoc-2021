package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem11B struct {

}



func (this *Problem11B) Solve() {
	Log.Info("Problem 10B solver beginning!")

	file, err := os.Open("source-data/input-day-11b.txt");
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

	totalFlashes := 0;
	stepCount := 0;
	didSync := false;
	for{

		if(didSync){
			break;
		}

		rowMin := grid.MinY();
		rowMax := grid.MaxY();

		colMin := grid.MinX();
		colMax := grid.MaxX();


		flashed := &IntegerGrid2D{};
		flashed.Init();

		flashQueue := make([]*IntVec2, 0);
		gridSize := 0;
		// Step 1, increment everything
		for i := rowMin; i <= rowMax; i++{
			for j := colMin; j <= colMax; j++{
				if(!grid.HasValue(j, i)){
					continue;
				}
				val := grid.GetValue(j, i);
				gridSize++;
				val++;
				grid.SetValue(j, i, val);
			}

		}

		flashesThisRound := 0;
		// Step 2, process flashes
		for{
			for i := rowMin; i <= rowMax; i++{
				for j := colMin; j <= colMax; j++{
					if(!grid.HasValue(j, i)){
						continue;
					}
					val := grid.GetValue(j, i);

					if(val == 10){
						if(flashed.HasValue(j, i)){
							continue;
						}
						v := &IntVec2{};
						v.X = j;
						v.Y = i;
						flashQueue = append(flashQueue, v);
					}
				}
			}
			if(len(flashQueue) == 0){
				break;
			}

			for _, flash := range flashQueue{
				j := flash.X;
				i := flash.Y;
				if(flashed.HasValue(j, i)){
					continue;
				}
				totalFlashes++;
				flashesThisRound++;
				flashed.SetValue(j, i, 1);
				for x := -1; x <= 1; x++{
					for y := -1; y <= 1; y++{
						if(x == 0 && y == 0){
							continue;
						}
						dJ := j + x;
						dI := i + y;
						if(!grid.HasValue(dJ, dI)){
							continue;
						}
						val := grid.GetValue(dJ, dI);
						if(val == 10){
							continue;
						}
						val++;
						grid.SetValue(dJ, dI, val);
					}
				}
			}
			flashQueue = make([]*IntVec2, 0);
		}
		if(flashesThisRound == gridSize){
			didSync = true;
			break;
		}

		for i := rowMin; i <= rowMax; i++{
			for j := colMin; j <= colMax; j++{
				if(!grid.HasValue(j, i)){
					continue;
				}
				val := grid.GetValue(j, i);

				if(val == 10){
					grid.SetValue(j, i, 0);
				}
			}
		}

		stepCount++;
		//Log.Info("\nAfter %d steps, %d flashes", stepCount, totalFlashes);
		//fmt.Print(grid.PrintWithZero("0"));
	}


	Log.Info("After %d steps we found a sync point", stepCount+1)
}
