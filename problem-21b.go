package main

import (
	"fmt"
)

type Problem21B struct {
	StateCache map[string]*DiceSystem;
	WinningScore int;
	Explored int;
	WinningsPerPlayer map[int]int64;
}

type DiceSystem struct {
	ActivePlayer int;

	Player1Position int;
	Player1Score int;

	Player2Position int;
	Player2Score int;

	WinningsPerPlayer map[int]int64;
}

func (this *DiceSystem) Clone() *DiceSystem {
	res := &DiceSystem{};

	res.ActivePlayer = this.ActivePlayer;

	res.Player1Position = this.Player1Position;
	res.Player2Position = this.Player2Position;
	res.Player1Score = this.Player1Score;
	res.Player2Score = this.Player2Score;

	res.WinningsPerPlayer = make(map[int]int64);
	//for k, v := range this.WinningsPerPlayer{
	//	res.WinningsPerPlayer[k] = 0;
	//}

	return res;
}

func (this *DiceSystem) StateHash() string {
	return fmt.Sprintf("%d-%d-%d-%d-%d", this.ActivePlayer, this.Player1Position, this.Player1Score, this.Player2Position, this.Player2Score)
}


func (this *Problem21B) Solve() {
	Log.Info("Problem 21B solver beginning!")
	this.WinningsPerPlayer = make(map[int]int64);
	this.StateCache = make(map[string]*DiceSystem);
	this.WinningScore = 21;

	initial := &DiceSystem{};
	initial.Player1Score = 0;
	initial.Player2Score = 0;
	initial.WinningsPerPlayer = make(map[int]int64);
	initial.WinningsPerPlayer[1] = 0;
	initial.WinningsPerPlayer[2] = 0;

	initial.Player1Position = 2;
	initial.Player2Position = 8;



	initial.Player1Position--;
	initial.Player2Position--;

	initial.ActivePlayer = 1;

	final := this.Explore(initial);

	Log.Info("Finished exploring dice system - explored %d", this.Explored);

	for k, v := range final.WinningsPerPlayer{
		Log.Info("%d => %d", k, v);
	}
}

func (this *Problem21B) Explore(system *DiceSystem) *DiceSystem {
	this.Explored++;
	//startExplore := this.Explored;
	hash := system.StateHash();
	cached, exists := this.StateCache[hash];
	if(exists){
		return cached;
	}

	if(system.Player1Score >= this.WinningScore){
		system.WinningsPerPlayer[1] = 1;
		this.StateCache[hash] = system;
		this.WinningsPerPlayer[1]++;
		return system;
	} else if(system.Player2Score >= this.WinningScore){
		system.WinningsPerPlayer[2] = 1;
		this.StateCache[hash] = system;
		this.WinningsPerPlayer[2]++;
		return system;
	}

	cs := system.Clone();
	childWinners := make(map[int]int64);
	childWinners[1] = 0;
	childWinners[2] = 0;
	for i := 1; i <= 3; i++{
		for j := 1; j <= 3; j++{
			for k := 1; k <= 3; k++{
				sum := i + j + k;
				childState := cs.Clone();
				if(childState.ActivePlayer == 1){
					newPos := childState.Player1Position + sum;
					newPos = newPos % 10;
					score := newPos + 1;
					childState.Player1Score += score;
					childState.Player1Position = newPos;
					childState.ActivePlayer = 2;
				} else{
					newPos := childState.Player2Position + sum;
					newPos = newPos % 10;
					score := newPos + 1;
					childState.Player2Score += score;
					childState.Player2Position = newPos;
					childState.ActivePlayer = 1;
				}
				explored := this.Explore(childState);
				for z, v := range explored.WinningsPerPlayer{
					childWinners[z] += v;
				}
			}
		}
	}
	system.WinningsPerPlayer = childWinners;
	this.StateCache[hash] = system;
	return system;
}