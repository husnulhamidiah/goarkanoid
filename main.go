package main

func main() {
	free := buildLayout()
	defer free()

	render()
}
