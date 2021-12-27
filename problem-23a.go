package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Problem23A struct {
	BestCosts map[string]int;
	BestEnergy int;
}

type AmpSystem struct {
	Grid *IntegerGrid2D;
	EnergySpentSoFar int;
	GoalColumns []int;
	HallwayRow int;
	Units []*AmpUnit;
	UpperGoal int;
	LowerGoal int;
	Depth int;
	ParentState *AmpSystem;
	ParentStateSerialized string;
	MoveDescription string;
}


type AmpUnit struct {
	Letter rune;
	PosX int;
	PosY int;
	HasMovedToGoal bool;
	IsInHallway bool;
	HasMovedToHallway bool;
	UnitId int;
}

func (this *AmpSystem) Clone(moveDescription string) *AmpSystem {
	res := &AmpSystem{};
	res.ParentState = this;
	res.Depth = this.Depth + 1;
	res.Grid = this.Grid.Clone();
	res.ParentStateSerialized = this.Grid.PrintAscii();
	res.EnergySpentSoFar = this.EnergySpentSoFar;
	res.MoveDescription = moveDescription;

	res.GoalColumns = make([]int, len(this.GoalColumns));
	for i, v := range this.GoalColumns{
		res.GoalColumns[i] = v;
	}

	res.HallwayRow = this.HallwayRow;

	res.Units = make([]*AmpUnit, len(this.Units));
	for j, u := range this.Units{
		res.Units[j] = u.Clone();
	}
	res.UpperGoal = this.UpperGoal;
	res.LowerGoal = this.LowerGoal;

	return res;
}

func (this *AmpSystem) Path() []*AmpSystem {
	next := this;
	res := make([]*AmpSystem, 0);
	for{
		if(next == nil){
			break;
		}
		res = append(res, next);
		next = next.ParentState;
	}
	ReverseSlice(res);
	return res;
}

func (this *AmpUnit) Clone() *AmpUnit {
	res := &AmpUnit{};
	res.Letter = this.Letter;
	res.PosX = this.PosX;
	res.PosY = this.PosY;
	res.HasMovedToGoal = this.HasMovedToGoal;
	res.HasMovedToHallway = this.HasMovedToHallway;
	res.UnitId = this.UnitId;
	return res;
}

func (this *AmpSystem) MoveUnit(unitId int, posX int, posY int) {
	unit := this.GetUnitById(unitId);
	this.Grid.SetValue(unit.PosX, unit.PosY, '.');
	this.Grid.SetValue(posX, posY, int(unit.Letter));

	dX := posX - unit.PosX;
	if(dX < 0){
		dX *= -1;
	}
	dY := 0;
	if(unit.PosY > this.HallwayRow && posY != this.HallwayRow){ // We are currently in a goal and are moving directly into another goal
		dY = this.HallwayRow - unit.PosY; // Up
		dY += this.HallwayRow - posY; // Down;
		if(dY < 0){
			dY *= -1;
		}
	} else{
		dY = posY - unit.PosY;
		if(dY < 0){
			dY *= -1;
		}
	}

	steps := dX + dY;
	unit.PosX = posX;
	unit.PosY = posY;
	if(unit.Letter == 'A'){
		this.EnergySpentSoFar += steps * 1;
	} else if(unit.Letter == 'B'){
		this.EnergySpentSoFar += steps * 10;
	} else if(unit.Letter == 'C'){
		this.EnergySpentSoFar += steps * 100;
	} else if(unit.Letter == 'D'){
		this.EnergySpentSoFar += steps * 1000;
	}
}

func (this *AmpSystem) GetUnitById(unitId int) *AmpUnit {
	for _, v := range this.Units{
		if(v.UnitId == unitId){
			return v;
		}
	}
	return nil;
}

func (this *AmpSystem) TargetIndex(letter rune) int {
	if(letter == 'A'){
		return 3;
	}
	if(letter == 'B'){
		return 5;
	}
	if(letter == 'C'){
		return 7;
	}
	if(letter == 'D'){
		return 9;
	}
	return -1;
}


