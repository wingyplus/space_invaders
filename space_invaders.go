// Space Invaders: the clone of Space Invaders see https://en.wikipedia.org/wiki/Space_Invaders
// for more details.

package main

func main() {
	game := NewGame()
	defer game.Cleanup()

	game.Init()

	for game.Running() {
		game.HandleEvent()
		game.Update()
		game.Render()
	}
}
