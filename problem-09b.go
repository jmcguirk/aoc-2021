package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Problem9B struct {
	CurrentBasinSize int;
	Grid *IntegerGrid2D;
	Visited *IntegerGrid2D;
}



func (this *Problem9B) Solve() {
	Log.Info("Problem 9B solver beginning!")

	file, err := os.Open("source-data/input-day-09b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	this.Grid = &IntegerGrid2D{};
	this.Grid.Init()

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
				this.Grid.SetValue(col, row, parsed);
				col++;
			}
			row++;
		}
	}


	lowPoints := 0;
	lowPointsSum := 0;

	rowMin := this.Grid.MinY();
	rowMax := this.Grid.MaxY();

	colMin := this.Grid.MinX();
	colMax := this.Grid.MaxX();

	//Log.Info("Starting scan %d:%d %d:%d", rowMin, rowMax, colMin, colMax);

	basinSizes := make([]int, 0);

	for i := rowMin; i <= rowMax; i++{
		for j := colMin; j <= colMax; j++{
			if(!this.Grid.HasValue(j, i)){
				continue;
			}
			val := this.Grid.GetValue(j, i);
			//Log.Info("Considering %d, %d : %d", i, j, val);

			if(this.Grid.HasValue(j - 1, i) && this.Grid.GetValue(j - 1, i) <= val){
				continue;
			}
			if(this.Grid.HasValue(j + 1, i) && this.Grid.GetValue(j + 1, i) <= val){
				continue;
			}
			if(this.Grid.HasValue(j, i - 1) && this.Grid.GetValue(j, i - 1) <= val){
				continue;
			}
			if(this.Grid.HasValue(j, i + 1) && this.Grid.GetValue(j, i + 1) <= val){
				continue;
			}
			//Log.Info("Found low point %d, at pos %d, %d", val, i, j);
			lowPoints++;
			lowPointsSum += val+1;
			this.Visited = &IntegerGrid2D{};
			this.Visited.Init();
			this.CurrentBasinSize = 0;
			this.ExploreBasin(j, i)
			Log.Info("Finished exploring basin at %d, %d for value %d. Basin size is %d", i, j, val, this.CurrentBasinSize);
			basinSizes = append(basinSizes, this.CurrentBasinSize);
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)));

	product := 1;
	for i := 0; i < 3; i++ {
		product *= basinSizes[i];
	}

	Log.Info("Finished processing heatmap, %d low points, total sum %d. Explored %d basins. Checksum was %d", lowPoints, lowPointsSum, len(basinSizes), product);
}

func (this *Problem9B) ExploreBasin(j int, i int) {
	if(!this.Grid.HasValue(j, i)){
		return;
	}
	if(this.Visited.IsVisited(j, i)){
		return;
	}
	this.Visited.Visit(j, i);
	val := this.Grid.GetValue(j, i);
	if(val == 9){
		return;
	}
	this.CurrentBasinSize++;
	this.ExploreBasin(j - 1, i);
	this.ExploreBasin(j + 1, i);
	this.ExploreBasin(j, i - 1);
	this.ExploreBasin(j, i + 1);
}