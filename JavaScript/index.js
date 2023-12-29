// First color value in hexadecimal
const hex = 0xffffff;

class Color {
  constructor(r, g, b) {
    this.r = r;
    this.g = g;
    this.b = b;
  }

  // Convert RGB to hexadecimal
  toHex() {
    return parseInt((this.r << 16) + (this.g << 8) + this.b, 10).toString(16);
  }
}

// Convert hexadecimal to RGB
function hexToRgb(hex) {
  let blue_and_green = hex % 0x010000;
  let r = hex >> 16;
  let g = blue_and_green >> 8;
  let b = blue_and_green % 0x0100;
  return new Color(r, g, b);
}

function calculateRange(color, range, endpoint) {
  color_range = [];
  for (let i = 0; i <= range - 1; i++) {
    color_range.push(
      new Color(
        color.r + (endpoint.r - color.r) * (i / (range - 1)),
        color.g + (endpoint.g - color.g) * (i / (range - 1)),
        color.b + (endpoint.b - color.b) * (i / (range - 1))
      ).toHex()
    );
  }
  return color_range;
}

console.log(calculateRange(hexToRgb(hex), 10, hexToRgb(0x000000)));
