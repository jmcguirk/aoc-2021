package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem2A struct {

}

func (this *Problem2A) Solve() {
	Log.Info("Problem 2A solver beginning!")


	file, err := os.Open("source-data/input-day-02a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	pos := &IntVec3{};
	pos.X = 0;
	pos.Z = 0;

	scanner := bufio.NewScanner(file)

	instructions := 0;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){

			parts := strings.Split(line, " ");
			if(len(parts) != 2){
				Log.Fatal("Failed to parse line " + line)
			}
			val, err := strconv.ParseInt(parts[1], 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			valI := int(val);
			dir := parts[0];
			switch(dir){
				case "forward":
					pos.X += valI;
					break;
				case "down":
					pos.Z += valI;
					break;
				case "up":
					pos.Z -= valI;
					break;
				default:
					Log.Fatal("Unhandled instruction %s", dir);
					break;
			}
			instructions++;
		}
	}

	Log.Info("Finished parsing file - %d instructions handled, final pos is %s. Checksum is %d", instructions, pos.ToString(), pos.X * pos.Z);
}
