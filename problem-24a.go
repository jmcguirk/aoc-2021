package main

import (
	"math"
	"math/big"
	"os"
	"strconv"
)

type Problem24A struct {
	Registers []int64;
	MaxRegisters []int64;
	MinRegisters []int64;
	Machine *IntcodeMachine;
	LeastValue int64;
}

func (this *Problem24A) Solve() {
	Log.Info("Problem 24A solver beginning!")


	this.Machine = &IntcodeMachine{};
	err := this.Machine.Load("source-data/input-day-24a.txt");
	if(err != nil) {
		Log.FatalError(err)
	}


	/*
	for i := 0; i <= 10; i++{
		this.Machine.Reset();
		this.Machine.QueueInput(int64(i));
		this.Machine.Run();
		z := this.Machine.GetRegisterValue("z");
		y := this.Machine.GetRegisterValue("y");
		x := this.Machine.GetRegisterValue("x");
		w := this.Machine.GetRegisterValue("w");
		Log.Info("%d - %d%d%d%d", i, w,x,y,z);
	}*/



	this.LeastValue = int64(math.MaxInt64);
	this.Registers = make([]int64, 14);
	this.MinRegisters = make([]int64, 14);
	this.MaxRegisters = make([]int64, 14);


	min := big.NewInt(11111111111111);
	max := big.NewInt(99999999999999);



	this.LoadIntoBuff(this.Registers, min);
	this.LoadIntoBuff(this.MinRegisters, min);
	this.LoadIntoBuff(this.MaxRegisters, max);

	Log.Info("Starting a scan using %d registers", len(this.MaxRegisters));

	this.LoopOdometer();
}



func (this *Problem24A) LoadIntoBuff(registers []int64, bigInt *big.Int) {

	Log.Info("Loading %d " , bigInt.Int64());
	for i, _ := range registers{
		cpy := big.NewInt(0);
		cpy.Set(bigInt)
		registers[i] = nthDigit(cpy,int64(len(registers) - i - 1));
	}
}

func (this *Problem24A) LogRegisters() {
	this.LogBuff(this.Registers);
}

func (this *Problem24A) LogAndExit() {
	Log.Info("Completed odometer loop");
	os.Exit(0)
}


func (this *Problem24A) LoopOdometer() {

	for {
		atLim := false;
		for j := len(this.Registers) - 1; j >= 0; j--{
			if(this.Registers[j] + 1 < 9){
				this.Registers[j]++;
				break;
			} else{
				if(j == 0){
					atLim = true;
					break;
				}
				this.Registers[j] = 1;
			}
		}

		this.Machine.Reset();
		for _, v := range this.Registers{
			this.Machine.QueueInput(v);
		}
		success := this.Machine.Run();
		if(!success){
			Log.Fatal("Failed to execute machine");
		}
		output := this.Machine.GetRegisterValue("z");
		if(output < this.LeastValue){
			this.LeastValue = output;
			buff := "";
			for _, v := range this.Registers{
				buff += strconv.Itoa(int(v));
			}
			Log.Info(buff + " %d", output);
		}
		if(output == 0){
			Log.Info("Found terminal output!");
			this.LogRegisters();
			os.Exit(1);
		}
		if(atLim) {
			break;
		}
	}
}

func (this *Problem24A) LogBuff(registers[]int64) {

	buff := "";
	for _, v := range registers{
		buff += strconv.Itoa(int(v));
	}
	Log.Info(buff);
}