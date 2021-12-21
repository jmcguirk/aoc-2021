package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem20B struct {

}


func (this *Problem20B) Solve() {
	Log.Info("Problem 20B solver beginning!")


	file, err := os.Open("source-data/input-day-20b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hasParsedInstructionLine := false;
	instructions := make([]int, 0);
	grid := &IntegerGrid2D{};
	grid.Init()
	rows := 0;
	cols := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			if(!hasParsedInstructionLine){
				for _, r := range line {
					if(r == '.') {
						instructions = append(instructions, 0);
					}
					if(r == '#') {
						instructions = append(instructions, 1);
					}
				}
			} else{
				Log.Info("Parse line %s", line);
				cols = 0;
				for _, v := range line{
					val := 0;
					if(v == '#'){
						val = 1;
					}
					grid.SetValue(cols, rows, val);
					cols++;
				}
				rows++;
			}
		} else{
			hasParsedInstructionLine = true;
		}
	}

	Log.Info("Parsed instruction set %d length. Parsed grid %d by %d", len(instructions), rows, cols);
	//fmt.Println("");
	//fmt.Println(grid.Print())

	maxStep := 50;
	currStep := 0;


	for{
		if(currStep >= maxStep){
			break;
		}

		nextGrid := &IntegerGrid2D{};
		nextGrid.Init();

		xMin := grid.MinX() - 1;
		xMax := grid.MaxX() + 1;

		yMin := grid.MinY() - 1;
		yMax := grid.MaxY() + 1;

		for y := yMin; y <= yMax; y++{
			for x := xMin; x <= xMax; x++{

				index := 0;
				for m := -1; m <= 1; m++{
					for n := -1; n <= 1; n++{
						index = index << 1;
						vY := y + m;
						vX := x + n;

						if(!grid.HasValue(vX, vY)){
							if(currStep % 2 == 1 && instructions[0] == 1){ // This square has not been simulated. If we're in an odd step, and the first index in the instruction requires us to be on, turn it on
								index = index | 1;
							}
						} else{
							val := grid.GetValue(vX, vY);
							if(val > 0){
								index = index | 1;
							}
						}
					}
				}
				nextGrid.SetValue(x, y, instructions[index]);
			}
		}
		grid = nextGrid;

		currStep++;



		//Log.Info("After %d steps", currStep);
		//
	}


	lightsOn := grid.CountGreaterThan(0);
	Log.Info("Simulation of %d steps complete, %d lights on", maxStep, lightsOn)

	//fmt.Println(grid.PrintAscii());
}
