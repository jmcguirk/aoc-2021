package main

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"strconv"
)

type M3 struct {
	Data map[int]*map[int]int;
	Inverse *M3;
}

func (this *M3) Init() {
	this.Data = make(map[int]*map[int]int);
}

func (this *M3) SetValue(x int, y int, value int) {
	_, exists := this.Data[x];
	if(!exists){
		newMap := make(map[int]int);
		this.Data[x] = &newMap;
	}
	rowData := *this.Data[x];
	_, exists = rowData[y];
	if(!exists){
		rowData[y] = 0;
	}
	rowData[y] = value;
}

func (this *M3) GetValue(x int, y int) int {
	_, exists := this.Data[x];
	if(!exists){
		return 0;
	}
	rowData := *this.Data[x];
	_, exists = rowData[y];
	if(!exists){
		return 0;
	}
	return rowData[y];
}

func (this *M3) HasValue(x int, y int) bool {
	_, exists := this.Data[x];
	if(!exists){
		return false;
	}
	rowData := *this.Data[x];
	_, exists = rowData[y];
	if(!exists){
		return false;
	}
	return true;
}


func (this *M3) Print() string {
	return this.PrintWithZero(".");
}



func (this *M3) MaxRow() int {
	res := math.MinInt32;
	for x, _ := range this.Data{
		if(x > res){
			res = x;
		}
	}
	return res;
}

func (this *M3) MinRow() int {
	res := math.MaxInt32;
	for x, _ := range this.Data{
		if(x < res){
			res = x;
		}
	}
	return res;
}

func (this *M3) MaxCol() int {
	res := math.MinInt32;
	for _, vals := range this.Data{
		for y, _ := range *vals{
			if(y > res){
				res = y;
			}
		}
	}
	return res;
}

func (this *M3) MinCol() int {
	res := math.MaxInt32;
	for _, vals := range this.Data{
		for y, _ := range *vals{
			if(y < res){
				res = y;
			}
		}
	}
	return res;
}



func (this *M3) MaxX() int {
	res := math.MinInt32;
	for x, _ := range this.Data{
		if(x > res){
			res = x;
		}
	}
	return res;
}

func (this *M3) MinX() int {
	res := math.MaxInt32;
	for x, _ := range this.Data{
		if(x < res){
			res = x;
		}
	}
	return res;
}

func (this *M3) MaxY() int {
	res := math.MinInt32;
	for _, vals := range this.Data{
		for y, _ := range *vals{
			if(y > res){
				res = y;
			}
		}
	}
	return res;
}

func (this *M3) MinY() int {
	res := math.MaxInt32;
	for _, vals := range this.Data{
		for y, _ := range *vals{
			if(y < res){
				res = y;
			}
		}
	}
	return res;
}


func (this *M3) PrintWithZero(zeroStr string) string {
	xMin := this.MinRow();
	xMax := this.MaxRow();

	yMin := this.MinCol();
	yMax := this.MaxCol();

	buff := "";
	for j := yMin; j<= yMax; j++{
		for i := xMin; i<= xMax; i++{
			if(!this.HasValue(i, j)){
				buff += ".";
			} else{
				val := this.GetValue(i, j);
				if(val != 0){
					buff += strconv.Itoa(this.GetValue(i, j));
				} else{
					if(val == 0){
						buff += zeroStr;
					}
				}
			}
		}
		buff += "\n";
	}
	//Log.Info("Furthest point is %d,%d", furthestX, furthestY);

	return buff;
}

func (this *M3) InverseMultiplyVec(val *IntVec3) *IntVec3 {
	return this.Inverse.MultiplyVec(val);
}

func (this *M3) MultiplyVec(val *IntVec3) *IntVec3 {
	res := &IntVec3{};

	res.X = 0;
	res.X += val.X * this.GetValue(0, 0);
	res.X += val.Y * this.GetValue(1, 0);
	res.X += val.Z * this.GetValue(2, 0);

	res.Y = 0;
	res.Y += val.X * this.GetValue(0, 1);
	res.Y += val.Y * this.GetValue(1, 1);
	res.Y += val.Z * this.GetValue(2, 1);

	res.Z = 0;
	res.Z += val.X * this.GetValue(0, 2);
	res.Z += val.Y * this.GetValue(1, 2);
	res.Z += val.Z * this.GetValue(2, 2);

	return res;
}

func (this *M3) CalculateInverse() {
	flat := make([]float64, 9);

	flat[0] = float64(this.GetValue(0, 0));
	flat[1] = float64(this.GetValue(1, 0));
	flat[2] = float64(this.GetValue(2, 0));

	flat[3] = float64(this.GetValue(0, 1));
	flat[4] = float64(this.GetValue(1, 1));
	flat[5] = float64(this.GetValue(2, 1));

	flat[6] = float64(this.GetValue(0, 2));
	flat[7] = float64(this.GetValue(1, 2));
	flat[8] = float64(this.GetValue(2, 2));

	a := mat.NewDense(3, 3, flat);

	zero := make([]float64, 9);
	b := mat.NewDense(3, 3, zero);

	err := b.Inverse(a);
	if err != nil {
		panic("A is not invertible")
	}

	this.Inverse = &M3{};
	this.Inverse.Init();

	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			this.Inverse.SetValue(i, j, int(a.At(i, j)));
		}
	}
}