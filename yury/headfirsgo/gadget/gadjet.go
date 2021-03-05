package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Now playing ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Now playing ", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Stopped!")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
