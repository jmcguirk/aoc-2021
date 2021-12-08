package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SevenSegmentSystem struct{
	Entries []*SevenSegmentLogEntry;
	Displays []*SevenSegmentDisplay;
	CanonicalLetters []rune;
	CanonicalPositions []int;
	CanonicalMapping map[rune]int;

	Zero []int;
	One []int;
	Two []int;
	Three []int;
	Four []int;
	Five []int;
	Six []int;
	Seven []int;
	Eight []int;
	Nine []int;

	AllNumbers [][]int;
}

type SevenSegmentDisplay struct{
	Index int;
	CurrentAssignments map[rune]int;
	ContainingSystem *SevenSegmentSystem;
	HasBeenSolved bool;
	PermutationsConsidered int;
}

type SevenSegmentLogEntry struct{
	Inputs []*SevenSegmentReading;
	Outputs[]*SevenSegmentReading;
	Row int;
}

type SevenSegmentReading struct{
	Runes []rune;
}

const SevenSegmentDisplay_Top = 1;
const SevenSegmentDisplay_TopLeft = 2;
const SevenSegmentDisplay_TopRight = 3;
const SevenSegmentDisplay_Center = 4;
const SevenSegmentDisplay_BottomLeft = 5;
const SevenSegmentDisplay_BottomRight = 6;
const SevenSegmentDisplay_Bottom = 7;


func(this *SevenSegmentDisplay) ToString() string{
	return "";
}

