package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem4A struct {

}


type BingoBoard struct {
	Id 		int;
	Values *IntegerGrid2D;
	Hits map[int]bool;
	Cols int;
	Rows int;
	HasBingo bool;
}

func (this *BingoBoard) Init(id int) {
	this.Id = id;
	this.Values = &IntegerGrid2D{};
	this.Values.Init();
	this.Hits = make(map[int]bool);
}


func (this *BingoBoard) Mark(value int) {
	this.Hits[value] = true;
}

func (this *BingoBoard) IsMarked(value int) bool {
	val, _ := this.Hits[value];
	return val;
}

func (this *BingoBoard) Checksum(bingoNumber int) int {

	unMarkedSum := 0;

	for i := 0; i < this.Cols; i++{
		for j := 0; j < this.Rows; j++{
			if(!this.IsCellMarked(j, i)){
				unMarkedSum += this.Values.GetValue(j, i);
			}
		}
	}

	return unMarkedSum * bingoNumber;
}

func (this *BingoBoard) IsCellMarked(row int, col int) bool {
	return this.IsMarked(this.Values.GetValue(row, col));
}

func (this *BingoBoard) IsBingo() bool {
	for i := 0; i < this.Cols; i++{
		allCol := true;
		for j := 0; j < this.Rows; j++{
			if(!this.IsCellMarked(j, i)){
				allCol = false;
				break;
			}
		}
		if(allCol){
			return true;
		}
	}

	for j := 0; j < this.Rows; j++{
		allRow := true;
		for i := 0; i < this.Cols; i++{
			if(!this.IsCellMarked(j, i)){
				allRow = false;
				break;
			}
		}
		if(allRow){
			return true;
		}
	}

	return false;
}

func (this *BingoBoard) Log() {
	Log.Info("");
	Log.Info("Bingo Board %d (%d rows, %d cols)", this.Id, this.Rows, this.Cols)
	Log.Info("");

	for i := 0; i < this.Cols; i++{

		row := "";
		for j := 0; j < this.Rows; j++{
			val := this.Values.GetValue(j, i);
			if(this.IsMarked(val)){
				row += "*";
			} else{
				row += " ";
			}
			row += fmt.Sprintf("%d", val);
			if(this.IsMarked(val)){
				row += "*";
			} else{
				row += " ";
			}
		}
		Log.Info(row);
		Log.Info("");
	}
}


func (this *Problem4A) Solve() {
	Log.Info("Problem 4A solver beginning!")


	file, err := os.Open("source-data/input-day-04a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hasParsedBingoLine := false;
	var board *BingoBoard;
	rows := 0;
	boardId := 0;
	bingoNumbers := make([]int, 0);

	allBoards := make([]*BingoBoard, 0);

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			if(!hasParsedBingoLine){
				hasParsedBingoLine = true;
				bingoParts := strings.Split(line, ",");
				for _, v := range bingoParts{
					parsed, err := strconv.Atoi(v);
					if(err != nil){
						Log.FatalError(err);
					}
					bingoNumbers = append(bingoNumbers, parsed);
				}
			} else{
				if(board == nil){
					board = &BingoBoard{};
					rows = 0;
					board.Init(boardId);
					boardId++;
					allBoards = append(allBoards, board);
				}
				cols := 0;
				lineParts := strings.Split(line, " ");
				for _, v := range lineParts{
					if(v == ""){
						continue;
					}
					parsed, err := strconv.Atoi(v);
					if(err != nil){
						Log.FatalError(err);
					}
					board.Values.SetValue(cols, rows, parsed);
					cols++;
				}
				board.Cols = cols;
				rows++;
			}
		} else{
			if(board != nil){
				board.Rows = rows;
				board = nil;
			}
		}
	}

	if(board != nil){
		board.Rows = rows;
		board = nil;
	}

	Log.Info("Finished parsing file. Processing %d bingo numbers and %d bingo boards", len(bingoNumbers), len(allBoards));

	bingoFound := false;
	for _, bingoValue := range bingoNumbers{
		if(bingoFound){
			break;
		}
		for _, markedBoard := range allBoards{
			markedBoard.Mark(bingoValue);
		}
		for _, markedBoard := range allBoards{
			if(markedBoard.IsBingo()){
				Log.Info("Bingo found with number %d checksum is %d", bingoValue, markedBoard.Checksum(bingoValue));
				//markedBoard.Log();
				bingoFound = true;
			}
		}
	}

	Log.Info("Finished bingo simulation found bingo %t", bingoFound)

}
