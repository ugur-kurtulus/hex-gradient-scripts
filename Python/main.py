# Constant definition for hex
hex = 0xabcdef

class Color:
    def __init__(self, name, color):
        self.name = name
        self.color = color

def main():
    blue_and_green = hex % 0x010000
    red = Color("red", hex >> 16)
    green = Color("blue", blue_and_green >> 8)
    blue = Color("green", blue_and_green % 0x0100)
    print([hex(color) for color in calculate_range(red, green, blue, 5, 0xFFFFFF)])

def calculate_range(red: Color, green: Color, blue: Color, step: int, endpoint_hex: int):
    color_range = []
    for i in range(step):
        red_value: hex = int(red.color + (endpoint_hex-red.color) * i / step)
        green_value: hex = int(green.color + (endpoint_hex-green.color) * i / step)
        blue_value: hex = int(blue.color + (endpoint_hex-blue.color) * i / step)
        color_range.append(red_value * (2**16) + green_value * (2**8) + blue_value)
	
    return color_range

main()