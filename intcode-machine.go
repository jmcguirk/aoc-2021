package main

import (
	"bufio"
	"os"
	"strings"
)

type IntcodeMachine struct {
	Instructions[] *IntcodeInstruction;
	InstructionPointer int64;
	Registers map[string]int64;
	InputQueue []int64;
}



func (this *IntcodeMachine) QueueInput(value int64) {
	this.InputQueue = append(this.InputQueue, value);
}

func (this *IntcodeMachine) GetRegisterValue(name string) int64 {
	v, exists := this.Registers[name];
	if(!exists){
		return 0;
	}
	return v;
}

func (this *IntcodeMachine) SetRegisterValue(name string, val int64) {
	this.Registers[name] = val;
}

func (this *IntcodeMachine) Reset() {
	this.InstructionPointer = 0;
	this.Registers = make(map[string]int64);
	this.InputQueue = make([]int64, 0);
}

func (this *IntcodeMachine) Load(fileName string) error {
	//Log.Info("Loading intcode v3 machine from %s", fileName)
	this.Instructions = make([]*IntcodeInstruction, 0);
	this.Reset();
	file, err := os.Open(fileName);
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0;
	for scanner.Scan() {
		trimmed := strings.TrimSpace(scanner.Text());
		if(trimmed != ""){
			lineNumber ++;
			instruction := &IntcodeInstruction{}
			instruction.Line = lineNumber;
			instruction.Parse(trimmed);
			this.Instructions = append(this.Instructions, instruction);
		}
	}

	Log.Info("Loaded program, contains %d lines", len(this.Instructions));
	this.InstructionPointer = 0;

	return nil;
}

func (this *IntcodeMachine) Run() (bool) {
	for{
		nextInstruction := this.Instructions[this.InstructionPointer];
		//Log.Info("Executing %s", nextInstruction.Describe())
		this.ExecuteInstruction(nextInstruction);
		if(int(this.InstructionPointer) >= len(this.Instructions)){
			return true
		}
	}
}

func (this *IntcodeMachine) ExecuteInstruction(instruction *IntcodeInstruction) {
	switch(instruction.Operation){
		case IntCodeOpCodeInput:
			this.ExecuteInput(instruction);
		case IntCodeOpCodeAdd:
			this.ExecuteAdd(instruction);
		case IntCodeOpCodeMul:
			this.ExecuteMul(instruction);
		case IntCodeOpCodeDiv:
			this.ExecuteDiv(instruction);
		case IntCodeOpCodeMod:
			this.ExecuteMod(instruction);
		case IntCodeOpCodeEql:
			this.ExecuteEql(instruction);
		default:
			Log.Fatal("Couldn't execute unknown instruction %d at line %d", instruction.Operation, instruction.Line)
	}
}

func (this *IntcodeMachine) ExecuteInput(instruction *IntcodeInstruction) {

	if(len(this.InputQueue) <= 0){
		Log.Fatal("Tried to process an input queue instruction and no input was provided");
	}

	next := this.InputQueue[0];
	this.InputQueue = this.InputQueue[1:];
	this.SetRegisterValue(instruction.ParameterOne, next);
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteAdd(instruction *IntcodeInstruction) {

	arg1 := this.GetRegisterValue(instruction.ParameterOne);

	arg2 := instruction.ParameterTwoInt;
	if(instruction.ParameterTwo != ""){
		arg2 = this.GetRegisterValue(instruction.ParameterTwo);
	}


	this.SetRegisterValue(instruction.ParameterOne, arg1 + arg2);
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteMul(instruction *IntcodeInstruction) {

	arg1 := this.GetRegisterValue(instruction.ParameterOne);

	arg2 := instruction.ParameterTwoInt;
	if(instruction.ParameterTwo != ""){
		arg2 = this.GetRegisterValue(instruction.ParameterTwo);
	}


	this.SetRegisterValue(instruction.ParameterOne, arg1 * arg2);
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteDiv(instruction *IntcodeInstruction) {

	arg1 := this.GetRegisterValue(instruction.ParameterOne);

	arg2 := instruction.ParameterTwoInt;
	if(instruction.ParameterTwo != ""){
		arg2 = this.GetRegisterValue(instruction.ParameterTwo);
	}


	this.SetRegisterValue(instruction.ParameterOne, arg1 / arg2);
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteMod(instruction *IntcodeInstruction) {

	arg1 := this.GetRegisterValue(instruction.ParameterOne);

	arg2 := instruction.ParameterTwoInt;
	if(instruction.ParameterTwo != ""){
		arg2 = this.GetRegisterValue(instruction.ParameterTwo);
	}


	this.SetRegisterValue(instruction.ParameterOne, arg1 % arg2);
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteEql(instruction *IntcodeInstruction) {

	arg1 := this.GetRegisterValue(instruction.ParameterOne);

	arg2 := instruction.ParameterTwoInt;
	if(instruction.ParameterTwo != ""){
		arg2 = this.GetRegisterValue(instruction.ParameterTwo);
	}

	val := 0;
	if(arg1 == arg2){
		val = 1
	}

	this.SetRegisterValue(instruction.ParameterOne, int64(val));
	this.InstructionPointer++;
}

func (this *IntcodeMachine) Describe() {
	Log.Info("Loaded a machine with %d instructions", len(this.Instructions));
	for _, v := range this.Instructions{
		Log.Info(v.Describe());
	}
}