func(this *SevenSegmentSystem) Init(numDisplays int){
	this.Displays = make([]*SevenSegmentDisplay, numDisplays);
	for i, _ := range this.Displays{
		this.Displays[i] = &SevenSegmentDisplay{};
		this.Displays[i].Index = i;
		this.Displays[i].ContainingSystem = this;
		this.Displays[i].CurrentAssignments = make(map[rune]int);
	}
	this.Entries = make([]*SevenSegmentLogEntry, 0);
	this.CanonicalLetters = make([]rune, 7);
	this.CanonicalLetters[0] = 'a';
	this.CanonicalLetters[1] = 'b';
	this.CanonicalLetters[2] = 'c';
	this.CanonicalLetters[3] = 'd';
	this.CanonicalLetters[4] = 'e';
	this.CanonicalLetters[5] = 'f';
	this.CanonicalLetters[6] = 'g';

	this.CanonicalPositions = make([]int, 7);
	this.CanonicalPositions[0] = SevenSegmentDisplay_Top;
	this.CanonicalPositions[1] = SevenSegmentDisplay_TopLeft;
	this.CanonicalPositions[2] = SevenSegmentDisplay_TopRight;
	this.CanonicalPositions[3] = SevenSegmentDisplay_Center;
	this.CanonicalPositions[4] = SevenSegmentDisplay_BottomLeft;
	this.CanonicalPositions[5] = SevenSegmentDisplay_BottomRight;
	this.CanonicalPositions[6] = SevenSegmentDisplay_Bottom;

	this.CanonicalMapping = make(map[rune]int);
	this.CanonicalMapping['a'] = SevenSegmentDisplay_Top;
	this.CanonicalMapping['b'] = SevenSegmentDisplay_TopLeft;
	this.CanonicalMapping['c'] = SevenSegmentDisplay_TopRight;
	this.CanonicalMapping['d'] = SevenSegmentDisplay_Center;
	this.CanonicalMapping['e'] = SevenSegmentDisplay_BottomLeft;
	this.CanonicalMapping['f'] = SevenSegmentDisplay_BottomRight;
	this.CanonicalMapping['g'] = SevenSegmentDisplay_Bottom;


	this.Zero = make([]int, 6);
	this.Zero[0] = SevenSegmentDisplay_TopRight;
	this.Zero[1] = SevenSegmentDisplay_BottomRight;
	this.Zero[2] = SevenSegmentDisplay_TopLeft;
	this.Zero[3] = SevenSegmentDisplay_BottomLeft;
	this.Zero[4] = SevenSegmentDisplay_Bottom;
	this.Zero[5] = SevenSegmentDisplay_Top;

	this.One = make([]int, 2);
	this.One[0] = SevenSegmentDisplay_TopRight;
	this.One[1] = SevenSegmentDisplay_BottomRight;


	this.Two = make([]int, 5);
	this.Two[0] = SevenSegmentDisplay_TopRight;
	this.Two[1] = SevenSegmentDisplay_BottomLeft;
	this.Two[2] = SevenSegmentDisplay_Center;
	this.Two[3] = SevenSegmentDisplay_Bottom;
	this.Two[4] = SevenSegmentDisplay_Top;

	this.Three = make([]int, 5);
	this.Three[0] = SevenSegmentDisplay_TopRight;
	this.Three[1] = SevenSegmentDisplay_BottomRight;
	this.Three[2] = SevenSegmentDisplay_Center;
	this.Three[3] = SevenSegmentDisplay_Bottom;
	this.Three[4] = SevenSegmentDisplay_Top;

	this.Four = make([]int, 4);
	this.Four[0] = SevenSegmentDisplay_TopRight;
	this.Four[1] = SevenSegmentDisplay_BottomRight;
	this.Four[2] = SevenSegmentDisplay_Center;
	this.Four[3] = SevenSegmentDisplay_TopLeft;

	this.Five = make([]int, 5);
	this.Five[0] = SevenSegmentDisplay_Top;
	this.Five[1] = SevenSegmentDisplay_TopLeft;
	this.Five[2] = SevenSegmentDisplay_Center;
	this.Five[3] = SevenSegmentDisplay_BottomRight;
	this.Five[4] = SevenSegmentDisplay_Bottom;

	this.Six = make([]int, 6);
	this.Six[0] = SevenSegmentDisplay_Top;
	this.Six[1] = SevenSegmentDisplay_TopLeft;
	this.Six[2] = SevenSegmentDisplay_Center;
	this.Six[3] = SevenSegmentDisplay_BottomRight;
	this.Six[4] = SevenSegmentDisplay_Bottom;
	this.Six[5] = SevenSegmentDisplay_BottomLeft;

	this.Seven = make([]int, 3);
	this.Seven[0] = SevenSegmentDisplay_Top;
	this.Seven[1] = SevenSegmentDisplay_TopRight;
	this.Seven[2] = SevenSegmentDisplay_BottomRight;

	this.Eight = make([]int, 7);
	this.Eight[0] = SevenSegmentDisplay_Top;
	this.Eight[1] = SevenSegmentDisplay_TopRight;
	this.Eight[2] = SevenSegmentDisplay_TopLeft;
	this.Eight[3] = SevenSegmentDisplay_Bottom;
	this.Eight[4] = SevenSegmentDisplay_BottomRight;
	this.Eight[5] = SevenSegmentDisplay_BottomLeft;
	this.Eight[6] = SevenSegmentDisplay_Center;

	this.Nine = make([]int, 6);
	this.Nine[0] = SevenSegmentDisplay_Top;
	this.Nine[1] = SevenSegmentDisplay_TopRight;
	this.Nine[2] = SevenSegmentDisplay_TopLeft;
	this.Nine[3] = SevenSegmentDisplay_Bottom;
	this.Nine[4] = SevenSegmentDisplay_BottomRight;
	this.Nine[5] = SevenSegmentDisplay_Center;

	this.AllNumbers = make([][]int, 0);
	this.AllNumbers = append(this.AllNumbers, this.Zero);
	this.AllNumbers = append(this.AllNumbers, this.One);
	this.AllNumbers = append(this.AllNumbers, this.Two);
	this.AllNumbers = append(this.AllNumbers, this.Three);
	this.AllNumbers = append(this.AllNumbers, this.Four);
	this.AllNumbers = append(this.AllNumbers, this.Five);
	this.AllNumbers = append(this.AllNumbers, this.Six);
	this.AllNumbers = append(this.AllNumbers, this.Seven);
	this.AllNumbers = append(this.AllNumbers, this.Eight);
	this.AllNumbers = append(this.AllNumbers, this.Nine);
}

func(this *SevenSegmentReading) Parse(line string){
	this.Runes = make([]rune, 0);
	for _, v := range line{
		this.Runes = append(this.Runes, v);
	}
}

func(this *SevenSegmentSystem) Solve() int{
	sum := 0;
	for _, v := range this.Displays {
		for _, ent := range this.Entries{
			v.HasBeenSolved = false;
			v.PermutationsConsidered = 0;
			v.CurrentAssignments = make(map[rune]int);
			sum += v.Solve(ent,0);
			if(!v.HasBeenSolved){
				Log.Fatal("Failed to solve display - %d permutations considered", v.PermutationsConsidered);
			}
		}

	}
	Log.Info("Solved all log entries");
	return sum;
}

func(this *SevenSegmentSystem) CountUniques() int{
	total := 0
	for _, v := range this.Entries{
		for _, output := range v.Outputs {
			l := len(output.Runes);
			if(l == 2 || l == 3 || l == 4 || l == 7){
				total++;
			}
		}
	}
	return total;
}


