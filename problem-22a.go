package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem22A struct {
	Grid *IntegerGrid3D;
	Instructions []*BootInstruction;
	MinRange int;
	MaxRange int;
}

type BootInstruction struct {
	On bool;
	XRange IntVec2;
	YRange IntVec2;
	ZRange IntVec2;
	InstructionNumber int;

}

func (this *BootInstruction) Describe() string {
	op := "off";
	if(this.On){
		op = "on";
	}
	return fmt.Sprintf("%s x=%d..%d,y=%d..%d,z=%d..%d", op, this.XRange.X, this.XRange.Y, this.YRange.X, this.YRange.Y, this.ZRange.X, this.ZRange.Y);
}

func (this *BootInstruction) Parse(line string) bool{
	spaceParts := strings.Split(line, " ");
	op := strings.TrimSpace(spaceParts[0]);
	if(op == "off"){
		this.On = false;
	} else if(op == "on"){
		this.On = true;
	} else{
		Log.Fatal("Unknown operation")
	}

	ranges := strings.Split(strings.TrimSpace(spaceParts[1]), ",");

	for _, v := range ranges{
		rangeAxisParts := strings.Split(strings.TrimSpace(v), "=");
		axis := strings.TrimSpace(rangeAxisParts[0]);
		axisParts := strings.Split(strings.TrimSpace(rangeAxisParts[1]), "..");
		if(len(axisParts) != 2){
			Log.Fatal("Incorrect number of axis parts");
		}
		axisMin, err := strconv.Atoi(strings.TrimSpace(axisParts[0]));
		if(err != nil){
			Log.FatalError(err);
		}
		axisMax, err := strconv.Atoi(strings.TrimSpace(axisParts[1]));
		if(err != nil){
			Log.FatalError(err);
		}
		r := IntVec2{};
		r.X = axisMin;
		r.Y = axisMax;
		if(axis == "x"){
			this.XRange = r;
		} else if(axis == "y"){
			this.YRange = r;
		} else if(axis == "z") {
			this.ZRange = r;
		} else {
			Log.Fatal("Unknown axis %s", axis);
		}
	}
	return true;
}

func (this *Problem22A) Solve() {
	Log.Info("Problem 22A solver beginning!")


	file, err := os.Open("source-data/input-day-22a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	this.Grid = &IntegerGrid3D{};
	this.Grid.Init();

	this.Instructions = make([]*BootInstruction, 0);


	this.MinRange = -50;
	this.MaxRange = 50;


	scanner := bufio.NewScanner(file)

	instructionNumber := 1;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			instruction := &BootInstruction{};
			if(!instruction.Parse(line)){
				Log.Fatal("Failed ot parse line")
			}
			instruction.InstructionNumber = instructionNumber;
			this.Instructions = append(this.Instructions, instruction);
			instructionNumber++;
		}
	}


	Log.Info("Finished parsing boot instructions - %d total instructions", len(this.Instructions));
	for _, instruction := range this.Instructions{
		this.Apply(instruction);
	}
	Log.Info("Finished processing %d boot instructions total lights on: %d", len(this.Instructions), this.Grid.CountGreaterThan(0));
}

func (this *Problem22A) Apply(instruction *BootInstruction) {
	for x := instruction.XRange.X; x <= instruction.XRange.Y; x++{
		if(x < this.MinRange || x > this.MaxRange){
			continue;
		}
		for y := instruction.YRange.X; y <= instruction.YRange.Y; y++{
			if(y < this.MinRange || y > this.MaxRange){
				continue;
			}
			for z := instruction.ZRange.X; z <= instruction.ZRange.Y; z++{
				if(z < this.MinRange || z > this.MaxRange){
					continue;
				}

				val := 0;
				if(instruction.On){
					val = 1;
				}
				this.Grid.SetValue(x, y, z, val);
			}
		}
	}
}
