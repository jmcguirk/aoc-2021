package main

import (
	"math"
)

type IntegerGrid3D struct {
	Data map[int]*map[int]*map[int]int;
}

func (this *IntegerGrid3D) Init() {
	this.Data = make(map[int]*map[int]*map[int]int);
}

func (this *IntegerGrid3D) Clone() *IntegerGrid3D {
	res := &IntegerGrid3D{};
	res.Init();


	for k, v := range this.Data{
		cpy := make(map[int]*map[int]int);
		for j, v2 := range *v{
			cln := make(map[int]int);
			for g, v3 := range *v2{
				cln[g] = v3;
			}
			cpy[j] = &cln;
		}
		res.Data[k] = &cpy;
	}

	return res;
}



func (this *IntegerGrid3D) SetValue(x int, y int, z int, value int) {
	_, exists := this.Data[x];
	if(!exists){
		newMap := make(map[int]*map[int]int);
		this.Data[x] = &newMap;
	}
	rowData := *this.Data[x];
	grid, exists := rowData[y];
	if(!exists){
		newGrid := make(map[int]int);
		grid = &newGrid;
		rowData[y] = grid;
	}
	(*grid)[z] = value;
}



func (this *IntegerGrid3D) MaxX() int {
	res := math.MinInt32;
	for x, _ := range this.Data{
		if(x > res){
			res = x;
		}
	}
	return res;
}

func (this *IntegerGrid3D) MinX() int {
	res := math.MaxInt32;
	for x, _ := range this.Data{
		if(x < res){
			res = x;
		}
	}
	return res;
}

func (this *IntegerGrid3D) MaxY() int {
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

func (this *IntegerGrid3D) MinY() int {
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

func (this *IntegerGrid3D) MinZ() int {
	res := math.MaxInt32;
	for _, vals := range this.Data{
		for _, cells := range *vals{
			for z, _ := range *cells {
				if (z < res) {
					res = z;
				}
			}
		}
	}
	return res;
}

func (this *IntegerGrid3D) MaxZ() int {
	res := math.MinInt32;
	for _, vals := range this.Data{
		for _, cells := range *vals{
			for z, _ := range *cells {
				if (z > res) {
					res = z;
				}
			}
		}
	}
	return res;
}

func (this *IntegerGrid3D) GetValue(x int, y int, z int) int {
	_, exists := this.Data[x];
	if(!exists){
		return 0;
	}
	rowData := *this.Data[x];
	cellData, exists := rowData[y];
	if(!exists){
		return 0;
	}
	val, exists := (*cellData)[z];
	if(!exists){
		return 0;
	}
	return val;
}


func (this *IntegerGrid3D) CountGreaterThan(threshold int) int {
	xMin := this.MinX();
	xMax := this.MaxX();

	yMin := this.MinY();
	yMax := this.MaxY();

	zMin := this.MinZ();
	zMax := this.MaxZ();

	total := 0;
	for j := yMin; j<= yMax; j++{
		for i := xMin; i<= xMax; i++{
			for k := zMin; k <= zMax; k++{
				val := this.GetValue(i, j, k);
				if(val > threshold){
					total++;
				}
			}

		}
	}
	return total;
}