func (this *AmpSystem) SuccessorStates() []*AmpSystem {
	res := make([]*AmpSystem, 0);

	minX := this.Grid.MinX();
	maxX := this.Grid.MaxX();

	for _, unit := range this.Units{
		if(unit.HasMovedToGoal){
			continue;
		}

		if(!unit.HasMovedToHallway){
			if(this.Grid.GetValue(unit.PosX, unit.PosY-1) != int('.')){
				// We haven't moved into the hallway, and there is someone above us
				continue;
			}

			for i := minX; i <= maxX; i++{
				if(this.Grid.GetValue(i, this.HallwayRow) != int('.')){ // This is occupied, or is a wall
					continue;
				}
				isExit := false;
				for _, j := range this.GoalColumns{
					if(j == i){
						isExit = true;
						break;
					}
				}
				if(isExit){
					continue;
				}

				startScan := i;
				endScan := unit.PosX;
				if(unit.PosX < i){
					startScan = unit.PosX;
					endScan = i;
				}

				inWay := false;
				for j := startScan; j <= endScan; j++{
					if(this.Grid.GetValue(j, this.HallwayRow) != int('.')){ // This is occupied, or is a wall
						if(j == unit.PosX && this.HallwayRow == unit.PosY){
							continue;
						}
						inWay = true;
					}
				}
				if(inWay){
					continue;
				}

				next := this.Clone(fmt.Sprintf("Moving unit %d (%c) into hallway at position %d, %d", unit.UnitId, unit.Letter, i, this.HallwayRow));
				next.MoveUnit(unit.UnitId, i, this.HallwayRow);
				next.GetUnitById(unit.UnitId).HasMovedToHallway = true;
				res = append(res, next);
			}
		}
		if(!unit.HasMovedToGoal){
			if(unit.PosY == this.LowerGoal){
				if(this.Grid.GetValue(unit.PosY, this.UpperGoal) != int('.')){ // This is occupied, or is a wall
					continue;
				}
			}
			for _, i := range this.GoalColumns {
				if(i != this.TargetIndex(unit.Letter)){
					continue;
				}
				lowerIsAvailable := false;
				upperIsAvailable := false;

				startScan := i;
				endScan := unit.PosX;
				if(unit.PosX < i){
					startScan = unit.PosX;
					endScan = i;
				}

				inWay := false;
				for j := startScan; j <= endScan; j++{
					if(this.Grid.GetValue(j, this.HallwayRow) != int('.')){ // This is occupied, or is a wall
						if(j == unit.PosX && this.HallwayRow == unit.PosY){
							continue;
						}
						inWay = true;
					}
				}
				if(inWay){
					continue;
				}



				if(this.Grid.GetValue(i, this.UpperGoal) == int('.') && this.Grid.GetValue(i, this.LowerGoal) == int('.')){
					lowerIsAvailable = true;
				}

				if(this.Grid.GetValue(i, this.UpperGoal) == int('.') && this.Grid.GetValue(i, this.LowerGoal) == int(unit.Letter)){
					upperIsAvailable = true;
				}

				if(lowerIsAvailable){
					next := this.Clone(fmt.Sprintf("Moving unit %d (%c) into lower goal position %d, %d", unit.UnitId, unit.Letter, i, this.LowerGoal));
					next.MoveUnit(unit.UnitId, i, this.LowerGoal);
					next.GetUnitById(unit.UnitId).HasMovedToGoal = true;
					res = append(res, next);
				} else if(upperIsAvailable){
					next := this.Clone(fmt.Sprintf("Moving unit %d (%c) into upper goal position %d, %d", unit.UnitId, unit.Letter, i, this.UpperGoal));
					next.MoveUnit(unit.UnitId, i, this.UpperGoal);
					next.GetUnitById(unit.UnitId).HasMovedToGoal = true;
					res = append(res, next);
				}
			}
		}
	}
	return res;
}

