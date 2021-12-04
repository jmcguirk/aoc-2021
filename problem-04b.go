package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem4B struct {

}



func (this *Problem4B) Solve() {
	Log.Info("Problem 4B solver beginning!")


	file, err := os.Open("source-data/input-day-04b.txt");
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

	var lastWinner *BingoBoard;
	lastWinnerNumber := 0;

	for _, bingoValue := range bingoNumbers{
		for _, markedBoard := range allBoards{
			if(!markedBoard.HasBingo){
				markedBoard.Mark(bingoValue);
			}
		}
		for _, markedBoard := range allBoards{
			if(!markedBoard.HasBingo && markedBoard.IsBingo()){
				//Log.Info("Bingo found with number %d checksum is %d", bingoValue, markedBoard.Checksum(bingoValue));
				//markedBoard.Log();
				lastWinner = markedBoard
				lastWinnerNumber = bingoValue;
				markedBoard.HasBingo = true;
			}
		}
	}
	if(lastWinner != nil){
		Log.Info("Found last winner with number %d, checksum %d", lastWinnerNumber, lastWinner.Checksum(lastWinnerNumber))
	}

}
