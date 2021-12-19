package main

import (
	"fmt"
	"math"
)

type Problem17B struct {
	BestHeight int;

	MinX int;
	MaxX int;
	MinY int;
	MaxY int;

	MaxSimulationSteps int;
}

func (this *Problem17B) Solve() {
	Log.Info("Problem 17B solver beginning!")


	/*
	this.MinX = 20;
	this.MaxX = 30;

	this.MinY = -10;
	this.MaxY = -5;
	*/

		this.MinX = 155;
		this.MaxX = 215;

		this.MinY = -132;
		this.MaxY = -72;

		/*
	this.MinX = 282;
	this.MaxX = 314;

	this.MinY = -80;
	this.MaxY = -45;*/


		this.MaxSimulationSteps = 20 * (this.MaxX - this.MinX);

		minVx := 0;
		maxVx := 1000;

		minVy := -1000;
		maxVy := 1000;

		peakHeight := math.MinInt32;
		hitCount := 0;

		seen := make(map[string]int);

		for x := minVx; x <= maxVx; x++{
			for y := minVy; y <= maxVy; y++{
				peak, hit := this.Simulate(x,y);
				if(hit){

					str := fmt.Sprintf("%d,%d", x, y);
					//Log.Info(str);
					_, exists := seen[str];
					if(!exists){
						seen[str] = 1;
						hitCount++;
					}

					if(peak > peakHeight){
						peakHeight = peak;
					}
				}
			}
		}

		// 9601 too many
		Log.Info("Completed simulation, peak height was %d, %d velocities hit ", peakHeight, hitCount);

	}

	func (this *Problem17B) Simulate(velocityX int, velocityY int) (int, bool) {

		vX := velocityX;
		vY := velocityY;
		currX := 0;
		currY := 0;
		stepCount := 0;
		peakY := math.MinInt32;

		for {



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

			if(currX >= this.MinX && currX <= this.MaxX && currY >= this.MinY && currY <= this.MaxY){
				return peakY, true;
			}



			stepCount++;
		}
	}
