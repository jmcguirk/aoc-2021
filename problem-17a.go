package main

import "math"

type Problem17A struct {
	BestHeight int;

	MinX int;
	MaxX int;
	MinY int;
	MaxY int;

	MaxSimulationSteps int;
}

func (this *Problem17A) Solve() {
	Log.Info("Problem 17A solver beginning!")


	/*
	this.MinX = 20;
	this.MaxX = 30;

	this.MinY = -10;
	this.MaxY = -5;
	*/

	this.MinX = 115;
	this.MaxX = 215;

	this.MinY = -132;
	this.MaxY = -72;


	this.MaxSimulationSteps = 20 * (this.MaxX - this.MinX);

	minVx := -1000;
	maxVx := 1000;

	minVy := -1000;
	maxVy := 1000;

	peakHeight := math.MinInt32;

	for x := minVx; x <= maxVx; x++{
		for y := minVy; y <= maxVy; y++{
			peak, hit := this.Simulate(x,y);
			if(hit){
				if(peak > peakHeight){
					peakHeight = peak;
				}
			}
		}
	}

	Log.Info("Completed simulation, peak height was %d ", peakHeight);

}

func (this *Problem17A) Simulate(velocityX int, velocityY int) (int, bool) {

	vX := velocityX;
	vY := velocityY;
	currX := 0;
	currY := 0;
	stepCount := 0;
	peakY := math.MinInt32;

	for {
		if(currX >= this.MinX && currX <= this.MaxX && currY >= this.MinY && currY <= this.MaxY){
			return peakY, true;
		}

		if(currY < this.MinY && vY <= 0){
			return -1, false;
		}

		if(currX < this.MinX && vX <= 0){
			return -1, false;
		}

		if(currX > this.MaxX && vX >= 0){
			return -1, false;
		}

		if(stepCount >= this.MaxSimulationSteps){
			return 0, false;
		}

		currX += vX;
		currY += vY;

		if(currY > peakY){
			peakY = currY;
		}

		vY--;
		if(vX > 0){
			vX--;
		} else if(vX < 0){
			vX++;
		}

		stepCount++;
	}
}
