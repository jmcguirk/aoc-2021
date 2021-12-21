package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2021");

	solver := Problem21B{};
	solver.Solve()
	Log.Info("Solver complete - exiting");
}