func(this *SevenSegmentDisplay) Solve(entry *SevenSegmentLogEntry, index int) int{
	sum := 0;
	didAssign := false;
	for _, c := range this.ContainingSystem.CanonicalLetters {
		_, exists := this.CurrentAssignments[c];
		if(exists){
			continue;
		}
		didAssign = true;
		this.CurrentAssignments[c] = this.ContainingSystem.CanonicalPositions[index];
		if(index == (len(this.ContainingSystem.CanonicalPositions) - 1)){
			this.PermutationsConsidered++;
			if(this.VerifyEntry(entry)){
				buff := "";
				for _, v := range entry.Outputs{
					buff += fmt.Sprintf("%d", this.Map(v));
				}
				val, err := strconv.Atoi(buff);
				if(err != nil){
					Log.FatalError(err);
				}
				Log.Info("Solved entry %d after %d permutations. Value is %d", entry.Row, this.PermutationsConsidered, val);
				this.HasBeenSolved = true;
				sum = val;
				return sum;
			}
		} else{
			sum = this.Solve(entry, index+1);
			if(this.HasBeenSolved){
				return sum;
			}
		}
		delete(this.CurrentAssignments, c);
	}
	if(!didAssign){
		Log.Fatal("Failed to assign during search %d", index);
	}
	return -1;
}

func(this *SevenSegmentDisplay) Map(reading *SevenSegmentReading) int{
	mapped := make([]int, len(reading.Runes));
	for i, v := range reading.Runes {
		mappedV, _ := this.CurrentAssignments[v];
		mapped[i] = mappedV;
	}

	if(this.HasAll(mapped, this.ContainingSystem.Zero)){
		return 0;
	}
	if(this.HasAll(mapped, this.ContainingSystem.One)){
		return 1;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Two)){
		return 2;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Three)){
		return 3;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Four)){
		return 4;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Five)){
		return 5;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Six)){
		return 6;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Seven)){
		return 7;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Eight)){
		return 8;
	}
	if(this.HasAll(mapped, this.ContainingSystem.Nine)){
		return 9;
	}

	Log.Fatal("Failed to map output value");
	return -1;
}

func(this *SevenSegmentDisplay) VerifyEntry(entry *SevenSegmentLogEntry) bool{
	for _, v := range entry.Inputs{
		if(!this.VerifyInput(v)){
			return false;
		}
	}
	return true;
}

func(this *SevenSegmentDisplay) VerifyInput(reading *SevenSegmentReading) bool{

	mapped := make([]int, len(reading.Runes));
	for i, v := range reading.Runes {
		mappedV, _ := this.CurrentAssignments[v];
		mapped[i] = mappedV;
	}

	for _, num := range this.ContainingSystem.AllNumbers{
		if(len(num) == len(mapped)){
			if(this.HasAll(mapped, num)){
				return true;
			}
		}
	}
	return false;
}


func(this *SevenSegmentDisplay) HasAll(v []int, target []int) bool{
	if(len(v) != len(target)){
		return false;
	}
	for _, tV := range target{
		found := false;
		for _, vV := range v{
			if(tV == vV){
				found = true;
				break;
			}
		}
		if(!found){
			return false;
		}
	}
	return true;
}

func(this *SevenSegmentDisplay) Log(){
	for k, v := range this.CurrentAssignments {
		Log.Info("%c, %v", k, v);
	}
}

func(this *SevenSegmentSystem) AddEntry(line string){
	lineParts := strings.Split(line, "|");
	if(len(lineParts) != 2){
		Log.Fatal("Incorrect number of line parts");
	}
	entry := &SevenSegmentLogEntry{};
	entry.Inputs = make([]*SevenSegmentReading, 0);
	entry.Outputs = make([]*SevenSegmentReading, 0);
	inputs := strings.Split(strings.TrimSpace(lineParts[0]), " ");
	for _, v := range inputs{
		if(v != ""){
			reading := &SevenSegmentReading{};
			reading.Parse(v);
			entry.Inputs = append(entry.Inputs, reading);
		}
	}
	outputs := strings.Split(strings.TrimSpace(lineParts[1]), " ");
	for _, v := range outputs{
		if(v != ""){
			reading := &SevenSegmentReading{};
			reading.Parse(v);
			entry.Outputs = append(entry.Outputs, reading);
		}
	}
	entry.Row = len(this.Entries);
	this.Entries = append(this.Entries, entry);
}

