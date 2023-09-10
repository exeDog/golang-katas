package main

func main() {
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		println(color, hex)
	}
}