func (this *AmpSystem) IsCompleteState() bool {
	byUnit := make(map[rune][]*AmpUnit);
	//Log.Info("Check %d", len(this.Units));
	for _, unit := range this.Units{
		existing, exists := byUnit[unit.Letter];
		if(!exists){
			existing = make([]*AmpUnit, 0);
			byUnit[unit.Letter] = existing;
		}
		existing = append(existing, unit);
		byUnit[unit.Letter] = existing;
	}

	misMatch := false;
	for _, r := range byUnit{
		thisMisMatch := false;
		for _, v := range r{
			for _, v2 := range r{
				if(v2.PosX != this.TargetIndex(v2.Letter)){
					thisMisMatch = true;
				}
			}
			if(v.PosY <= this.HallwayRow){
				thisMisMatch = true;
			}
		}
		if(thisMisMatch){
			misMatch = true;
			break;
		}
	}
	if(misMatch){
		return false;
	}
	return true;
}

func (this *Problem23A) Solve() {
	Log.Info("Problem 23A solver beginning!")


	file, err := os.Open("source-data/input-day-23a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	this.BestCosts = make(map[string]int);

	system := &AmpSystem{};
	system.Grid = &IntegerGrid2D{};
	system.Grid.Init();
	system.MoveDescription = "Initial";
	system.GoalColumns = make([]int, 0);
	system.Units = make([]*AmpUnit, 0);

	frontier := make([]*AmpSystem, 0);

	rows := 0;

	unitId := 1;

	for scanner.Scan() {
		line := scanner.Text();
		if(line != ""){
			cols := 0;
			for _, v := range line{
				if(v == '#' || v == ' ') {
					// Wall, just add it in
				} else if(v == '.'){
					//Log.Fatal("Found hallway row at %d", rows);
					system.HallwayRow = rows;
				} else {
					// This is a letter, add a unit
					unit := &AmpUnit{};
					unit.PosX = cols;
					unit.PosY = rows;
					unit.UnitId = unitId;
					unitId++;
					unit.Letter = v;
					goalColumnFound := false;
					for _, c := range system.GoalColumns{
						if(c == cols){
							goalColumnFound = true;
							break;
						}
					}
					if(!goalColumnFound){
						system.GoalColumns = append(system.GoalColumns, cols);
					}
					system.Units = append(system.Units, unit);
				}
				if(v == ' '){
					system.Grid.SetValue(cols, rows, int('#'));
				} else{
					system.Grid.SetValue(cols, rows, int(v));
				}

				cols++;
			}
			rows++;
		}
	}

	//Log.Fatal("Parsed grid %d", system.HallwayRow);
	system.LowerGoal = system.HallwayRow + 2;
	system.UpperGoal = system.LowerGoal - 1;

	Log.Info("Parsed grid");
	fmt.Println(system.Grid.PrintAscii());

	frontier = append(frontier, system);
	this.BestEnergy = math.MaxInt32;
	statesExplored := 0;
	statesPrinted := 0;
	bestDepth := 0;
	for {
		if(len(frontier) <= 0){
			Log.Info("Search complete!");
			break;
		}
		next := frontier[0];
		statesExplored++;
		frontier = frontier[1:];
		if(next.EnergySpentSoFar >= this.BestEnergy){
			continue;
		}
		hash := next.Grid.PrintAscii();
		best, exists := this.BestCosts[hash];
		if(exists){
			if(best <= next.EnergySpentSoFar){
				continue;
			}
		}
		this.BestCosts[hash] = next.EnergySpentSoFar;
		if(next.Depth > bestDepth){
			bestDepth = next.Depth;
			statesPrinted++;
		}
		if(next.IsCompleteState()){
			if(next.EnergySpentSoFar < this.BestEnergy){
				this.BestEnergy = next.EnergySpentSoFar;
				Log.Info("New best energy %d", this.BestEnergy);
			}
			//fmt.Println(next.Grid.PrintAscii());
			continue;
		}
		successorStates := next.SuccessorStates();
		frontier = append(successorStates, frontier...)
	}

	Log.Info("Search complete, best energy is %d - states explored %d", this.BestEnergy, statesExplored);

}
