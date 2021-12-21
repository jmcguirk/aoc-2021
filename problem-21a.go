package main

type Problem21A struct {
	DiceIndex int;
}

type DiceRacer struct {
	PlayerId int;
	Score int;
	Position int;
}

func (this *Problem21A) Roll() int {
	this.DiceIndex++;
	return this.DiceIndex;
}

func (this *Problem21A) Solve() {
	Log.Info("Problem 12A solver beginning!")

	racers := make([]*DiceRacer, 0);

	p1 := &DiceRacer{};
	p1.PlayerId = 1;
	p1.Position = 2;
	racers = append(racers, p1);

	p2 := &DiceRacer{};
	p2.PlayerId = 2;
	p2.Position = 8;
	racers = append(racers, p2);

	activePlayer := 0;
	stepCount := 0;
	targetScore := 1000;

	// Normalize input
	for _, player := range racers {
		player.Position--;
	}

	winningPlayerId := -1;

	for {
		nextPlayer := racers[activePlayer]

		diceSum := this.Roll() + this.Roll() + this.Roll();

		newPos := nextPlayer.Position + diceSum;
		newPos = newPos % 10;
		score := newPos + 1;
		nextPlayer.Position = newPos;
		nextPlayer.Score += score;

		//Log.Info("%d.) Player %d moves %d steps and scores %d points", stepCount, nextPlayer.PlayerId, diceSum, score);

		if(nextPlayer.Score >= targetScore){
			winningPlayerId = nextPlayer.PlayerId;
			Log.Info("Game concluded after %d steps, %d player is the winner with score %d", stepCount, nextPlayer.PlayerId, nextPlayer.Score);
			break;
		}

		activePlayer = activePlayer + 1;
		activePlayer = activePlayer % len(racers);
		stepCount++;
	}

	losingScore := 0;

	for _, player := range racers{
		if(player.PlayerId != winningPlayerId){
			losingScore = player.Score;
		}
	}
	Log.Info("Checksum is %d x %d = %d", losingScore, this.DiceIndex, losingScore * this.DiceIndex)

}
