package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem16B struct {

}



func (this *Problem16B) Solve() {
	Log.Info("Problem 16B solver beginning!")

	file, err := os.Open("source-data/input-day-16b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	packetRaw := "";

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			packetRaw = line;
		}
	}

	deserializer := &PacketDeserializer{};
	packet := deserializer.DeserializeFromBinaryArray(HexToBinaryArray(packetRaw));
	Log.Info("Parsed packet system, value is %d", packet.Value());
}


