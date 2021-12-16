package main

import (
	"fmt"
	"math"
	"strconv"
)

const PacketOperation_Sum int = 0;
const PacketOperation_Product int = 1;
const PacketOperation_Min int = 2;
const PacketOperation_Max int = 3;
const PacketOperation_Literal int = 4;
const PacketOperation_GT int = 5;
const PacketOperation_LT int = 6;
const PacketOperation_EQ int = 7;

const PacketLengthMode_Bits int = 0;
const PacketLengthMode_Length int = 1;


type Packet struct {
	Version int;
	Operation int;
	LengthMode int;
	Length int;
	Depth int;
	UniqueId int;

	SubPackets []*Packet;
	LiteralValues []int;
}

type PacketDeserializer struct{
	Buffer []int;
	Position int;

	LastPacketId int;
	Depth int;
}

func (this *PacketDeserializer) DeserializeFromBinaryArray(bytes []int) *Packet {
	this.Buffer = bytes;
	this.Position = 0;
	this.LastPacketId = 1;
	return this.readPacket();
}


func (this *PacketDeserializer) readPacket() *Packet {
	packet := &Packet{};

	packet.UniqueId = this.LastPacketId;
	this.LastPacketId++;
	packet.Version = this.readIntegerOfSize(3);
	packet.Operation = this.readIntegerOfSize(3);
	packet.Depth = this.Depth;
	packet.LiteralValues = make([]int, 0);
	packet.SubPackets = make([]*Packet, 0);

	if(packet.Operation == PacketOperation_Literal){
		packet.LiteralValues = append(packet.LiteralValues, this.readLiteral());
	} else{

		this.Depth++;

		packet.LengthMode = this.readIntegerOfSize(1);
		if(packet.LengthMode == PacketLengthMode_Bits){
			packet.Length = this.readIntegerOfSize(15);
		} else{
			packet.Length = this.readIntegerOfSize(11);
		}

		if(packet.LengthMode == PacketLengthMode_Bits){
			startPos := this.Position;
			for {
				if(this.Position - startPos >= packet.Length){
					break;
				}
				packet.SubPackets = append(packet.SubPackets, this.readPacket());
			}
		} else{
			packetsRead := 0;
			for{
				if(packetsRead >= packet.Length){
					break;
				}
				packet.SubPackets = append(packet.SubPackets, this.readPacket());
				packetsRead++;
			}
		}
		this.Depth--;
	}
	return packet;
}

func (this *PacketDeserializer) readIntegerOfSize(size int) int {
	buff := "";
	for i := 0; i < size; i++{
		buff += fmt.Sprintf("%d", this.Buffer[this.Position]);
		this.Position++;
	}
	res, err := strconv.ParseInt(buff, 2, 64);
	if(err != nil){
		Log.FatalError(err);
	}
	return int(res);
}


func (this *PacketDeserializer) readLiteral() int {
	buff := "";
	for {
		lastBit := this.Buffer[this.Position];
		this.Position++;
		for i := 0; i < 4; i++{
			buff += fmt.Sprintf("%d", this.Buffer[this.Position]);
			this.Position++;
		}
		if(lastBit == 0){
			break;
		}
	}
	res, err := strconv.ParseInt(buff, 2, 64);
	if(err != nil){
		Log.FatalError(err);
	}
	return int(res);
}

func (this *Packet) Log() {
	depthBuffer := "";
	for i := 0; i < this.Depth; i++{
		depthBuffer += "    ";
	}
	Log.Info("%sPacket -- %d (%d)", depthBuffer, this.UniqueId, this.Depth);
	Log.Info("%s - Version: %d", depthBuffer, this.Version);
	Log.Info("%s - Operation: %d", depthBuffer, this.Operation);
	if(len(this.LiteralValues) > 0){
		Log.Info("%s - Literals: %s", depthBuffer, PrintIntArray(this.LiteralValues));
	}
	if(this.Operation != PacketOperation_Literal){
		Log.Info("%s - Length Mode: %d", depthBuffer, this.LengthMode);
	}
	for _, v := range this.SubPackets{
		v.Log();
	}
}

func (this *Packet) Checksum() int {
	sum := this.Version;
	for _, v := range this.SubPackets{
		sum += v.Checksum();
	}
	return sum;
}

func (this *Packet) Value() int {
	switch(this.Operation){
		case PacketOperation_Sum:
			return this.sumValue();
			break;
		case PacketOperation_Product:
			return this.productValue();
			break;
		case PacketOperation_Min:
			return this.minValue();
			break;
		case PacketOperation_Max:
			return this.maxValue();
			break;
		case PacketOperation_Literal:
			return this.literalValue();
			break;
		case PacketOperation_GT:
			return this.gtValue();
			break;
		case PacketOperation_LT:
			return this.ltValue();
			break;
		case PacketOperation_EQ:
			return this.eqValue();
			break;
	}
	Log.Fatal("Unsupported operation %d", this.Operation);
	return -1;
}

func (this *Packet) literalValue() int {
	return this.LiteralValues[0];
}

func (this *Packet) sumValue() int {
	sum := 0;
	for _, v := range this.SubPackets{
		sum += v.Value();
	}
	return sum;
}

func (this *Packet) productValue() int {
	m := 1;
	for _, v := range this.SubPackets{
		m *= v.Value();
	}
	return m;
}


func (this *Packet) minValue() int {
	x := math.MaxInt32;
	for _, v := range this.SubPackets{
		value := v.Value();
		if(value < x){
			x = value;
		}
	}
	return x;
}

func (this *Packet) maxValue() int {
	x := math.MinInt32;
	for _, v := range this.SubPackets{
		value := v.Value();
		if(value > x){
			x = value;
		}
	}
	return x;
}

func (this *Packet) gtValue() int {
	left := this.SubPackets[0].Value();
	right := this.SubPackets[1].Value();
	if(left > right){
		return 1;
	}
	return 0;
}

func (this *Packet) ltValue() int {
	left := this.SubPackets[0].Value();
	right := this.SubPackets[1].Value();
	if(left < right){
		return 1;
	}
	return 0;
}

func (this *Packet) eqValue() int {
	left := this.SubPackets[0].Value();
	right := this.SubPackets[1].Value();
	if(left == right){
		return 1;
	}
	return 0;
}