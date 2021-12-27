package main;

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const slopeEpsilon = 0.00001;
const distEpsilon = 0.00001;

func FormatDuration(duration time.Duration) string{
	return FormatDurationMS(int64(duration.Seconds() * 1000));
}

func FormatDurationMS(durationMS int64) string{
	if(durationMS < 1000){
		return strconv.FormatInt(durationMS, 10) + "ms";
	}
	secs := float64(durationMS) / 1000;
	if(secs < 300){
		return fmt.Sprintf("%.2fs", secs);
	}
	mins := secs / 60;
	if(mins < 60){
		return fmt.Sprintf("%.2fm", mins);
	}
	hr := mins / 60;
	return fmt.Sprintf("%.2fh", hr);
}


type IntVec2 struct{
	X 		int;
	Y		int;
}

type Interval struct{
	Min 		int;
	Max			int;
}

type Region struct {
	X Interval;
	Y Interval;
	Z Interval;
}


func (this Interval) Intersect(that Interval) Interval{
	res := Interval{};

	res.Min = this.Min;
	if(that.Min > res.Min){
		res.Min = that.Min;
	}

	res.Max = this.Max;
	if(that.Max < res.Max){
		res.Max = that.Max;
	}

	if(res.Min <= res.Max){
		return res;
	}

	zero := Interval{};
	zero.Min = 0;
	zero.Max = -1;
	return zero;
}

func (this Interval) Size() int{
	delta := this.Max - this.Min;
	delta += 1;
	if(delta > 0){
		return delta;
	}
	return 0;
}


func (this Interval) Equals(that Interval) bool{
	return this.Min == that.Min && this.Max == that.Max;
}

func (this Region) Equals(that Region) bool{
	return this.X.Equals(that.X) && this.Y.Equals(that.Y) && this.Z.Equals(that.Z);
}

func (this Region) Size() int{
	return this.X.Size() * this.Y.Size() * this.Z.Size();
}



func (this Interval) Split(that Interval) []Interval{
	res := make([]Interval, 0);
	intersection := this.Intersect(that);
	if(intersection.Size() <= 0){
		res = append(res, this);
		return res;
	}

	if(intersection.Min > this.Min){
		left := Interval{};
		left.Min = this.Min;
		left.Max = intersection.Min - 1;
		res = append(res, left);
	}

	res = append(res, intersection);

	if(intersection.Max < this.Max){
		right := Interval{};
		right.Min = intersection.Max + 1;
		right.Max = this.Max
		res = append(res, right);
	}
	return res;

}


func (this Region) SplitX(alongAxis Interval) []Region{
	res := make([]Region, 0);
	splits := this.X.Split(alongAxis);
	for _, v := range splits {
		subCube := Region{};
		subCube.X = v;
		subCube.Y = this.Y;
		subCube.Z = this.Z;
		res = append(res, subCube);
	}
	return res;
}

func (this Region) SplitY(alongAxis Interval) []Region{
	res := make([]Region, 0);
	splits := this.Y.Split(alongAxis);
	for _, v := range splits {
		subCube := Region{};
		subCube.X = this.X;
		subCube.Y = v;
		subCube.Z = this.Z;
		res = append(res, subCube);
	}
	return res;
}

func (this Region) SplitZ(alongAxis Interval) []Region{
	res := make([]Region, 0);
	splits := this.Z.Split(alongAxis);
	for _, v := range splits {
		subCube := Region{};
		subCube.X = this.X;
		subCube.Y = this.Y;
		subCube.Z = v;
		res = append(res, subCube);
	}
	return res;
}

func (this Region) Intersect(that Region) Region{
	res := Region{};
	res.X = this.X.Intersect(that.X);
	res.Y = this.Y.Intersect(that.Y);
	res.Z = this.Z.Intersect(that.Z);
	return res;
}


func (this Region) Subtract(that Region) []Region{
	res := make([]Region, 0);
	intersection := this.Intersect(that);
	if(intersection.X.Size() <= 0 || intersection.Y.Size() <= 0 || intersection.Z.Size() <= 0){ // One of our axis doesn't overlap this region
		res = append(res, this);
		return res;
	}
	if(this.Equals(intersection)){ // If we are subtracting the entire region, return an empty region list
		return res;
	}


	// These two cubes overlap, lets split up our entire region into subregions by splitting each access individually
	xCubes := this.SplitX(intersection.X);
	for _, xCube := range xCubes{
		yCubes := xCube.SplitY(intersection.Y);
		for _, yCube := range yCubes{
			zCubes := yCube.SplitZ(intersection.Z);
			for _, zCube := range zCubes{
				if(!zCube.Equals(intersection)){
					res = append(res, zCube);
				}
			}
		}
	}
	return res;
}


