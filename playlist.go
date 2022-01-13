package main

import (
	"fmt"
	"test_project/github.com/headfirstgo/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}
func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
