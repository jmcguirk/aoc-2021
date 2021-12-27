package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Problem25A struct {

}


func (this *Problem25A) Solve() {
	Log.Info("Problem 25A solver beginning!")


	file, err := os.Open("source-data/input-day-25a-test.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := &IntegerGrid2D{};
	grid.Init()
	rows := 0;
	cols := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			//Log.Info("Parse line %s", line);
			cols = 0;
			for _, v := range line{
				grid.SetValue(cols, rows, int(v));
				cols++;
			}
			rows++;
		}

	}

	Log.Info("Parsed grid %d by %d", rows, cols);
	//fmt.Println("");
	fmt.Println(grid.PrintAscii());

	maxStep := math.MaxInt32;
	currStep := 0;

	xMin := grid.MinRow();
	xMax := grid.MaxX();

	yMin := grid.MinY();
	yMax := grid.MaxY();

	for{
		if(currStep >= maxStep){
			break;
		}

		nextGrid := &IntegerGrid2D{};
		nextGrid.Init();



		//Log.Info("Going from %d to %d x", xMin, xMax);
		//Log.Info("Going from %d to %d y", yMin, yMax);

		for y := yMin; y <= yMax; y++{
			for x := xMin; x <= xMax; x++{
				nextGrid.SetValue(x, y, int('.'));
			}
		}

		//Log.Info("Filled, now x range is %d to %d x", grid.MinX(), grid.MaxX());

		eastMovers := 0;

		for y := yMin; y <= yMax; y++{
			for x := xMin; x <= xMax; x++{
				nextX := x;
				nextY := y;
				v := grid.GetValue(x, y);
				if(v == '>'){
					nextX = x+1
					nextX = nextX % (xMax+1);
					//Log.Info("Considering move")
					if(grid.GetValue(nextX, nextY) == '.'){
						nextGrid.SetValue(nextX, nextY, '>');
						//Log.Info("Moving from %d,%d to %d, %d", x, y, nextX, nextY);
						eastMovers++;
					} else{
						//Log.Info("Cannot move %c to %d,%d from %d,%d", grid.GetValue(nextX, nextY), nextX, nextY, x, y);
						nextGrid.SetValue(x, y, '>');
					}
				}
			}
		}

		southMovers := 0;
		halfGrid := nextGrid.Clone();

		nextGrid = &IntegerGrid2D{};
		nextGrid.Init();

		for y := yMin; y <= yMax; y++{
			for x := xMin; x <= xMax; x++{
				nextGrid.SetValue(x, y, halfGrid.GetValue(x, y));
			}
		}


		for y := yMin; y <= yMax; y++{
			for x := xMin; x <= xMax; x++{
				nextX := x;
				nextY := y;
				v := grid.GetValue(x, y);
				if(v == 'v'){
					nextY = y + 1;
					nextY = nextY % (yMax+1);
					//Log.Info("Moving down from %d,%d to %d,%d", x, y, nextX, nextY)
					if(halfGrid.GetValue(nextX, nextY) == '.' && grid.GetValue(nextX, nextY) != 'v'){
						nextGrid.SetValue(nextX, nextY, 'v');
						southMovers++;
					} else{
						nextGrid.SetValue(x, y, 'v');
					}
				}

			}
		}

		grid = nextGrid;



		currStep++;
		//Log.Info("After %d steps", currStep)
		//fmt.Println(grid.PrintAscii());

		if(eastMovers == 0 && southMovers == 0){
			break;
		}

		//Log.Info("After %d steps", currStep);
		//
	}
	Log.Info("Simulation of %d steps complete", currStep)
}
