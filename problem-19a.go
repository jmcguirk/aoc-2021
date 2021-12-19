package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Problem19A struct {
	AllMatrices []*M3;
}

type ProbeScanner struct {
	Id int;
	ScannedPoints []*IntVec3;
	CanonicalPoints []*IntVec3;
	TransformedPoints map[int][]*IntVec3;
	AssignedTransform *M3;
	AssignedTransformIndex int;
	PairedWith *ProbeScanner;
	IsCanonical bool;
	Pos *IntVec3;
}

func (this *ProbeScanner) Log() {
	Log.Info("--- scanner %d ---", this.Id);
	for _, v := range this.ScannedPoints{
		Log.Info("%s", v.ToString());
	}
}


func (this *Problem19A) Solve() {
	Log.Info("Problem 1A solver beginning!")


	file, err := os.Open("source-data/canonical-matricies.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	this.AllMatrices = make([]*M3, 0);


	var matrix *M3;

	rows := 0;

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
				if(matrix == nil){
					matrix = &M3{};
					rows = 0;
					matrix.Init();
					this.AllMatrices = append(this.AllMatrices, matrix);
				}
				cols := 0;
				rowParts := strings.Split(line, "\t");
				for _, v := range rowParts{
					parsed, err := strconv.Atoi(strings.TrimSpace(v));
					if(err != nil){
						Log.FatalError(err);
					}
					matrix.SetValue(cols, rows, parsed);
					cols++;
				}
				rows++;
		} else{
			if(matrix != nil){
				matrix = nil;
			}
		}
	}

	for _, v := range this.AllMatrices{
		v.CalculateInverse();
	}



	Log.Info("Finished parsing matrix file - loaded %d canonical matrices", len(this.AllMatrices));

	allScanners := make([]*ProbeScanner, 0);

	file, err = os.Open("source-data/input-day-19a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	var pScanner *ProbeScanner;
	scannerId := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			if(pScanner == nil){
				pScanner = &ProbeScanner{};
				pScanner.Id = scannerId;
				scannerId++;
				pScanner.ScannedPoints = make([]*IntVec3, 0);
				pScanner.TransformedPoints = make(map[int][]*IntVec3, 0);
				allScanners = append(allScanners, pScanner);
			} else{
				rowParts := strings.Split(line, ",");
				vec := &IntVec3{};

				for i, v := range rowParts{
					parsed, err := strconv.Atoi(strings.TrimSpace(v));
					if(err != nil){
						Log.FatalError(err);
					}
					if(i == 0){
						vec.X = parsed;
					} else if(i == 1){
						vec.Y = parsed;
					} else if(i == 2){
						vec.Z = parsed;
					}
				}
				pScanner.ScannedPoints = append(pScanner.ScannedPoints, vec);
			}
		} else{
			if(pScanner != nil){
				pScanner = nil;
			}
		}
	}

	Log.Info("Finished parsing, loaded %d scanners", len(allScanners));




	canonicalScanner := allScanners[0];
	canonicalScanner.IsCanonical = true;
	canonicalScanner.AssignedTransformIndex = 0;
	canonicalScanner.AssignedTransform = this.AllMatrices[0];
	canonicalScanner.CanonicalPoints = make([]*IntVec3, 0)
	canonicalScanner.CanonicalPoints = append(canonicalScanner.CanonicalPoints, canonicalScanner.ScannedPoints...);
	//Log.Fatal("Canoncial scanner has %d points", len(canonicalScanner.CanonicalPoints ));

	solvedScanners := make([]*ProbeScanner, 1);
	solvedScanners[0] = canonicalScanner;

	remainingScanners := make([]*ProbeScanner, 0);

	for i, v := range allScanners {
		if(i == 0){
			continue;
		}
		remainingScanners = append(remainingScanners, v);
	}

	for{
		if(len(remainingScanners) <= 0){
			break;
		}
		next := remainingScanners[0];
		remainingScanners = remainingScanners[1:];
		matchFound := false;

		for _, solved := range solvedScanners{
			//Log.Info("Consider %d vs %d", next.Id, solved.Id)

			bestMatrix, shiftAmount, overlapAmount := this.FindBestMatrix(solved, next);

			if(overlapAmount >= 12){
				matchFound = true;
				next.CanonicalPoints = make([]*IntVec3, len(next.ScannedPoints));
				for i, p := range next.ScannedPoints{
					rot := bestMatrix.MultiplyVec(p);
					shift := rot.Sub(shiftAmount);
					next.CanonicalPoints[i] = shift;
				}
				solvedScanners = append(solvedScanners, next);
				Log.Info("Aligned scanner %d, %d remain", next.Id, len(remainingScanners));
			}

			if(matchFound){
				break;
			}
		}

		if(!matchFound){
			//Log.Fatal("Failed to match");
			remainingScanners = append(remainingScanners, next);
		}
	}

	Log.Info("Solved all scanners");

	unique := make(map[string]int, 0);
	for _, pScanner := range solvedScanners{
		for _, v3 := range pScanner.CanonicalPoints{
			unique[v3.ToString()] = 1;
		}
	}

	Log.Info("Unique beacons found %d", len(unique));

	/*
	for _, pScanner := range allScanners {
		pScanner.Log();
		fmt.Println();
	}*/

}

func (this *Problem19A) FindBestOverlapAndShiftGivenRotation(referencePoints []*IntVec3, rotatedPoints []*IntVec3) (*IntVec3, int){
	overlapMap := make(map[string]int);
	for _, v := range referencePoints{
		overlapMap[v.ToString()] = 1;
	}

	mostOverlap := 0;
	leastShift := &IntVec3{};

	for i, canonicalPoint := range referencePoints{
		for j, rotatedPoint := range rotatedPoints{
			if (j > i){
				continue;
			}

			shiftAmount := rotatedPoint.Sub(canonicalPoint);
			shifted := make([]*IntVec3, len(rotatedPoints));
			for k, rp := range rotatedPoints{
				shifted[k] = rp.Sub(shiftAmount);
			}
			numOverlap := 0;
			for _, v := range shifted{
				_, exists := overlapMap[v.ToString()];
				if(exists){
					numOverlap++;
				}
			}
			if(numOverlap > mostOverlap){
				mostOverlap = numOverlap;
				leastShift = shiftAmount;
				if(mostOverlap >= 12){
					return leastShift, mostOverlap;
				}
			}
		}
	}
	return leastShift, mostOverlap;
}

func (this *Problem19A) FindBestMatrix(solvedScanner *ProbeScanner, scannerInQuestion *ProbeScanner) (*M3, *IntVec3, int){

	mostOverlap := math.MinInt32;
	var leastShift *IntVec3;
	var bestMatrix *M3;

	for _, m := range this.AllMatrices{
		rotated := make([]*IntVec3, len(scannerInQuestion.ScannedPoints));
		for i, v := range scannerInQuestion.ScannedPoints{
			rotated[i] = m.MultiplyVec(v);
		}
		shiftAmount, overlapAmount := this.FindBestOverlapAndShiftGivenRotation(solvedScanner.CanonicalPoints, rotated);
		//Log.Info("Overlap count %d", overlapAmount)
		if(overlapAmount >= 12){
			leastShift = shiftAmount;
			bestMatrix = m;
			mostOverlap = overlapAmount;
			break;
		}
	}
	return bestMatrix, leastShift, mostOverlap;
}
