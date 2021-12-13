package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem12A struct {
	Graph *UndirectedGraph;
	TotalPathCount int;
	NodeCount int;
	Goal *Node;
}



func (this *Problem12A) Solve() {
	Log.Info("Problem 12A solver beginning!")

	file, err := os.Open("source-data/input-day-12a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)


	this.Graph = &UndirectedGraph{};
	this.Graph.Init();

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			parts := strings.Split(line, "-");
			if(len(parts) != 2){
				Log.Fatal("Incorrect number of line parts");
			}
			source := parts[0];
			dest := parts[1];
			srcN := this.Graph.GetOrCreateNode(source);
			dstN := this.Graph.GetOrCreateNode(dest);
			this.Graph.CreateEdge(srcN, dstN);
		}
	}

	pathSoFar := make([]string, 0);
	visitCounts := make(map[string]int);

	this.NodeCount = len(this.Graph.Nodes);

	start := this.Graph.GetOrCreateNode("start");
	this.Goal = this.Graph.GetOrCreateNode("end");
	this.Explore(start, pathSoFar, visitCounts);

	Log.Info("Finished exploration of %d nodes, found %d paths", this.NodeCount, this.TotalPathCount)
}

func (this *Problem12A) describePath(path []string) string {
	buff := "";
	for _, v := range path{
		if(buff != ""){
			buff += ",";
		}
		buff += v;
	}
	return buff;
}


func (this *Problem12A) Explore(next *Node, pathSoFar []string, visitCounts map[string]int) {
	cnt, _ := visitCounts[next.Label];
	max := 1;
	if(IsUpper(next.Label)){
		max = this.NodeCount * 2;
	}
	if(cnt >= max){
		return;
	}

	if(next == this.Goal){
		pathSoFar = append(pathSoFar, next.Label);
		this.TotalPathCount++;
		Log.Info("New path found %s - total paths %d", this.describePath(pathSoFar), this.TotalPathCount)
		return;
	}

	pathSoFar = append(pathSoFar, next.Label);
	visitCounts[next.Label] = cnt+1;

	for _, v := range next.Edges{
		nextNext := v.To;
		if(nextNext == next){
			nextNext = v.From;
		}

		nextCnt, _ := visitCounts[nextNext.Label];
		max = 1;

		if(IsUpper(nextNext.Label)){
			max = this.NodeCount * 2;
		}
		if(nextCnt >= max){
			continue;
		}

		cpyPath := make([]string, 0);
		cpyPath = append(cpyPath, pathSoFar...);

		cpyCount := make(map[string]int);
		for k, vC := range visitCounts {
			cpyCount[k] = vC;
		}

		this.Explore(nextNext, cpyPath, cpyCount);
	}
}
