package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem3B struct {

}




func (this *Problem3B) Solve() {
	Log.Info("Problem 3B solver beginning!")


	file, err := os.Open("source-data/input-day-03b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	readOuts := make([]*SubmarineDiagnosticReadout, 0);

	colLen := 0;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			readOut := &SubmarineDiagnosticReadout{};
			readOut.Bytes = make([]byte, 0);
			for _, letter := range line {
				val, err := strconv.ParseInt(string(letter), 10, 64);
				if(err == nil){
					readOut.Bytes = append(readOut.Bytes, byte(val));
				}
			}
			readOuts = append(readOuts, readOut);
			colLen = len(readOut.Bytes);
		}
	}

	gamma := &SubmarineDiagnosticReadout{};
	gamma.Bytes = make([]byte, 0);

	epsilon := &SubmarineDiagnosticReadout{};
	epsilon.Bytes = make([]byte, 0);



	Log.Info("Finished parsing file - %d diagnostic readouts parsed. Col len is %d", len(readOuts), colLen);



	workingSet := make([]*SubmarineDiagnosticReadout, 0);
	workingSet = append(workingSet, readOuts...);

	index := 0;
	for{
		//Log.Info("Index %d - %d candidates remain", index, len(workingSet));
		if(len(workingSet) == 1){
			//Log.Info("Found 02 Readout");
			break;
		}


		zeroes := 0;
		ones := 0;
		luckyVal := byte(0);
		for _, readOut := range workingSet {
			if(readOut.Bytes[index] == 0){
				zeroes++;
			} else{
				ones++;
			}
		}
		if(ones >= zeroes){
			luckyVal = 1;
		}

		newWorkingSet := make([]*SubmarineDiagnosticReadout, 0);
		for _, readOut := range workingSet {
			if(readOut.Bytes[index] == luckyVal){
				newWorkingSet = append(newWorkingSet, readOut);
			}
		}
		workingSet = newWorkingSet;
		index++;
	}

	o2Value := workingSet[0].Value();

	Log.Info("Found 02Value Reading: %s", workingSet[0].ToString());

	workingSet = make([]*SubmarineDiagnosticReadout, 0);
	workingSet = append(workingSet, readOuts...);

	index = 0;
	for{
		//Log.Info("Index %d - %d candidates remain", index, len(workingSet));
		if(len(workingSet) == 1){
			//Log.Info("Found 02 Readout");
			break;
		}


		zeroes := 0;
		ones := 0;
		luckyVal := byte(0);
		for _, readOut := range workingSet {
			if(readOut.Bytes[index] == 0){
				zeroes++;
			} else{
				ones++;
			}
		}
		if(ones < zeroes){
			luckyVal = 1;
		}

		//Log.Info("Index %d - %d  is the lucky value", index, luckyVal);
		newWorkingSet := make([]*SubmarineDiagnosticReadout, 0);
		for _, readOut := range workingSet {
			if(readOut.Bytes[index] == luckyVal){
				newWorkingSet = append(newWorkingSet, readOut);
			}
		}
		workingSet = newWorkingSet;
		index++;
	}

	c02Value := workingSet[0].Value();
	Log.Info("Found C02Value Reading: %s", workingSet[0].ToString());

	checkSum := c02Value * o2Value;

	Log.Info("Checksum: %d", checkSum);
}
