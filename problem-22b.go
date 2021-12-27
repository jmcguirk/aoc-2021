package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Problem22B struct {
	Grid *IntegerGrid3D;
	Instructions []*BootInstruction2;
	MinRange int;
	MaxRange int;
}


type BootInstruction2 struct {
	On bool;
	Region Region;
	InstructionNumber int;

}

func (this *BootInstruction2) Parse(line string) bool{
	this.Region = Region{};
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
		r := Interval{};
		r.Min = axisMin;
		r.Max = axisMax;
		if(axis == "x"){
			this.Region.X = r;
		} else if(axis == "y"){
			this.Region.Y = r;
		} else if(axis == "z") {
			this.Region.Z  = r;
		} else {
			Log.Fatal("Unknown axis %s", axis);
		}
	}



	return true;
}

func (this *Problem22B) Solve() {
	Log.Info("Problem 22B solver beginning!")


	file, err := os.Open("source-data/input-day-22b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	this.Grid = &IntegerGrid3D{};
	this.Grid.Init();

	this.Instructions = make([]*BootInstruction2, 0);


	this.MinRange = math.MinInt32;
	this.MaxRange = math.MaxInt32;


	scanner := bufio.NewScanner(file)

	instructionNumber := 1;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			instruction := &BootInstruction2{};
			if(!instruction.Parse(line)){
				Log.Fatal("Failed ot parse line")
			}
			instruction.InstructionNumber = instructionNumber;
			this.Instructions = append(this.Instructions, instruction);
			instructionNumber++;
		}
	}


	Log.Info("Finished parsing boot instructions - %d total instructions", len(this.Instructions));

	regionsOn := make([]Region, 0);

	for _, v := range this.Instructions{
		if(v.On){
			intersected := make([]Region, 0);
			intersected = append(intersected, v.Region);
			for _, r := range regionsOn{
				newIntersected := make([]Region, 0)
				for _, r2 := range intersected{
					subCubes := r2.Subtract(r);
					for _, subCube := range subCubes{
						if(subCube.Size() > 0){
							newIntersected = append(newIntersected, subCube);
						}
					}
				}
				intersected = newIntersected;
			}
			for _, r := range intersected{
				regionsOn = append(regionsOn, r);
			}
		} else{
			intersected := make([]Region, 0);
			for _, r := range regionsOn{
				subCubes := r.Subtract(v.Region);
				for _, subCube := range subCubes {
					if(subCube.Size() > 0){
						intersected = append(intersected, subCube);
					}
				}
			}
			regionsOn = intersected;
		}
	}

	Log.Info("Completed simulation, total lights on %d", this.Size(regionsOn));
}

func (this *Problem22B) Size(regions []Region) int{
	size := 0;
	for _, v := range regions {
		size += v.Size();
	}
	return size;
}
