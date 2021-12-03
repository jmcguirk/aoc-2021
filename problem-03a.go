package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem3A struct {

}

type SubmarineDiagnosticReadout struct {
	Bytes []byte;
}

func (this *SubmarineDiagnosticReadout) ToString() string {
	str := "";
	for _, v := range this.Bytes{
		str += fmt.Sprintf("%d", v);
	}
	return str + " - " + fmt.Sprintf("%d", this.Value());
}

func (this *SubmarineDiagnosticReadout) Value() int64 {
	str := "";
	for _, v := range this.Bytes{
		str += fmt.Sprintf("%d", v);
	}
	val, _ := strconv.ParseInt(str, 2, 64);
	return val;
}

func (this *Problem3A) Solve() {
	Log.Info("Problem 3A solver beginning!")


	file, err := os.Open("source-data/input-day-03a.txt");
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

	for i := 0; i < colLen; i++{
		zeroes := 0;
		ones := 0;
		for _, readOut := range readOuts{
			if(readOut.Bytes[i] == 0){
				zeroes++;
			} else{
				ones++;
			}
		}
		if(ones > zeroes){
			gamma.Bytes = append(gamma.Bytes, 1);
			epsilon.Bytes = append(epsilon.Bytes, 0);
		} else{
			gamma.Bytes = append(gamma.Bytes, 0);
			epsilon.Bytes = append(epsilon.Bytes, 1);
		}
	}


	Log.Info("Finished parsing file - %d diagnostic readouts parsed. Col len is %d", len(readOuts), colLen);

	gammaV := gamma.Value();
	epsilonV := epsilon.Value();
	product := gammaV * epsilonV;

	Log.Info("Gamma: %d", gammaV);
	Log.Info("Epsilon: %d", epsilonV);

	Log.Info("Checksum: %d", product);
}
