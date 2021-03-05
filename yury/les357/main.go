package main

import (
	"src/yury/headfirsgo/gadget"
)

func playList(device gadget.TapePlayer, song []string) {
	for _, song := range song {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
