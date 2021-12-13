package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem13A struct {
	Grid *IntegerGrid2D;
	Instructions []*FoldInstruction;
}

type FoldInstruction struct {
	Pivot int;
	IsVerticalFold bool;
}

func (this *FoldInstruction) Describe() string {
	if(this.IsVerticalFold){
		return fmt.Sprintf("fold along y=%d", this.Pivot)
	} else{
		return fmt.Sprintf("fold along x=%d", this.Pivot)
	}
}

func (this *Problem13A) NumDots() int {
	xMin := this.Grid.MinRow();
	xMax := this.Grid.MaxRow();

	yMin := this.Grid.MinCol();
	yMax := this.Grid.MaxCol();

	total := 0;
	for j := yMin; j<= yMax; j++{
		for i := xMin; i<= xMax; i++{
			if(!this.Grid.HasValue(i, j)){
				continue;
			} else{
				val := this.Grid.GetValue(i, j);
				if(val > 0){
					total++;
				}
			}
		}
	}
	return total;
}


func (this *Problem13A) Solve() {
	Log.Info("Problem 13A solver beginning!")

	file, err := os.Open("source-data/input-day-13a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)


	this.Grid = &IntegerGrid2D{};
	this.Grid.Init();

	this.Instructions = make([]*FoldInstruction, 0);

	hasParsedGrid := false;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			if(!hasParsedGrid){
				lineParts := strings.Split(line, ",");
				x, err := strconv.Atoi(strings.TrimSpace(lineParts[0]));
				Log.FatalIfError(err);
				y, err := strconv.Atoi(strings.TrimSpace(lineParts[1]));
				Log.FatalIfError(err);
				this.Grid.SetValue(x, y, int('#'));
				Log.Info("Parsed value %d,%d", x,y);
			} else {
				lineParts := strings.Split(line, " ");
				if (len(lineParts) != 3) {
					Log.Fatal("Wrong number of line parts");
				}

				instructionRaw := strings.Split(strings.TrimSpace(lineParts[2]), "=");
				if (len(instructionRaw) != 2) {
					Log.Fatal("Wrong number of instruction parts");
				}

				pivot, err := strconv.Atoi(strings.TrimSpace(instructionRaw[1]));
				Log.FatalIfError(err);

				instruction := &FoldInstruction{};
				instruction.Pivot = pivot;
				if(instructionRaw[0] == "y"){
					instruction.IsVerticalFold = true;
				}
				this.Instructions = append(this.Instructions, instruction);
			}
		} else{
			hasParsedGrid = true;
		}
	}

	//fmt.Println(this.Grid.PrintAscii())

	Log.Info("Parsed grid and %d instructions. Grid starts with %d dots", len(this.Instructions), this.NumDots());
	firstInstructionOnly := true;

	for i, v := range this.Instructions{
		Log.Info("Applying fold %s", v.Describe());
		if(v.IsVerticalFold){
			this.ApplyVerticalFold(v);
		} else{
			this.ApplyHorizontalFold(v);
		}
		//fmt.Println(this.Grid.PrintAscii())
		Log.Info("Fold complete, %d dots", this.NumDots());
		if(i == 0 && firstInstructionOnly){
			break;
		}
	}
}


func (this *Problem13A) ApplyVerticalFold(fold *FoldInstruction) {
	newGrid := &IntegerGrid2D{};
	newGrid.Init();

	xMin := this.Grid.MinRow();
	xMax := this.Grid.MaxRow();

	yMin := this.Grid.MinCol();
	yMax := this.Grid.MaxCol();

	//total := 0;
	for j := yMin; j<= yMax; j++{
		for i := xMin; i<= xMax; i++{
			if(!this.Grid.HasValue(i, j)){
				continue;
			} else{
				val := this.Grid.GetValue(i, j);
				if(val > 0){
					if(j < fold.Pivot){
						newGrid.SetValue(i, j, '#');
					} else{
						delta := j - fold.Pivot;
						newPos := fold.Pivot - delta;
						newGrid.SetValue(i, newPos, '#');
					}
				}
 			}
		}
	}
	this.Grid = newGrid;

}

func (this *Problem13A) ApplyHorizontalFold(fold *FoldInstruction) {
	newGrid := &IntegerGrid2D{};
	newGrid.Init();

	xMin := this.Grid.MinRow();
	xMax := this.Grid.MaxRow();

	yMin := this.Grid.MinCol();
	yMax := this.Grid.MaxCol();

	//total := 0;
	for j := yMin; j<= yMax; j++{
		for i := xMin; i<= xMax; i++{
			if(!this.Grid.HasValue(i, j)){
				continue;
			} else{
				val := this.Grid.GetValue(i, j);
				if(val > 0){
					if(i < fold.Pivot){
						newGrid.SetValue(i, j, '#');
					} else{
						delta := i - fold.Pivot;
						newPos := fold.Pivot - delta;
						newGrid.SetValue(newPos, j, '#');
					}
				}
			}
		}
	}
	this.Grid = newGrid;

}


