package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem15A struct {

}



func (this *Problem15A) Solve() {
	Log.Info("Problem 15A solver beginning!")

	file, err := os.Open("source-data/input-day-15a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)

	grid := IntegerGrid2D{};
	grid.Init()

	row := 0;
	lastX := 0;
	lastY := 0;

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
				lastX = col;
				lastY = row;
				col++;
			}
			row++;
		}
	}

	fmt.Println(grid.Print());
	start := &IntVec2{};
	finish := &IntVec2{};
	finish.X = lastX;
	finish.Y = lastY;

	Log.Info("Parsed initial graph, charting a path from %s to %s", start.ToString(), finish.ToString());
	path := grid.ShortestPath(start, finish, -1);
	Log.Info("Found path of length %d", len(path));
	totalRisk := 0;
	for _, v := range path {
		totalRisk += grid.GetValue(v.X, v.Y);
		//Log.Info("%d,%d : %d", v.X, v.Y, grid.GetValue(v.X, v.Y));
	}
	Log.Info("Best path total risk is %d", totalRisk)
}