func (this *IntVec2) ManhattanDistance(other *IntVec2) int{
	xComp := this.X - other.X;
	if(xComp < 0){
		xComp *= -1;
	}
	yComp := this.Y - other.Y;
	if(yComp < 0){
		yComp *= -1;
	}
	return xComp + yComp;
}

func (this *IntVec2) Slope(other *IntVec2) float32{
	if(other.X == this.X){
		return math.MaxFloat32;
	}
	return float32((other.Y - this.Y)) / float32((other.X - this.X));
}

func (this *IntVec2) Distance(other *IntVec2) float32{
	distX := (this.X - other.X);
	distY := (this.Y - other.Y);
	return float32(math.Sqrt(float64((distX*distX) + (distY*distY))));
}

func (this *IntVec2) Angle (other *IntVec2) float32{
	return float32(math.Atan2(float64(other.Y - this.Y), float64(other.X - this.X)));
	//atan2(y2 - y1, x2 - x1) * 180 / PI;
}

func (this *IntVec2) ToString () string{
	return "[X:" + strconv.Itoa(this.X) + ",Y:" + strconv.Itoa(this.Y) + "]";
}

func (this *IntVec2) GetVisiblePoints(points []*IntVec2) []*IntVec2{
	res := make([]*IntVec2, 0);
	candidate := this;
	for _, neighbor := range points {
		if(neighbor == candidate){
			continue;
		}
		isOccluded := false;
		slopeN := candidate.Slope(neighbor);
		distN := candidate.Distance(neighbor);
		for _, occluder := range points {
			if(occluder == neighbor || occluder == candidate){
				continue;
			}
			slopeO := candidate.Slope(occluder);
			if(math.Abs(float64(slopeN - slopeO)) <= slopeEpsilon){

				if(math.Abs(float64((candidate.Distance(occluder) + neighbor.Distance(occluder)) - distN)) <= distEpsilon){
					isOccluded = true;
				}
			}
		}
		if(!isOccluded){
			res = append(res, neighbor);
		}
	}
	return res;
}

func Filter(target *IntVec2, points []*IntVec2) []*IntVec2{
	res := make([]*IntVec2, 0);
	for _, candidate := range points {
		if(candidate.X != target.X || candidate.Y != target.Y){
			res = append(res, candidate);
		}
	}
	return res;
}



func nthDigit(input *big.Int, n int64) int64 {
	var quotient big.Int
	quotient.Exp(big.NewInt(10), big.NewInt(n), nil)

	bigI := new(big.Int);
	bigI.Set(input);

	bigI.Div(bigI, &quotient)

	var result big.Int
	result.Mod(bigI, big.NewInt(10))

	return result.Int64();
}



