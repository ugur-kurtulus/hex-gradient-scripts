package main

import (
	"fmt"
)

type Color struct {
	name  string
	color int
}

func main() {

	var hex_input int
	hex_input = 0xFabcff

	red, green, blue := hex_to_rgb(hex_input)

	final_hexes := calculate_range(red, green, blue, 10, 0xFFFFFF)

	fmt.Println(fmt.Sprintf("%x", final_hexes))

}

func hex_to_rgb(hex int) (Color, Color, Color) {

	blue_and_green_hex := hex % 0x010000

	red := Color{"red", hex >> 16}
	green := Color{"green", blue_and_green_hex >> 8}
	blue := Color{"blue", blue_and_green_hex % 0x0100}

	return red, green, blue

}

func calculate_range(red Color, green Color, blue Color, step int, endpoint_color int) []int {

	var output_colors []int
	endpoint_red, endpoint_green, endpoint_blue := hex_to_rgb(endpoint_color)

	for i := 0; i <= step-1; i++ {

		red_value := red.color + (endpoint_red.color-red.color)*i/(step-1)
		green_value := green.color + (endpoint_green.color-green.color)*i/(step-1)
		blue_value := blue.color + (endpoint_blue.color-blue.color)*i/(step-1)

		output_colors = append(output_colors, (red_value<<16 + green_value<<8 + blue_value))
	}

	return output_colors

}
