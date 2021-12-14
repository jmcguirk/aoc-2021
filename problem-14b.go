package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Problem14B struct {
	Rules map[rune]map[rune]rune;
	PairMap map[rune]map[rune]int;
}



func (this *Problem14B) Solve() {
	Log.Info("Problem 14B solver beginning!")

	file, err := os.Open("source-data/input-day-14b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	this.Rules = make(map[rune]map[rune]rune);

	scanner := bufio.NewScanner(file)

	hasParsedInitialState := false;

	flatState := make([]rune, 0);
	ruleCount := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			if(!hasParsedInitialState){
				for _, v := range line {
					flatState = append(flatState, v);
				}
				hasParsedInitialState = true;
			} else{
				ruleParts := strings.Split(line, "->");
				inputParts := strings.TrimSpace(ruleParts[0]);
				output := rune(strings.TrimSpace(ruleParts[1])[0]);
				input1 := rune(inputParts[0]);
				input2 := rune(inputParts[1]);

				existing, exists := this.Rules[input1];
				if(!exists){
					existing = make(map[rune]rune);
					this.Rules[input1] = existing;
				}
				existing[input2] = output;
				ruleCount++;
			}
		}
	}

	Log.Info("Parsed initial state %d", len(flatState));
	Log.Info("Loaded %d rules", ruleCount)

	this.PairMap = make(map[rune]map[rune]int);

	for i, v := range flatState {
		if(i < len(flatState) - 1){
			existing, exists := this.PairMap[v];
			if(!exists){
				existing = make(map[rune]int);
				this.PairMap[v] = existing;
			}
			cnt, _ := existing[flatState[i+1]]
			cnt++;
			existing[flatState[i+1]] = cnt;
		}
	}

	currStep := 0;
	maxStep := 40;
	for{
		if(currStep >= maxStep){
			break;
		}


		newMap := make(map[rune]map[rune]int);





		for input1, inputs := range this.PairMap {
			for input2, count := range inputs {
				recipes, exists := this.Rules[input1];
				outputFound := false;
				if(exists){
					output, outputExists := recipes[input2];
					if(outputExists){
						outputFound = true;
						existingOutputsLeft, existingOutputsLeftExists := newMap[input1];
						if(!existingOutputsLeftExists){
							existingOutputsLeft = make(map[rune]int);
							newMap[input1] = existingOutputsLeft;
						}
						cnt, _ := existingOutputsLeft[output];
						cnt+= count;
						existingOutputsLeft[output] = cnt;

						existingOutputsRight, existingOutputsRightExists := newMap[output];
						if(!existingOutputsRightExists){
							existingOutputsRight = make(map[rune]int);
							newMap[output] = existingOutputsRight;
						}
						cnt, _ = existingOutputsRight[input2];
						cnt+= count;
						existingOutputsRight[input2] = cnt;
					}
				}
				if(!outputFound){
					Log.Fatal("Failed to convert");
				}
			}
		}

		this.PairMap = newMap;

		currStep++;
	}

	freq := make(map[rune]int);
	for _, v := range this.PairMap {
		for k, v2 := range v {
			cnt, _ := freq[k]
			cnt+=v2;
			freq[k] = cnt;
		}
	}

	cnt, _ := freq[flatState[0]]; // First letter needs +1
	cnt+=1;
	freq[flatState[0]] = cnt;

	flat := make([]*PolymerFrequency, 0);
	for k, v := range freq{
		f := &PolymerFrequency{};
		f.Letter = k;
		f.Count = v;
		flat = append(flat, f);
	}
	sort.SliceStable(flat, func(i, j int) bool {
		return flat[i].Count < flat[j].Count;
	});
	least := flat[0].Count;
	most := flat[len(flat) - 1].Count;
	checksum := most - least;
	Log.Info("Checksum is %d", checksum)
}

