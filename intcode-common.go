package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntcodeInstruction struct{
	Operation int;
	Line int;
	ParameterOne int;
}

func (this *IntcodeInstruction) Parse(lineNum string) bool {
	parts := strings.Split(strings.TrimSpace(lineNum), " ");
	this.Operation = this.OpCodeToInt(parts[0]);
	arg, err := strconv.Atoi(parts[1]);
	if(err != nil){
		return false;
	}
	this.ParameterOne = arg;
	return true;
}

func (this *IntcodeInstruction) Describe() string {
	return fmt.Sprintf("%d %s %d", this.Line, this.IntToOpCode(this.Operation), this.ParameterOne);
}

func (this *IntcodeInstruction) IntToOpCode(opCode int) string {
	switch opCode{
		case IntCodeOpCodeJump:
			return "jmp";
		case IntCodeOpCodeInc:
			return "acc";
		case IntCodeOpCodeNoOp:
			return "nop";
	}
	return "unk";
}

func (this *IntcodeInstruction) OpCodeToInt(opCode string) int {
	switch opCode{
		case "jmp":
			return IntCodeOpCodeJump;
		case "acc":
			return IntCodeOpCodeInc;
		case "nop":
			return IntCodeOpCodeNoOp;
	}
	return IntCodeOpCodeUnknown;
}