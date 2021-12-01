package main

import (
	"bufio"
	"os"
	"strings"
)

type IntcodeMachine struct {
	Instructions[] *IntcodeInstruction;
	AccumulatedValue int64;
	InstructionPointer int64;
	ExecutedInstructions map[int]bool;
}






func (this *IntcodeMachine) Load(fileName string) error {
	//Log.Info("Loading intcode v3 machine from %s", fileName)
	this.Instructions = make([]*IntcodeInstruction, 0);
	this.ExecutedInstructions = make(map[int]bool, 0);
	file, err := os.Open(fileName);
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()


	graph := DirectedGraph{};
	graph.Init();
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

func (this *IntcodeMachine) RunTillFirstDuplicate() (int64, bool) {
	for{
		nextInstruction := this.Instructions[this.InstructionPointer];
		_, exists := this.ExecutedInstructions[nextInstruction.Line];
		if(exists){
			return this.AccumulatedValue, false;
		}
		//Log.Info("Executing %s", nextInstruction.Describe())
		this.ExecutedInstructions[nextInstruction.Line] = true;
		this.ExecuteInstruction(nextInstruction);
		if(int(this.InstructionPointer) >= len(this.Instructions)){
			return this.AccumulatedValue, true;
		}
	}
}

func (this *IntcodeMachine) ExecuteInstruction(instruction *IntcodeInstruction) {
	switch(instruction.Operation){
		case IntCodeOpCodeNoOp:
			this.ExecuteNoOp(instruction);
		case IntCodeOpCodeJump:
			this.ExecuteJump(instruction);
		case IntCodeOpCodeInc:
			this.ExecuteInc(instruction);
		default:
			Log.Fatal("Couldn't execute unknown instruction %d at line %d", instruction.Operation, instruction.Line)
	}
}

func (this *IntcodeMachine) ExecuteNoOp(instruction *IntcodeInstruction) {
	// Do nothing;
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteInc(instruction *IntcodeInstruction) {
	this.AccumulatedValue += int64(instruction.ParameterOne);
	this.InstructionPointer++;
}

func (this *IntcodeMachine) ExecuteJump(instruction *IntcodeInstruction) {
	this.InstructionPointer += int64(instruction.ParameterOne);
}