// Perm calls f with each permutation of a.
func PermInt(a []int, f func([]int)) {
	permInt(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func permInt(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permInt(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permInt(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Perm calls f with each permutation of a.
func Perm(a []int64, f func([]int64)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int64, f func([]int64), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func nthDigit64(val int64, n int64) int {
	var quotient big.Int
	quotient.Exp(big.NewInt(10), big.NewInt(n), nil)

	bigI := big.NewInt(val);
	bigI.Div(bigI, &quotient)

	var result big.Int
	result.Mod(bigI, big.NewInt(10))

	return int(result.Int64());
}


func UpperAlphaCharacters() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return strings.ToUpper(string(p));
}

func IsGTEOrEqual(registersA []int, registersB []int) bool {
	for i, v := range registersA{
		if(v > registersB[i]){
			return true;
		}
		if(v < registersB[i]){
			return false;
		}
	}
	return true;
}

func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func IsGTE(registersA []int, registersB []int) bool {
	for i, v := range registersA{
		if(v > registersB[i]){
			return true;
		}
		if(v < registersB[i]){
			return false;
		}
	}
	return false;
}

func IsGTE64(registersA []int64, registersB []int64) bool {
	for i, v := range registersA{
		if(v > registersB[i]){
			return true;
		}
		if(v < registersB[i]){
			return false;
		}
	}
	return false;
}


func IsEQ(registersA []int, registersB []int) bool {
	for i, v := range registersA{
		if(v != registersB[i]){
			return false;
		}
	}
	return true;
}

const OrientationNorth = 0;
const OrientationEast = 1;
const OrientationSouth = 2;
const OrientationWest = 3;

const DirectionNorth = 1;
const DirectionSouth = 2;
const DirectionWest = 3;
const DirectionEast = 4;

func PrintOrientation(val int) string {
	switch (val) {
		case OrientationEast:
			return "E";
		case OrientationSouth:
			return "S";
		case OrientationWest:
			return "W";
	}
	return "N";
}

type IntVec3 struct{
	X 		int;
	Y		int;
	Z  		int;
}

func(this *IntVec3) Equals(that *IntVec3) bool {
	return this.X == that.X && this.Y == that.Y && this.Z == that.Z;
}

func(this *IntVec3) ToString() string{
	return fmt.Sprintf("%d,%d,%d", this.X, this.Y, this.Z);
}

func(this *IntVec3) Add(that *IntVec3) *IntVec3{
	res := &IntVec3{};
	res.X = this.X + that.X;
	res.Y = this.Y + that.Y;
	res.Z = this.Z + that.Z;
	return res;
}

func(this *IntVec3) Sub(that *IntVec3) *IntVec3{
	res := &IntVec3{};
	res.X = this.X - that.X;
	res.Y = this.Y - that.Y;
	res.Z = this.Z - that.Z;
	return res;
}


func(this *IntVec3) Manhattan(that *IntVec3) int{
	dX := this.X - that.X;
	if(dX < 0){
		dX = dX * -1;
	}
	dY := this.Y - that.Y;
	if(dY < 0){
		dY = dY * -1;
	}
	dZ := this.Z - that.Z;
	if(dZ < 0){
		dZ = dZ * -1;
	}
	return dX + dY + dZ;
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func (this *IntVec2) Clone() *IntVec2 {
	res := &IntVec2{};
	res.X = this.X;
	res.Y = this.Y;
	return res;
}

const TileIndexSize = 1000;
const TileIndexOffset = 50000;

func (this *IntVec2) TileIndex() int {
	return (this.X + TileIndexSize) + ((this.Y + TileIndexSize) * TileIndexOffset);
}

func (this *IntVec2) FromTileIndex(tileIndex int) {
	this.Y = (tileIndex/TileIndexOffset)-TileIndexSize;
	this.X = (tileIndex%TileIndexOffset)-TileIndexSize;
}

func (this *IntVec2) Eq(that *IntVec2) bool {
	return this.X == that.X && this.Y == that.Y;
}

func AllSubstrings(val string, n int) []string{
	res := make([]string, 0);
	for len := 1; len <= n; len++{
		for i := 0; i <= n - len; i++{
			j := i + len - 1;
			res = append(res, val[i:j]);
		}
	}
	return res;
}

// return list of primes less than
// source; https://stackoverflow.com/questions/21854191/generating-prime-numbers-in-go
func PrimesLessThan(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true { continue }
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func HexToBinaryArray(s string) []int {
	bin := make([]int, 0);

	for _, v := range s{
		switch(v){
			case '0':
				bin = append(bin, 0,0,0,0);
				break;
			case '1':
				bin = append(bin, 0,0,0,1);
				break;
			case '2':
				bin = append(bin, 0,0,1,0);
				break;
			case '3':
				bin = append(bin, 0,0,1,1);
				break;
			case '4':
				bin = append(bin, 0,1,0,0);
				break;
			case '5':
				bin = append(bin, 0,1,0,1);
				break;
			case '6':
				bin = append(bin, 0,1,1,0);
				break;
			case '7':
				bin = append(bin, 0,1,1,1);
				break;
			case '8':
				bin = append(bin, 1,0,0,0);
				break;
			case '9':
				bin = append(bin, 1,0,0,1);
				break;
			case 'A':
				bin = append(bin, 1,0,1,0);
				break;
			case 'B':
				bin = append(bin, 1,0,1,1);
				break;
			case 'C':
				bin = append(bin, 1,1,0,0);
				break;
			case 'D':
				bin = append(bin, 1,1,0,1);
				break;
			case 'E':
				bin = append(bin, 1,1,1,0);
				break;
			case 'F':
				bin = append(bin, 1,1,1,1);
				break;
		}
	}

	return bin;
}

func PrintBinaryArray(bytes []int) string  {
	buff := "";
	for _, v := range bytes{
		buff += fmt.Sprintf("%d", v);
	}
	return buff;
}

func PrintIntArray(vals []int) string  {
	buff := "[";
	for _, v := range vals{
		if(buff != "["){
			buff += ",";
		}
		buff += fmt.Sprintf("%d", v);
	}
	buff += "]";
	return buff;
}

