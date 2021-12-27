package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntcodeInstruction struct{
	Operation int;
	Line int;
	ParameterOne string;
	ParameterTwo string;
	ParameterTwoInt int64;
}

func (this *IntcodeInstruction) Parse(lineNum string) bool {
	parts := strings.Split(strings.TrimSpace(lineNum), " ");
	this.Operation = this.OpCodeToInt(parts[0]);

	this.ParameterOne = strings.TrimSpace(parts[1])

	if(len(parts) > 2){
		arg, err := strconv.Atoi(strings.TrimSpace(parts[2]));
		if(err != nil){
			this.ParameterTwo = parts[2];
		} else{
			this.ParameterTwoInt = int64(arg);
		}

	}
	return true;
}

func (this *IntcodeInstruction) Describe() string {
	if(this.ParameterTwo != ""){
		return fmt.Sprintf("%d %s %s %s", this.Line, this.IntToOpCode(this.Operation), this.ParameterOne, this.ParameterTwo);
	} else if(this.Operation == IntCodeOpCodeInput){
		return fmt.Sprintf("%d %s %s", this.Line, this.IntToOpCode(this.Operation), this.ParameterOne);
	} else{
		return fmt.Sprintf("%d %s %s %d", this.Line, this.IntToOpCode(this.Operation), this.ParameterOne, this.ParameterTwoInt);
	}

}

func (this *IntcodeInstruction) IntToOpCode(opCode int) string {
	switch opCode{
		case IntCodeOpCodeInput:
			return "inp";
		case IntCodeOpCodeAdd:
			return "add";
		case IntCodeOpCodeMul:
			return "mul";
		case IntCodeOpCodeDiv:
			return "div";
		case IntCodeOpCodeMod:
			return "mod";
		case IntCodeOpCodeEql:
			return "eql";
	}
	return "unk";
}

func (this *IntcodeInstruction) OpCodeToInt(opCode string) int {
	switch opCode{
		case "inp":
			return IntCodeOpCodeInput;
		case "add":
			return IntCodeOpCodeAdd;
		case "mul":
			return IntCodeOpCodeMul;
		case "div":
			return IntCodeOpCodeDiv;
		case "mod":
			return IntCodeOpCodeMod;
		case "eql":
			return IntCodeOpCodeEql;
	}
	return IntCodeOpCodeUnknown;
}