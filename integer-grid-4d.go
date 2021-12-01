package main

import (
	"math"
)

type IntegerGrid4D struct {
	Data map[int]*map[int]*map[int]*map[int]int;
}

func (this *IntegerGrid4D) Init() {
	this.Data = make(map[int]*map[int]*map[int]*map[int]int);
}

func (this *IntegerGrid4D) Clone() *IntegerGrid4D {
	res := &IntegerGrid4D{};
	res.Init();
	for k, v := range this.Data{
		cpy := make(map[int]*map[int]*map[int]int);
		for j, v2 := range *v{
			cln := make(map[int]*map[int]int);
			for g, v3 := range *v2{
				fcln := make(map[int]int);
				for n, v4 := range *v3{
					fcln[n] = v4;
				}
				cln[g] = &fcln;
			}
			cpy[j] = &cln;
		}
		res.Data[k] = &cpy;
	}

	return res;
}



func (this *IntegerGrid4D) SetValue(x int, y int, z int, w int, value int) {
	_, exists := this.Data[x];
	if(!exists){
		newMap := make(map[int]*map[int]*map[int]int);
		this.Data[x] = &newMap;
	}
	rowData := *this.Data[x];
	grid, exists := rowData[y];
	if(!exists){
		newGrid := make(map[int]*map[int]int);
		grid = &newGrid;
		rowData[y] = grid;
	}
	wgrid, exists := (*grid)[z];
	if(!exists){
		newGrid := make(map[int]int);
		wgrid = &newGrid;
		(*grid)[z] = wgrid;
	}
	(*wgrid)[w] = value;
}



func (this *IntegerGrid4D) MaxX() int {
	res := math.MinInt32;
	for x, _ := range this.Data{
		if(x > res){
			res = x;
		}
	}
	return res;
}

func (this *IntegerGrid4D) MinX() int {
	res := math.MaxInt32;
	for x, _ := range this.Data{
		if(x < res){
			res = x;
		}
	}
	return res;
}

func (this *IntegerGrid4D) MaxY() int {
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

func (this *IntegerGrid4D) MinY() int {
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

func (this *IntegerGrid4D) MinZ() int {
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

func (this *IntegerGrid4D) MaxZ() int {
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

func (this *IntegerGrid4D) MaxW() int {
	res := math.MinInt32;
	for _, vals := range this.Data{
		for _, cells := range *vals{
			for _, slice := range *cells {
				for w, _ := range *slice {
					if (w > res) {
						res = w;
					}
				}
			}
		}
	}
	return res;
}

func (this *IntegerGrid4D) MinW() int {
	res := math.MaxInt32;
	for _, vals := range this.Data{
		for _, cells := range *vals{
			for _, slice := range *cells {
				for w, _ := range *slice {
					if (w < res) {
						res = w;
					}
				}
			}
		}
	}
	return res;
}

func (this *IntegerGrid4D) GetValue(x int, y int, z int, w int) int {
	_, exists := this.Data[x];
	if(!exists){
		return 0;
	}
	rowData := *this.Data[x];
	cellData, exists := rowData[y];
	if(!exists){
		return 0;
	}
	zData, exists := (*cellData)[z];
	if(!exists){
		return 0;
	}
	val, exists := (*zData)[w];
	if(!exists){
		return 0;
	}
	return val;
}