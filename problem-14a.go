package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Problem14A struct {
	Head *PolymerNode;
	Rules map[rune]map[rune]rune;
}

type PolymerNode struct {
	Letter rune;
	Next *PolymerNode;
}

type PolymerFrequency struct {
	Letter rune;
	Count int;
}

func (this *Problem14A) PrintChain() string {
	buff := "";
	next := this.Head;
	for{
		if(next == nil){
			break;
		}
		buff += fmt.Sprintf("%c", next.Letter);
		next = next.Next;
	}
	return buff;
}


func (this *Problem14A) Solve() {
	Log.Info("Problem 14A solver beginning!")

	file, err := os.Open("source-data/input-day-14a-test.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	this.Rules = make(map[rune]map[rune]rune);

	scanner := bufio.NewScanner(file)

	hasParsedInitialState := false;

	ruleCount := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			if(!hasParsedInitialState){

				var prev *PolymerNode;
				for i, v := range line {
					node := &PolymerNode{};
					if(i == 0){
						this.Head = node;
					}
					node.Letter = v;
					if(prev != nil){
						prev.Next = node;
					}
					prev = node;
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

	Log.Info("Parsed initial state %s", this.PrintChain());
	Log.Info("Loaded %d rules", ruleCount)
	/*
	for k, r := range this.Rules{
		for k2, v := range r {
			Log.Info("%c%c -> %c", k, k2, v);
		}
	}*/

	stepCount := 0;
	maxStepCount := 10;
	lastLen := 0;
	for{
		if(stepCount >= maxStepCount){
			break;
		}

		curr := this.Head;
		lastLen = 0;
		for{
			if(curr == nil){
				break;
			}
			lastLen++;
			next := curr.Next;
			if(next == nil){
				break;
			}

			currL := curr.Letter;
			nextL := next.Letter;

			recipe, exists := this.Rules[currL];
			if(exists){
				output, outputExists := recipe[nextL];
				if(outputExists){
					reaction := &PolymerNode{};
					reaction.Letter = output;
					reaction.Next = next;
					curr.Next = reaction;
					lastLen++;
				}
			}

			curr = next;
		}

		stepCount++;
	}
	Log.Info("Finished simulation of %d steps - last chain length %d", stepCount, lastLen);

	next := this.Head;
	freq := make(map[rune]int);
	for{
		if(next == nil){
			break;
		}
		cnt, _ := freq[next.Letter];
		cnt++;
		freq[next.Letter] = cnt;
		next = next.Next;
	}

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

